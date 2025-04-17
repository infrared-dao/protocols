package protocols

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/infrared-dao/protocols/internal/sc"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

var _ Protocol = &WasabeeLPPriceProvider{}

// WasabeeConfig defines the configuration for Wasabee adapter
type WasabeeConfig struct {
	Token0      string `json:"token0"`
	Token1      string `json:"token1"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// WasabeeLPPriceProvider defines the provider for Wasabee LP price and TVL.
type WasabeeLPPriceProvider struct {
	address     common.Address
	block       *big.Int
	priceMap    map[string]Price
	logger      zerolog.Logger
	configBytes []byte
	config      *WasabeeConfig
	contract    *sc.WasabeeVault
}

// NewWasabeeLPPriceProvider creates a new instance of the WasabeeLPPriceProvider.
func NewWasabeeLPPriceProvider(
	address common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *WasabeeLPPriceProvider {
	w := &WasabeeLPPriceProvider{
		address:     address,
		block:       block,
		priceMap:    prices,
		logger:      logger,
		configBytes: config,
	}
	return w
}

// Initialize checks the configuration/data provided and instantiates the Wasabee smart contract.
func (w *WasabeeLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	var err error

	w.config = &WasabeeConfig{}
	err = json.Unmarshal(w.configBytes, w.config)
	if err != nil {
		w.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	// Validate that we have price data for the tokens
	_, ok := w.priceMap[w.config.Token0]
	if !ok {
		err = fmt.Errorf("no price data found for token0 (%s)", w.config.Token0)
		w.logger.Error().Msg(err.Error())
		return err
	}

	_, ok = w.priceMap[w.config.Token1]
	if !ok {
		err = fmt.Errorf("no price data found for token1 (%s)", w.config.Token1)
		w.logger.Error().Msg(err.Error())
		return err
	}

	// Initialize Bulla contract
	w.contract, err = sc.NewWasabeeVault(w.address, client)
	if err != nil {
		w.logger.Error().Err(err).Msg("failed to instantiate Wasabee smart contract")
		return err
	}

	return nil
}

// LPTokenPrice returns the current price of the protocol's LP token in USD.
func (w *WasabeeLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	// Fetch total supply
	totalSupply, err := w.getTotalSupply(ctx)
	if err != nil {
		return "", err
	}

	// Avoid division by zero
	if totalSupply.Sign() == 0 {
		err := errors.New("totalSupply is zero, cannot calculate LP token price")
		w.logger.Error().Err(err).Msg("Invalid totalSupply")
		return "", err
	}

	totalValue, err := w.totalValue(ctx)
	if err != nil {
		return "", err
	}

	totalSupplyDecimal := NormalizeAmount(totalSupply, w.config.LPTDecimals)
	pricePerToken := totalValue.Div(totalSupplyDecimal)

	w.logger.Info().
		Str("totalValue", totalValue.String()).
		Str("totalSupply", totalSupplyDecimal.String()).
		Str("pricePerToken", pricePerToken.String()).
		Msg("LP token price calculated successfully")

	return pricePerToken.StringFixed(roundingDecimals), nil
}

// TVL returns the Total Value Locked in the protocol in USD.
func (w *WasabeeLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := w.totalValue(ctx)
	if err != nil {
		return "", err
	}

	w.logger.Info().
		Str("totalValue", totalValue.String()).
		Msg("TVL calculated successfully")

	return totalValue.StringFixed(roundingDecimals), nil
}

// GetConfig returns the configuration for the Wasabee adapter.
func (w *WasabeeLPPriceProvider) GetConfig(ctx context.Context, address string, client *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	contract, err := sc.NewWasabeeVault(common.HexToAddress(address), client)
	if err != nil {
		err = fmt.Errorf("failed to instantiate Wasabee smart contract, %v", err)
		return nil, err
	}

	wc := WasabeeConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	// Token0
	addr, err := contract.Token0(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token0 address for Wasabee vault %s, %v", address, err)
		return nil, err
	}
	wc.Token0 = strings.ToLower(addr.Hex())

	// Token1
	addr, err = contract.Token1(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token1 address for Wasabee vault %s, %v", address, err)
		return nil, err
	}
	wc.Token1 = strings.ToLower(addr.Hex())

	// Decimals
	decimals, err := contract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for LP token %s, %v", address, err)
		return nil, err
	}
	wc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(wc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (w *WasabeeLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	w.block = block
	if prices != nil {
		w.priceMap = prices
	}
}

// Internal Helper methods not able to be called except in this file

// totalValue calculates the total value of assets in the pool
func (w *WasabeeLPPriceProvider) totalValue(ctx context.Context) (decimal.Decimal, error) {
	total0, total1, err := w.getPoolBalances(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	price0, err := w.getPrice(w.config.Token0)
	if err != nil {
		return decimal.Zero, err
	}

	price1, err := w.getPrice(w.config.Token1)
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
func (w *WasabeeLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := w.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		w.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of the LP token
func (w *WasabeeLPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: w.block,
	}

	totalSupply, err := w.contract.TotalSupply(opts)
	if err != nil {
		w.logger.Error().Msgf("failed to obtain total supply for Wasabee pool %s, %v", w.address.String(), err)
		return nil, fmt.Errorf("failed to get Wasabee total supply, err: %w", err)
	}

	return totalSupply, nil
}

// getPoolBalances fetches the base and quote token balances in the pool
func (w *WasabeeLPPriceProvider) getPoolBalances(ctx context.Context) (*big.Int, *big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: w.block,
	}

	totalAmounts, err := w.contract.GetTotalAmounts(opts)
	if err != nil {
		w.logger.Error().Msgf("failed to obtain total amounts for Wasabee pool %s, %v", w.address.String(), err)
		return nil, nil, fmt.Errorf("failed to get Wasabee total amounts, err: %w", err)
	}

	return totalAmounts.Total0, totalAmounts.Total1, nil
}
