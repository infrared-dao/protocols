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
var _ Protocol = &SteerLPPriceProvider{}

// Define core types for the Steer Adapter

type SteerConfig struct {
	Token0      string `json:"token0"`
	Token1      string `json:"token1"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// SteerLPPriceProvider defines the provider for Steer LP price and TVL.
type SteerLPPriceProvider struct {
	address     common.Address
	block       *big.Int
	priceMap    map[string]Price
	logger      zerolog.Logger
	configBytes []byte
	config      *SteerConfig
	contract    *sc.SteerPool
}

// NewSteerLPPriceProvider creates a new instance of the SteerLPPriceProvider.
func NewSteerLPPriceProvider(
	address common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *SteerLPPriceProvider {
	s := &SteerLPPriceProvider{
		address:     address,
		block:       block,
		priceMap:    prices,
		logger:      logger,
		configBytes: config,
	}
	return s
}

// Initialize checks the configuration/data provided and instantiates the SteerV1 smart contract.
func (s *SteerLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	var err error

	s.config = &SteerConfig{}
	err = json.Unmarshal(s.configBytes, s.config)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}
	_, ok := s.priceMap[s.config.Token0]
	if !ok {
		err = fmt.Errorf("no price data found for token0 (%s)", s.config.Token0)
		s.logger.Error().Msg(err.Error())
		return err
	}
	_, ok = s.priceMap[s.config.Token1]
	if !ok {
		err = fmt.Errorf("no price data found for token1 (%s)", s.config.Token1)
		s.logger.Error().Msg(err.Error())
		return err
	}

	// initialize Steer contract of correct type
	pool, err := sc.NewSteerPool(s.address, client)
	if err != nil {
		return err
	}
	s.contract = pool

	return nil
}

// LPTokenPrice returns the current price of LP token in USD
func (s *SteerLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	// Fetch total supply
	totalSupply, err := s.getTotalSupply(ctx)
	if err != nil {
		return "", err
	}

	// Avoid division by zero
	if totalSupply.Sign() == 0 {
		err := errors.New("totalSupply is zero, cannot calculate LP token price")
		s.logger.Error().Err(err).Msg("Invalid totalSupply")
		return "", err
	}

	totalValue, err := s.totalValue(ctx)
	if err != nil {
		return "", err
	}

	totalSupplyDecimal := NormalizeAmount(totalSupply, s.config.LPTDecimals)
	pricePerToken := totalValue.Div(totalSupplyDecimal)

	s.logger.Debug().
		Str("totalValue", totalValue.String()).
		Str("totalSupply", totalSupplyDecimal.String()).
		Str("pricePerToken", pricePerToken.String()).
		Msg("LP token price calculated successfully")

	return pricePerToken.StringFixed(roundingDecimals), nil
}

// TVL returns the Total Value Locked in the pool as USD
func (s *SteerLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := s.totalValue(ctx)
	if err != nil {
		return "", err
	}

	s.logger.Debug().
		Str("totalValue", totalValue.String()).
		Msg("TVL calculated successfully")

	return totalValue.StringFixed(roundingDecimals), nil
}

func (s *SteerLPPriceProvider) GetConfig(ctx context.Context, address string, client *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	// initialize Steer pool contract
	contract, err := sc.NewSteerPool(common.HexToAddress(address), client)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to instantiate Steer Pool contract")
		return nil, err
	}

	sc := SteerConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	// token0
	addr, err := contract.Token0(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token0 address for Steer pool %s, %v", address, err)
		return nil, err
	}
	sc.Token0 = strings.ToLower(addr.Hex())

	// token1
	addr, err = contract.Token1(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token1 address for Steer pool %s, %v", address, err)
		return nil, err
	}
	sc.Token1 = strings.ToLower(addr.Hex())

	// decimals
	decimals, err := contract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for LP token %s, %v", address, err)
		return nil, err
	}
	sc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(sc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (s *SteerLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	s.block = block
	if prices != nil {
		s.priceMap = prices
	}
}

// TVLBreakdown returns the breakdown of TVL by underlying tokens.
// TODO: Implement TVL breakdown for Steer protocol
func (s *SteerLPPriceProvider) TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error) {
	return nil, errors.New("TVLBreakdown not yet implemented for Steer")
}

// Internal Helper methods not able to be called except in this file

func (s *SteerLPPriceProvider) totalValue(ctx context.Context) (decimal.Decimal, error) {
	var err error

	// Fetch internal balances
	amount0, amount1, err := s.getBalances(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	price0, err := s.getPrice(s.config.Token0)
	if err != nil {
		return decimal.Zero, err
	}
	amount0Decimal := NormalizeAmount(amount0, price0.Decimals)

	price1, err := s.getPrice(s.config.Token1)
	if err != nil {
		return decimal.Zero, err
	}
	amount1Decimal := NormalizeAmount(amount1, price1.Decimals)
	totalValue := amount0Decimal.Mul(price0.Price).Add(amount1Decimal.Mul(price1.Price))
	return totalValue, nil
}

func (s *SteerLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := s.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		s.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of the LP token.
func (s *SteerLPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: s.block,
	}

	ts, err := s.contract.TotalSupply(opts)
	if err != nil {
		s.logger.Error().Msgf("failed to obtain total supply for Steer contract %s, %v",
			s.address.String(), err)
		return nil, fmt.Errorf("failed to get Steer contract total supply, err: %w", err)
	}

	return ts, err
}

// getTotalAmounts fetches the underlying token balances.
func (s *SteerLPPriceProvider) getBalances(ctx context.Context) (*big.Int, *big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: s.block,
	}

	ta, err := s.contract.GetTotalAmounts(opts)
	if err != nil {
		s.logger.Error().Msgf("failed to obtain balances for Steer contract %s, %v",
			s.address.String(), err)
		return nil, nil, fmt.Errorf("failed to get Steer contract balances, err: %w", err)
	}
	return ta.Total0, ta.Total1, err
}
