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
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

var _ Protocol = &BexLPPriceProvider{}

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

// LPTokenPrice returns the current price of the protocol's LP token in USD
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

	// Avoid division by zero
	if totalSupply.Sign() == 0 {
		err := errors.New("totalSupply is zero, cannot calculate LP token price")
		b.logger.Error().Err(err).Msg("Invalid totalSupply")
		return "", err
	}

	totalValue, err := b.totalValue(ctx)
	if err != nil {
		return "", err
	}

	totalSupplyDecimal := NormalizeAmount(totalSupply, b.config.LPTDecimals)
	pricePerToken := totalValue.Div(totalSupplyDecimal)

	b.logger.Debug().
		Str("totalValue", totalValue.String()).
		Str("totalSupply", totalSupplyDecimal.String()).
		Str("pricePerToken", pricePerToken.String()).
		Msg("LP token price calculated successfully")

	return pricePerToken.StringFixed(roundingDecimals), nil
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
