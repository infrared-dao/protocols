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

var _ Protocol = &BurrBearLPPriceProvider{}
var _ BatchablePriceProvider = &BurrBearLPPriceProvider{}

// BurrBear is based on BalancerV2 which is the same codebase which BEX uses on mainnet
// Decided to implement it as a parallel code instead of a wrapper so it can get in config
//  the VaultContract address off the contract itself so we don't need to pass this in as
//  an env var like we originally did for BEX. This leads to different New function signature
// It can use the same Balancer contracts so there are no new abi or smart contract bindings

type BurrBearPoolConfig struct {
	VaultContract string   `json:"vault_contract"`
	PoolID        [32]byte `json:"poolid"`
	LPTDecimals   uint     `json:"lpt_decimals"`
}

// BurrBearLPPriceProvider defines the provider for BEX LP price and Pool TVL.
type BurrBearLPPriceProvider struct {
	poolAddress   common.Address
	block         *big.Int
	priceMap      map[string]Price
	logger        zerolog.Logger
	configBytes   []byte
	config        *BurrBearPoolConfig
	vaultContract *sc.BalancerVault
	poolContract  *sc.BalancerBasePool
}

// NewBurrBearLPPriceProvider creates a new instance of the BurrBearLPPriceProvider.
func NewBurrBearLPPriceProvider(
	poolAddress common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *BurrBearLPPriceProvider {
	b := &BurrBearLPPriceProvider{
		poolAddress: poolAddress,
		block:       block,
		priceMap:    prices,
		logger:      logger,
		configBytes: config,
	}
	return b
}

// Initialize checks the configuration/data and instantiates the Vault and Base Pool contracts.
func (bb *BurrBearLPPriceProvider) Initialize(ctx context.Context, client bind.ContractBackend, _ fetchers.HttpClient) error {
	var err error

	bb.config = &BurrBearPoolConfig{}
	err = json.Unmarshal(bb.configBytes, bb.config)
	if err != nil {
		bb.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	vaultAddress := common.HexToAddress(bb.config.VaultContract)

	bb.vaultContract, err = sc.NewBalancerVault(vaultAddress, client)
	if err != nil {
		bb.logger.Error().Err(err).Msg("failed to instantiate Balancer Vault contract")
		return err
	}

	bb.poolContract, err = sc.NewBalancerBasePool(bb.poolAddress, client)
	if err != nil {
		bb.logger.Error().Err(err).Msg("failed to instantiate Balancer Base Pool contract on LP Token")
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
func (bb *BurrBearLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: bb.block,
	}
	totalSupply, err := bb.poolContract.GetActualSupply(opts)
	if err != nil {
		return "", err
	}
	balances, err := bb.getUnderlyingBalances(ctx)
	if err != nil {
		return "", err
	}
	price, err := bb.computeLPPriceFromReads(totalSupply, balances)
	if err != nil {
		return "", err
	}
	return price.StringFixed(roundingDecimals), nil
}

func (bb *BurrBearLPPriceProvider) PriceReads() ([]multicall3.Call3, error) {
	poolABI, err := sc.BalancerBasePoolMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("burrbear: pool ABI: %w", err)
	}
	totalSupplyData, err := poolABI.Pack("getActualSupply")
	if err != nil {
		return nil, fmt.Errorf("burrbear: pack getActualSupply: %w", err)
	}
	vaultABI, err := sc.BalancerVaultMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("burrbear: vault ABI: %w", err)
	}
	poolTokensData, err := vaultABI.Pack("getPoolTokens", bb.config.PoolID)
	if err != nil {
		return nil, fmt.Errorf("burrbear: pack getPoolTokens: %w", err)
	}
	return []multicall3.Call3{
		{Target: bb.poolAddress, AllowFailure: false, CallData: totalSupplyData},
		{Target: common.HexToAddress(bb.config.VaultContract), AllowFailure: false, CallData: poolTokensData},
	}, nil
}

func (bb *BurrBearLPPriceProvider) ComputePrice(responses []multicall3.Result3) (decimal.Decimal, error) {
	if len(responses) != 2 {
		return decimal.Zero, fmt.Errorf("burrbear: expected 2 responses, got %d", len(responses))
	}
	if !responses[0].Success {
		return decimal.Zero, errors.New("burrbear: getActualSupply call reverted in multicall")
	}
	if !responses[1].Success {
		return decimal.Zero, errors.New("burrbear: getPoolTokens call reverted in multicall")
	}
	poolABI, err := sc.BalancerBasePoolMetaData.GetAbi()
	if err != nil {
		return decimal.Zero, fmt.Errorf("burrbear: pool ABI: %w", err)
	}
	tsOut, err := poolABI.Methods["getActualSupply"].Outputs.Unpack(responses[0].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("burrbear: unpack getActualSupply: %w", err)
	}
	totalSupply, ok := tsOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("burrbear: totalSupply type %T", tsOut[0])
	}
	vaultABI, err := sc.BalancerVaultMetaData.GetAbi()
	if err != nil {
		return decimal.Zero, fmt.Errorf("burrbear: vault ABI: %w", err)
	}
	ptOut, err := vaultABI.Methods["getPoolTokens"].Outputs.Unpack(responses[1].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("burrbear: unpack getPoolTokens: %w", err)
	}
	if len(ptOut) < 2 {
		return decimal.Zero, fmt.Errorf("burrbear: getPoolTokens returned %d values, want >=2", len(ptOut))
	}
	tokens, ok := ptOut[0].([]common.Address)
	if !ok {
		return decimal.Zero, fmt.Errorf("burrbear: tokens type %T", ptOut[0])
	}
	rawBalances, ok := ptOut[1].([]*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("burrbear: balances type %T", ptOut[1])
	}
	if len(tokens) != len(rawBalances) {
		return decimal.Zero, fmt.Errorf("burrbear: tokens (%d) and balances (%d) length mismatch", len(tokens), len(rawBalances))
	}
	balances := make(map[string]*big.Int, len(tokens))
	poolAddressLower := strings.ToLower(bb.poolAddress.Hex())
	for i, t := range tokens {
		token := strings.ToLower(t.Hex())
		if token == poolAddressLower {
			continue
		}
		balances[token] = rawBalances[i]
	}
	return bb.computeLPPriceFromReads(totalSupply, balances)
}

func (bb *BurrBearLPPriceProvider) computeLPPriceFromReads(totalSupply *big.Int, balances map[string]*big.Int) (decimal.Decimal, error) {
	if totalSupply.Sign() == 0 {
		err := errors.New("totalSupply is zero, cannot calculate LP token price")
		bb.logger.Error().Err(err).Msg("Invalid totalSupply")
		return decimal.Zero, err
	}
	totalValue := decimal.Zero
	for token, balance := range balances {
		price, err := bb.getPrice(token)
		if err != nil {
			return decimal.Zero, err
		}
		balanceDecimal := NormalizeAmount(balance, price.Decimals)
		totalValue = totalValue.Add(balanceDecimal.Mul(price.Price))
	}
	totalSupplyDecimal := NormalizeAmount(totalSupply, bb.config.LPTDecimals)
	pricePerToken := totalValue.Div(totalSupplyDecimal)
	bb.logger.Debug().Str("pricePerToken", pricePerToken.String()).Msg("LP token price calculated successfully")
	return pricePerToken, nil
}

// TVL returns the Total Value Locked in the pool in USD cents (1 USD = 100 cents).
func (bb *BurrBearLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := bb.totalValue(ctx)
	if err != nil {
		return "", err
	}

	return totalValue.StringFixed(roundingDecimals), nil
}

func (bb *BurrBearLPPriceProvider) GetConfig(ctx context.Context, poolAddress string, client bind.ContractBackend) ([]byte, error) {
	var err error
	if !common.IsHexAddress(poolAddress) {
		err = fmt.Errorf("invalid smart contract address, '%s'", poolAddress)
		return nil, err
	}

	poolContract, err := sc.NewBalancerBasePool(common.HexToAddress(poolAddress), client)
	if err != nil {
		bb.logger.Error().Err(err).Msg("failed to instantiate Balancer Base Pool contract on LP Token")
		return nil, err
	}

	bbpc := BurrBearPoolConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	// returns as common.Address, need to get a string for saving with Hex() call
	vaultAddress, err := poolContract.GetVault(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain poolID for bex pool %s, %v", poolAddress, err)
		return nil, err
	}
	bbpc.VaultContract = vaultAddress.Hex()

	// returns as [32]byte
	poolID, err := poolContract.GetPoolId(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain poolID for bex pool %s, %v", poolAddress, err)
		return nil, err
	}
	bbpc.PoolID = poolID

	// decimals is uint8
	decimals, err := poolContract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for LP token %s, %v", poolAddress, err)
		return nil, err
	}
	bbpc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(bbpc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (bb *BurrBearLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	bb.block = block
	if prices != nil {
		bb.priceMap = prices
	}
}

// TVLBreakdown returns the breakdown of TVL by underlying tokens.
// TODO: Implement TVL breakdown for BurrBear protocol
func (bb *BurrBearLPPriceProvider) TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error) {
	return nil, ErrTVLBreakdownNotImplemented
}

// Internal Helper methods not able to be called except in this file

func (bb *BurrBearLPPriceProvider) totalValue(ctx context.Context) (decimal.Decimal, error) {
	var err error

	// Fetch underlying balances as map[string]*big.Int
	balanceData, err := bb.getUnderlyingBalances(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	bb.logger.Debug().
		Msgf("Token Balances: %+v", balanceData)

	totalValue := decimal.Zero
	for token, balance := range balanceData {
		price, err := bb.getPrice(token)
		if err != nil {
			return decimal.Zero, err
		}
		balanceDecimal := NormalizeAmount(balance, price.Decimals)
		totalValue = totalValue.Add(balanceDecimal.Mul(price.Price))
	}

	return totalValue, nil
}

func (bb *BurrBearLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := bb.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		bb.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getUnderlyingBalances fetches the underlying virtual token supply for each token.
func (bb *BurrBearLPPriceProvider) getUnderlyingBalances(ctx context.Context) (map[string]*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: bb.block,
	}

	/********************************************
		Returns data of type:
		struct {
			Tokens          []common.Address
			Balances        []*big.Int
			LastChangeBlock *big.Int
		}
	********************************************/
	poolTokens, err := bb.vaultContract.GetPoolTokens(opts, bb.config.PoolID)
	if err != nil {
		return nil, fmt.Errorf("failed to get pool tokens and balances from bex, err: %w", err)
	}

	var balanceData = make(map[string]*big.Int)
	for i, tokenAddress := range poolTokens.Tokens {
		token := strings.ToLower(tokenAddress.Hex())

		//verify this is always sound for all pool types
		if token == strings.ToLower(bb.poolAddress.Hex()) {
			// ignore when some of the LP token is locked in pool itself
			// this is why should use actualSupply instead of totalSupply
			continue
		}

		balance := poolTokens.Balances[i]
		balanceData[token] = balance
	}

	return balanceData, nil
}
