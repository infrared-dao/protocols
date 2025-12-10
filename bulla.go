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
	"github.com/infrared-dao/protocols/internal/sc"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

var _ Protocol = &BullaLPPriceProvider{}

// The bulla abi given is a contract type called hypervisor which manages automated pools
// For other non-automated pools we would need a slightly different adapter because the function
// getPoolBalances wouldn't exist on those contract types, we could generalize to work across
// all pools by looking at the reserves instead which has a helper call on the base Algrabra Pool

// BullaConfig defines the configuration for Bulla Exchange adapter
type BullaConfig struct {
	Token0      string `json:"token0"`
	Token1      string `json:"token1"`
	LPTDecimals uint   `json:"lpt_decimals"`
	// Add any other Bulla-specific configuration fields here
}

// BullaLPPriceProvider defines the provider for Bulla Exchange LP price and TVL.
type BullaLPPriceProvider struct {
	address     common.Address
	block       *big.Int
	priceMap    map[string]Price
	logger      zerolog.Logger
	configBytes []byte
	config      *BullaConfig
	contract    *sc.Bulla
	// Add any Bulla-specific contract interfaces here
}

// NewBullaLPPriceProvider creates a new instance of the BullaLPPriceProvider.
func NewBullaLPPriceProvider(
	address common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *BullaLPPriceProvider {
	b := &BullaLPPriceProvider{
		address:     address,
		block:       block,
		priceMap:    prices,
		logger:      logger,
		configBytes: config,
	}
	return b
}

// Initialize checks the configuration/data provided and instantiates the Bulla smart contract.
func (b *BullaLPPriceProvider) Initialize(ctx context.Context, client bind.ContractBackend) error {
	var err error

	b.config = &BullaConfig{}
	err = json.Unmarshal(b.configBytes, b.config)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	// Validate that we have price data for the tokens
	_, ok := b.priceMap[b.config.Token0]
	if !ok {
		err = fmt.Errorf("no price data found for token0 (%s)", b.config.Token0)
		b.logger.Error().Msg(err.Error())
		return err
	}

	_, ok = b.priceMap[b.config.Token1]
	if !ok {
		err = fmt.Errorf("no price data found for token1 (%s)", b.config.Token1)
		b.logger.Error().Msg(err.Error())
		return err
	}

	// Initialize Bulla contract
	b.contract, err = sc.NewBulla(b.address, client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instantiate Bulla smart contract")
		return err
	}

	return nil
}

// LPTokenPrice returns the current price of the protocol's LP token in USD.
func (b *BullaLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	// Fetch total supply
	totalSupply, err := b.getTotalSupply(ctx)
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

// TVL returns the Total Value Locked in the protocol in USD.
func (b *BullaLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := b.totalValue(ctx)
	if err != nil {
		return "", err
	}

	b.logger.Debug().
		Str("totalValue", totalValue.String()).
		Msg("TVL calculated successfully")

	return totalValue.StringFixed(roundingDecimals), nil
}

// GetConfig returns the configuration for the Bulla Exchange adapter.
func (b *BullaLPPriceProvider) GetConfig(ctx context.Context, address string, client bind.ContractBackend) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	contract, err := sc.NewBulla(common.HexToAddress(address), client)
	if err != nil {
		err = fmt.Errorf("failed to instantiate Bulla smart contract, %v", err)
		return nil, err
	}

	bc := BullaConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	// Token0
	addr, err := contract.Token0(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token0 address for Bulla pool %s, %v", address, err)
		return nil, err
	}
	bc.Token0 = strings.ToLower(addr.Hex())

	// Token1
	addr, err = contract.Token1(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token1 address for Bulla pool %s, %v", address, err)
		return nil, err
	}
	bc.Token1 = strings.ToLower(addr.Hex())

	// Decimals
	decimals, err := contract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for LP token %s, %v", address, err)
		return nil, err
	}
	bc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(bc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (b *BullaLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	b.block = block
	if prices != nil {
		b.priceMap = prices
	}
}

// TVLBreakdown returns the breakdown of TVL by underlying tokens.
// TODO: Implement TVL breakdown for Bulla protocol
func (b *BullaLPPriceProvider) TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error) {
	return nil, ErrTVLBreakdownNotImplemented
}

// Internal Helper methods not able to be called except in this file

// totalValue calculates the total value of assets in the pool
func (b *BullaLPPriceProvider) totalValue(ctx context.Context) (decimal.Decimal, error) {
	total0, total1, err := b.getPoolBalances(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	price0, err := b.getPrice(b.config.Token0)
	if err != nil {
		return decimal.Zero, err
	}

	price1, err := b.getPrice(b.config.Token1)
	if err != nil {
		return decimal.Zero, err
	}

	amount0Decimal := NormalizeAmount(total0, price0.Decimals)
	amount1Decimal := NormalizeAmount(total1, price1.Decimals)

	value0 := amount0Decimal.Mul(price0.Price)
	value1 := amount1Decimal.Mul(price1.Price)

	totalValue := value0.Add(value1)

	return totalValue, nil
}

// getPrice retrieves the price for a given token
func (b *BullaLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := b.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		b.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of the LP token
func (b *BullaLPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: b.block,
	}

	totalSupply, err := b.contract.TotalSupply(opts)
	if err != nil {
		b.logger.Error().Msgf("failed to obtain total supply for Bulla pool %s, %v", b.address.String(), err)
		return nil, fmt.Errorf("failed to get Bulla total supply, err: %w", err)
	}

	return totalSupply, nil
}

// getPoolBalances fetches the base and quote token balances in the pool
func (b *BullaLPPriceProvider) getPoolBalances(ctx context.Context) (*big.Int, *big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: b.block,
	}

	totalAmounts, err := b.contract.GetTotalAmounts(opts)
	if err != nil {
		b.logger.Error().Msgf("failed to obtain total amounts for Bulla pool %s, %v", b.address.String(), err)
		return nil, nil, fmt.Errorf("failed to get Bulla total amounts, err: %w", err)
	}

	return totalAmounts.Total0, totalAmounts.Total1, nil
}
