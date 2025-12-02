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

var _ Protocol = &AquaBeraLPPriceProvider{}

type AquaBeraConfig struct {
	Token0      string `json:"token0"`
	Token1      string `json:"token1"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// AquaBeraLPPriceProvider defines the provider for AquaBera LP price and TVL.
type AquaBeraLPPriceProvider struct {
	address     common.Address
	block       *big.Int
	priceMap    map[string]Price
	logger      zerolog.Logger
	configBytes []byte
	config      *AquaBeraConfig
	contract    *sc.AquaBera
}

// NewAquaBeraLPPriceProvider creates a new instance of the AquaBeraLPPriceProvider.
func NewAquaBeraLPPriceProvider(
	address common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *AquaBeraLPPriceProvider {
	a := &AquaBeraLPPriceProvider{
		address:     address,
		block:       block,
		priceMap:    prices,
		logger:      logger,
		configBytes: config,
	}
	return a
}

// Initialize checks the configuration/data provided and instantiates the AquaBera smart contract.
func (a *AquaBeraLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	var err error

	a.config = &AquaBeraConfig{}
	err = json.Unmarshal(a.configBytes, a.config)
	if err != nil {
		a.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}
	_, ok := a.priceMap[a.config.Token0]
	if !ok {
		err = fmt.Errorf("no price data found for token0 (%s)", a.config.Token0)
		a.logger.Error().Msg(err.Error())
		return err
	}
	_, ok = a.priceMap[a.config.Token1]
	if !ok {
		err = fmt.Errorf("no price data found for token1 (%s)", a.config.Token1)
		a.logger.Error().Msg(err.Error())
		return err
	}

	a.contract, err = sc.NewAquaBera(a.address, client)
	if err != nil {
		a.logger.Error().Err(err).Msg("failed to instantiate AquaBera smart contract")
		return err
	}
	return nil
}

// LPTokenPrice returns the current price of the protocol's LP token in USD.
func (a *AquaBeraLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	// Fetch total supply
	totalSupply, err := a.getTotalSupply(ctx)
	if err != nil {
		return "", err
	}

	// Avoid division by zero
	if totalSupply.Sign() == 0 {
		err := errors.New("totalSupply is zero, cannot calculate LP token price")
		a.logger.Error().Err(err).Msg("Invalid totalSupply")
		return "", err
	}

	totalValue, err := a.totalValue(ctx)
	if err != nil {
		return "", err
	}

	totalSupplyDecimal := NormalizeAmount(totalSupply, a.config.LPTDecimals)
	pricePerToken := totalValue.Div(totalSupplyDecimal)

	a.logger.Debug().
		Str("totalValue", totalValue.String()).
		Str("totalSupply", totalSupplyDecimal.String()).
		Str("pricePerToken", pricePerToken.String()).
		Msg("LP token price calculated successfully")

	return pricePerToken.StringFixed(roundingDecimals), nil
}

// TVL returns the Total Value Locked in the protocol in USD.
func (a *AquaBeraLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := a.totalValue(ctx)
	if err != nil {
		return "", err
	}

	a.logger.Debug().
		Str("totalValue", totalValue.String()).
		Msg("TVL calculated successfully")

	return totalValue.StringFixed(roundingDecimals), nil
}

// GetConfig fetches and returns the configuration for the AquaBera protocol.
func (a *AquaBeraLPPriceProvider) GetConfig(ctx context.Context, address string, client *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}
	contract, err := sc.NewAquaBera(common.HexToAddress(address), client)
	if err != nil {
		err = fmt.Errorf("failed to instantiate AquaBera smart contract, %v", err)
		return nil, err
	}

	ac := AquaBeraConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	// token0
	addr, err := contract.Token0(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token0 address for aquabera vault %s, %v", address, err)
		return nil, err
	}
	ac.Token0 = strings.ToLower(addr.Hex())

	// token1
	addr, err = contract.Token1(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token1 address for aquabera vault %s, %v", address, err)
		return nil, err
	}
	ac.Token1 = strings.ToLower(addr.Hex())

	// decimals
	decimals, err := contract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for LP token %s, %v", address, err)
		return nil, err
	}
	ac.LPTDecimals = uint(decimals)

	body, err := json.Marshal(ac)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (a *AquaBeraLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	a.block = block
	if prices != nil {
		a.priceMap = prices
	}
}

// TVLBreakdown returns the breakdown of TVL by underlying tokens.
// TODO: Implement TVL breakdown for AquaBera protocol
func (a *AquaBeraLPPriceProvider) TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error) {
	return nil, ErrTVLBreakdownNotImplemented
}

// Internal Helper methods not able to be called except in this file

// Helper method to calculate the total value of underlying assets
func (a *AquaBeraLPPriceProvider) totalValue(ctx context.Context) (decimal.Decimal, error) {
	var err error

	// Fetch underlying balances
	amount0, amount1, err := a.getUnderlyingBalances(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	price0, err := a.getPrice(a.config.Token0)
	if err != nil {
		return decimal.Zero, err
	}
	amount0Decimal := NormalizeAmount(amount0, price0.Decimals)

	price1, err := a.getPrice(a.config.Token1)
	if err != nil {
		return decimal.Zero, err
	}
	amount1Decimal := NormalizeAmount(amount1, price1.Decimals)
	totalValue := amount0Decimal.Mul(price0.Price).Add(amount1Decimal.Mul(price1.Price))
	return totalValue, nil
}

func (a *AquaBeraLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := a.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		a.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of the LP token.
func (a *AquaBeraLPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: a.block,
	}
	ts, err := a.contract.TotalSupply(opts)
	if err != nil {
		a.logger.Error().Msgf("failed to obtain total supply for aquabera vault %s, %v", a.address.String(), err)
		return nil, fmt.Errorf("failed to get aquabera total supply, err: %w", err)
	}

	return ts, err
}

// getUnderlyingBalances fetches the underlying token balances.
func (a *AquaBeraLPPriceProvider) getUnderlyingBalances(ctx context.Context) (*big.Int, *big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: a.block,
	}
	ubs, err := a.contract.GetTotalAmounts(opts)
	if err != nil {
		a.logger.Error().Msgf("failed to obtain underlying balances for aquabera vault %s, %v", a.address.String(), err)
		return nil, nil, fmt.Errorf("failed to get aquabera underlying balances, err: %w", err)
	}
	return ubs.Total0, ubs.Total1, err
}
