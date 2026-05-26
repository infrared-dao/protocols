package protocols

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/infrared-dao/protocols/fetchers"
	"github.com/infrared-dao/protocols/internal/sc"
	"github.com/infrared-dao/protocols/multicall3"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

var _ Protocol = &BexLPPriceProvider{}

// Compile-time check that BexLPPriceProvider satisfies BatchablePriceProvider.
var _ BatchablePriceProvider = &BexLPPriceProvider{}

type BexPoolConfig struct {
	PoolID      [32]byte `json:"poolid"`
	LPTDecimals uint     `json:"lpt_decimals"`
}

// BexLPPriceProvider defines the provider for BEX LP price and Pool TVL.
type BexLPPriceProvider struct {
	vaultAddress  common.Address
	poolAddress   common.Address
	block         *big.Int
	priceMap      map[string]Price
	logger        zerolog.Logger
	configBytes   []byte
	config        *BexPoolConfig
	vaultContract *sc.BalancerVault
	poolContract  *sc.BalancerBasePool
}

// NewBexLPPriceProvider creates a new instance of the BexLPPriceProvider.
func NewBexLPPriceProvider(
	vaultAddress common.Address,
	poolAddress common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *BexLPPriceProvider {
	b := &BexLPPriceProvider{
		vaultAddress: vaultAddress,
		poolAddress:  poolAddress,
		block:        block,
		priceMap:     prices,
		logger:       logger,
		configBytes:  config,
	}
	return b
}

// Initialize checks the configuration/data and instantiates the Vault and Base Pool contracts.
func (b *BexLPPriceProvider) Initialize(ctx context.Context, client bind.ContractBackend, _ fetchers.HttpClient) error {
	var err error

	b.config = &BexPoolConfig{}
	err = json.Unmarshal(b.configBytes, b.config)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	b.vaultContract, err = sc.NewBalancerVault(b.vaultAddress, client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instantiate Balancer Vault contract")
		return err
	}

	b.poolContract, err = sc.NewBalancerBasePool(b.poolAddress, client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instantiate Balancer Base Pool contract on LP Token")
		return err
	}
	return nil
}

// LPTokenPrice returns the current price of the protocol's LP token in USD.
//
// Legacy path — issues two sequential eth_calls (pool.getActualSupply +
// vault.getPoolTokens). The batchable path (PriceReads + ComputePrice)
// dispatches both as one Multicall3.aggregate3. Both paths share
// computeLPPriceFromReads so prices match for the same inputs.
func (b *BexLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: b.block,
	}

	// Using GetActualSupply because this is how much is circulating for pools which lock up some LP tokens
	totalSupply, err := b.poolContract.GetActualSupply(opts)
	if err != nil {
		return "", err
	}

	balances, err := b.getUnderlyingBalances(ctx)
	if err != nil {
		return "", err
	}

	price, err := b.computeLPPriceFromReads(totalSupply, balances)
	if err != nil {
		return "", err
	}
	return price.StringFixed(roundingDecimals), nil
}

// PriceReads describes the two eth_calls needed to compute the LP price:
// pool.getActualSupply (against the pool contract) and vault.getPoolTokens
// (against the Balancer vault contract, with the pool ID as argument).
//
// Targets differ across the two calls — multicall3 handles that naturally
// since each Call3 has its own target.
func (b *BexLPPriceProvider) PriceReads() ([]multicall3.Call3, error) {
	poolABI, err := sc.BalancerBasePoolMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("bex: get pool ABI: %w", err)
	}
	totalSupplyData, err := poolABI.Pack("getActualSupply")
	if err != nil {
		return nil, fmt.Errorf("bex: pack getActualSupply: %w", err)
	}

	vaultABI, err := sc.BalancerVaultMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("bex: get vault ABI: %w", err)
	}
	poolTokensData, err := vaultABI.Pack("getPoolTokens", b.config.PoolID)
	if err != nil {
		return nil, fmt.Errorf("bex: pack getPoolTokens: %w", err)
	}

	return []multicall3.Call3{
		{Target: b.poolAddress, AllowFailure: false, CallData: totalSupplyData},
		{Target: b.vaultAddress, AllowFailure: false, CallData: poolTokensData},
	}, nil
}

// ComputePrice decodes the two responses (totalSupply + pool tokens) and
// computes the LP token price. Pure — no network I/O.
func (b *BexLPPriceProvider) ComputePrice(responses []multicall3.Result3) (decimal.Decimal, error) {
	if len(responses) != 2 {
		return decimal.Zero, fmt.Errorf("bex: expected 2 responses, got %d", len(responses))
	}
	if !responses[0].Success {
		return decimal.Zero, errors.New("bex: getActualSupply call reverted in multicall")
	}
	if !responses[1].Success {
		return decimal.Zero, errors.New("bex: getPoolTokens call reverted in multicall")
	}

	// Unpack getActualSupply → *big.Int
	poolABI, err := sc.BalancerBasePoolMetaData.GetAbi()
	if err != nil {
		return decimal.Zero, fmt.Errorf("bex: get pool ABI: %w", err)
	}
	tsOut, err := poolABI.Methods["getActualSupply"].Outputs.Unpack(responses[0].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("bex: unpack getActualSupply: %w", err)
	}
	if len(tsOut) != 1 {
		return decimal.Zero, fmt.Errorf("bex: getActualSupply returned %d values, want 1", len(tsOut))
	}
	totalSupply, ok := tsOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("bex: getActualSupply value type %T, want *big.Int", tsOut[0])
	}

	// Unpack getPoolTokens → (tokens []common.Address, balances []*big.Int, lastChangeBlock *big.Int)
	vaultABI, err := sc.BalancerVaultMetaData.GetAbi()
	if err != nil {
		return decimal.Zero, fmt.Errorf("bex: get vault ABI: %w", err)
	}
	ptOut, err := vaultABI.Methods["getPoolTokens"].Outputs.Unpack(responses[1].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("bex: unpack getPoolTokens: %w", err)
	}
	if len(ptOut) < 2 {
		return decimal.Zero, fmt.Errorf("bex: getPoolTokens returned %d values, want >=2", len(ptOut))
	}
	tokens, ok := ptOut[0].([]common.Address)
	if !ok {
		return decimal.Zero, fmt.Errorf("bex: tokens value type %T, want []common.Address", ptOut[0])
	}
	rawBalances, ok := ptOut[1].([]*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("bex: balances value type %T, want []*big.Int", ptOut[1])
	}
	if len(tokens) != len(rawBalances) {
		return decimal.Zero, fmt.Errorf("bex: tokens (%d) and balances (%d) length mismatch", len(tokens), len(rawBalances))
	}

	// Filter out the pool's own LP token from the balances (same logic as
	// getUnderlyingBalances) — some pools lock LP tokens in themselves.
	balances := make(map[string]*big.Int, len(tokens))
	poolAddressLower := strings.ToLower(b.poolAddress.Hex())
	for i, t := range tokens {
		token := strings.ToLower(t.Hex())
		if token == poolAddressLower {
			continue
		}
		balances[token] = rawBalances[i]
	}

	return b.computeLPPriceFromReads(totalSupply, balances)
}

// computeLPPriceFromReads is the shared math between the legacy
// LPTokenPrice and the new batchable ComputePrice paths. Pure: takes
// the raw read values + config/prices, returns the LP token price.
func (b *BexLPPriceProvider) computeLPPriceFromReads(totalSupply *big.Int, balances map[string]*big.Int) (decimal.Decimal, error) {
	if totalSupply.Sign() == 0 {
		err := errors.New("totalSupply is zero, cannot calculate LP token price")
		b.logger.Error().Err(err).Msg("Invalid totalSupply")
		return decimal.Zero, err
	}

	totalValue := decimal.Zero
	for token, balance := range balances {
		price, err := b.getPrice(token)
		if err != nil {
			return decimal.Zero, err
		}
		balanceDecimal := NormalizeAmount(balance, price.Decimals)
		totalValue = totalValue.Add(balanceDecimal.Mul(price.Price))
	}

	totalSupplyDecimal := NormalizeAmount(totalSupply, b.config.LPTDecimals)
	pricePerToken := totalValue.Div(totalSupplyDecimal)

	b.logger.Debug().
		Str("totalValue", totalValue.String()).
		Str("totalSupply", totalSupplyDecimal.String()).
		Str("pricePerToken", pricePerToken.String()).
		Msg("LP token price calculated successfully")

	return pricePerToken, nil
}

// TVL returns the Total Value Locked in the pool in USD cents (1 USD = 100 cents).
func (b *BexLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := b.totalValue(ctx)
	if err != nil {
		return "", err
	}

	return totalValue.StringFixed(roundingDecimals), nil
}

func (b *BexLPPriceProvider) GetConfig(ctx context.Context, poolAddress string, client bind.ContractBackend) ([]byte, error) {
	var err error
	if !common.IsHexAddress(poolAddress) {
		err = fmt.Errorf("invalid smart contract address, '%s'", poolAddress)
		return nil, err
	}

	poolContract, err := sc.NewBalancerBasePool(common.HexToAddress(poolAddress), client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instantiate Balancer Base Pool contract on LP Token")
		return nil, err
	}

	bpc := BexPoolConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	// returns as [32]byte
	poolID, err := poolContract.GetPoolId(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain poolID for bex pool %s, %v", poolAddress, err)
		return nil, err
	}
	bpc.PoolID = poolID

	// decimals is uint8
	decimals, err := poolContract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for LP token %s, %v", poolAddress, err)
		return nil, err
	}
	bpc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(bpc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (b *BexLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	b.block = block
	if prices != nil {
		b.priceMap = prices
	}
}

// TVLBreakdown returns the breakdown of TVL by underlying tokens.
func (b *BexLPPriceProvider) TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error) {
	balanceData, err := b.getUnderlyingBalances(ctx)
	if err != nil {
		return nil, err
	}

	breakdown := make(map[string]TokenTVL, len(balanceData))
	type tokenData struct {
		amount   decimal.Decimal
		usdValue decimal.Decimal
		price    *Price
	}
	tokenCache := make(map[string]tokenData, len(balanceData))
	totalValue := decimal.Zero

	for token, balance := range balanceData {
		price, err := b.getPrice(token)
		if err != nil {
			return nil, err
		}
		balanceDecimal := NormalizeAmount(balance, price.Decimals)
		usdValue := balanceDecimal.Mul(price.Price)

		tokenCache[token] = tokenData{
			amount:   balanceDecimal,
			usdValue: usdValue,
			price:    price,
		}
		totalValue = totalValue.Add(usdValue)
	}

	// Build breakdown with pre-calculated values
	for token, data := range tokenCache {
		// Calculate ratio (handle zero TVL case)
		var ratio decimal.Decimal
		if totalValue.IsZero() {
			ratio = decimal.Zero
		} else {
			ratio = data.usdValue.Div(totalValue)
		}

		breakdown[token] = TokenTVL{
			TokenAddress: token,
			TokenSymbol:  data.price.TokenName,
			Amount:       data.amount,
			USDValue:     data.usdValue,
			Ratio:        ratio,
		}
	}

	b.logger.Debug().Msg("TVL breakdown calculated successfully")

	return breakdown, nil
}

// Internal Helper methods not able to be called except in this file

func (b *BexLPPriceProvider) totalValue(ctx context.Context) (decimal.Decimal, error) {
	var err error

	// Fetch underlying balances as map[string]*big.Int
	balanceData, err := b.getUnderlyingBalances(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	b.logger.Debug().
		Msgf("Token Balances: %+v", balanceData)

	totalValue := decimal.Zero
	for token, balance := range balanceData {
		price, err := b.getPrice(token)
		if err != nil {
			return decimal.Zero, err
		}
		balanceDecimal := NormalizeAmount(balance, price.Decimals)
		totalValue = totalValue.Add(balanceDecimal.Mul(price.Price))
	}

	return totalValue, nil
}

func (b *BexLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := b.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		b.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getUnderlyingBalances fetches the underlying virtual token supply for each token.
func (b *BexLPPriceProvider) getUnderlyingBalances(ctx context.Context) (map[string]*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: b.block,
	}

	/********************************************
		Returns data of type:
		struct {
			Tokens          []common.Address
			Balances        []*big.Int
			LastChangeBlock *big.Int
		}
	********************************************/
	poolTokens, err := b.vaultContract.GetPoolTokens(opts, b.config.PoolID)
	if err != nil {
		return nil, fmt.Errorf("failed to get pool tokens and balances from bex, err: %w", err)
	}

	var balanceData = make(map[string]*big.Int)
	for i, tokenAddress := range poolTokens.Tokens {
		token := strings.ToLower(tokenAddress.Hex())

		//verify this is always sound for all pool types
		if token == strings.ToLower(b.poolAddress.Hex()) {
			// ignore when some of the LP token is locked in pool itself
			// this is why should use actualSupply instead of totalSupply
			continue
		}

		balance := poolTokens.Balances[i]
		balanceData[token] = balance
	}

	return balanceData, nil
}
