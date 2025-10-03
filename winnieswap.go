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

// enforce interface adherence
var _ Protocol = &WinnieSwapLPPriceProvider{}

// Define core types for the WinnieSwap Adapter

type WinnieSwapConfig struct {
	Token0      string `json:"token0"`
	Token1      string `json:"token1"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// WinnieSwapLPPriceProvider defines the provider for WinnieSwap LP price and TVL.
type WinnieSwapLPPriceProvider struct {
	address     common.Address
	block       *big.Int
	priceMap    map[string]Price
	logger      zerolog.Logger
	configBytes []byte
	config      *WinnieSwapConfig
	contract    *sc.StickyVault
}

// NewWinnieSwapLPPriceProvider creates a new instance of the WinnieSwapLPPriceProvider.
func NewWinnieSwapLPPriceProvider(
	address common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *WinnieSwapLPPriceProvider {
	w := &WinnieSwapLPPriceProvider{
		address:     address,
		block:       block,
		priceMap:    prices,
		logger:      logger,
		configBytes: config,
	}
	return w
}

// Initialize checks the configuration/data provided and instantiates the StickyVault smart contract.
func (w *WinnieSwapLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	var err error

	w.config = &WinnieSwapConfig{}
	err = json.Unmarshal(w.configBytes, w.config)
	if err != nil {
		w.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}
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

	// initialize StickyVault contract
	vault, err := sc.NewStickyVault(w.address, client)
	if err != nil {
		return err
	}
	w.contract = vault

	return nil
}

// LPTokenPrice returns the current price of LP token in USD
func (w *WinnieSwapLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
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

	w.logger.Debug().
		Str("totalValue", totalValue.String()).
		Str("totalSupply", totalSupplyDecimal.String()).
		Str("pricePerToken", pricePerToken.String()).
		Msg("LP token price calculated successfully")

	return pricePerToken.StringFixed(roundingDecimals), nil
}

// TVL returns the Total Value Locked in the vault as USD
func (w *WinnieSwapLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := w.totalValue(ctx)
	if err != nil {
		return "", err
	}

	w.logger.Debug().
		Str("totalValue", totalValue.String()).
		Msg("TVL calculated successfully")

	return totalValue.StringFixed(roundingDecimals), nil
}

func (w *WinnieSwapLPPriceProvider) GetConfig(ctx context.Context, address string, client *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	// initialize StickyVault contract
	contract, err := sc.NewStickyVault(common.HexToAddress(address), client)
	if err != nil {
		w.logger.Error().Err(err).Msg("failed to instantiate StickyVault contract")
		return nil, err
	}

	wc := WinnieSwapConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	// token0
	addr, err := contract.Token0(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token0 address for WinnieSwap vault %s, %v", address, err)
		return nil, err
	}
	wc.Token0 = strings.ToLower(addr.Hex())

	// token1
	addr, err = contract.Token1(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token1 address for WinnieSwap vault %s, %v", address, err)
		return nil, err
	}
	wc.Token1 = strings.ToLower(addr.Hex())

	// decimals
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

func (w *WinnieSwapLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	w.block = block
	if prices != nil {
		w.priceMap = prices
	}
}

// Internal Helper methods not able to be called except in this file

func (w *WinnieSwapLPPriceProvider) totalValue(ctx context.Context) (decimal.Decimal, error) {
	var err error

	// Fetch internal balances
	amount0, amount1, err := w.getBalances(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	price0, err := w.getPrice(w.config.Token0)
	if err != nil {
		return decimal.Zero, err
	}
	amount0Decimal := NormalizeAmount(amount0, price0.Decimals)

	price1, err := w.getPrice(w.config.Token1)
	if err != nil {
		return decimal.Zero, err
	}
	amount1Decimal := NormalizeAmount(amount1, price1.Decimals)
	totalValue := amount0Decimal.Mul(price0.Price).Add(amount1Decimal.Mul(price1.Price))
	return totalValue, nil
}

func (w *WinnieSwapLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := w.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		w.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of the LP token.
func (w *WinnieSwapLPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: w.block,
	}

	ts, err := w.contract.TotalSupply(opts)
	if err != nil {
		w.logger.Error().Msgf("failed to obtain total supply for WinnieSwap vault %s, %v",
			w.address.String(), err)
		return nil, fmt.Errorf("failed to get WinnieSwap vault total supply, err: %w", err)
	}

	return ts, err
}

// getBalances fetches the underlying token balances.
func (w *WinnieSwapLPPriceProvider) getBalances(ctx context.Context) (*big.Int, *big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: w.block,
	}

	balances, err := w.contract.GetUnderlyingBalances(opts)
	if err != nil {
		w.logger.Error().Msgf("failed to obtain balances for WinnieSwap vault %s, %v",
			w.address.String(), err)
		return nil, nil, fmt.Errorf("failed to get WinnieSwap vault balances, err: %w", err)
	}
	return balances.Amount0Current, balances.Amount1Current, err
}
