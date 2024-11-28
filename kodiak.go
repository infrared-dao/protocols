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

type Price struct {
	TokenName string
	Decimals  uint
	Price     decimal.Decimal
}

type KodiakConfig struct {
	Token0      string `json:"token0"`
	Token1      string `json:"token1"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// KodiakLPPriceProvider defines the provider for Kodiak LP price and TVL.
type KodiakLPPriceProvider struct {
	address     common.Address
	logger      zerolog.Logger
	priceMap    map[string]Price
	prices      []Price
	configBytes []byte
	config      *KodiakConfig
	contract    *sc.KodiakV1
}

// NewKodiakLPPriceProvider creates a new instance of the KodiakLPPriceProvider.
func NewKodiakLPPriceProvider(address common.Address, prices map[string]Price, logger zerolog.Logger, config []byte) *KodiakLPPriceProvider {
	k := &KodiakLPPriceProvider{
		address:     address,
		logger:      logger,
		priceMap:    prices,
		configBytes: config,
	}
	k.prices = make([]Price, 2)
	return k
}

// Initialize instantiates the KodiakV1 smart contract.
func (k *KodiakLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	k.logger.Info().Msg("Initializing KodiakLPPriceProvider")
	var err error
	if len(k.prices) != 2 {
		err = fmt.Errorf("invalid price array, should have 2 elements, got %d instead", len(k.prices))
		k.logger.Error().Msg(err.Error())
		return err
	}

	k.config = &KodiakConfig{}
	err = json.Unmarshal(k.configBytes, k.config)
	if err != nil {
		k.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}
	err = k.setPrice(0)
	if err != nil {
		return err
	}
	err = k.setPrice(1)
	if err != nil {
		return err
	}

	k.contract, err = sc.NewKodiakV1(k.address, client)
	if err != nil {
		k.logger.Error().Err(err).Msg("failed to instatiate KodiakV1 smart contract")
		return err
	}
	return nil
}

func (k *KodiakLPPriceProvider) setPrice(idx int) error {
	var err error
	if len(k.prices) != 2 {
		err = fmt.Errorf("invalid price array, should have 2 elements, got %d instead", len(k.prices))
		k.logger.Error().Msg(err.Error())
		return err
	}
	switch idx {
	case 0:
		price, ok := k.priceMap[k.config.Token0]
		if !ok {
			err = fmt.Errorf("no price for token0 ('%s')", k.config.Token0)
			k.logger.Error().Msg(err.Error())
			return err
		}
		k.prices[0] = price
	case 1:
		price, ok := k.priceMap[k.config.Token1]
		if !ok {
			err = fmt.Errorf("no price for token1 ('%s')", k.config.Token1)
			k.logger.Error().Msg(err.Error())
			return err
		}
		k.prices[1] = price
	default:
		err = fmt.Errorf("invalid price index: %d", idx)
		k.logger.Error().Msg(err.Error())
		return err
	}
	return nil
}

// getTotalSupply fetches the total supply of the LP token.
func (k *KodiakLPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
	ts, err := k.contract.TotalSupply(opts)
	if err != nil {
		k.logger.Error().Msgf("failed to obtain total supply for kodiak vault %s, %v", k.address.String(), err)
		return nil, err
	}

	return ts, err
}

// getUnderlyingBalances fetches the underlying token balances.
func (k *KodiakLPPriceProvider) getUnderlyingBalances(ctx context.Context) (*big.Int, *big.Int, error) {
	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
	ubs, err := k.contract.GetUnderlyingBalances(opts)
	if err != nil {
		k.logger.Error().Msgf("failed to obtain underlying balances for kodiak vault %s, %v", k.address.String(), err)
		return nil, nil, err
	}
	return ubs.Amount0Current, ubs.Amount1Current, err
}

// LPTokenPrice returns the current price of the protocol's LP token in USD cents (1 USD = 100 cents).
func (k *KodiakLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	// Fetch total supply
	totalSupply, err := k.getTotalSupply(ctx)
	if err != nil {
		return "", err
	}

	// Avoid division by zero
	if totalSupply.Sign() == 0 {
		err := errors.New("totalSupply is zero, cannot calculate LP token price")
		k.logger.Error().Err(err).Msg("Invalid totalSupply")
		return "", err
	}

	totalValue, err := k.getTotalValue(ctx)
	if err != nil {
		return "", err
	}

	totalSupplyDecimal := NormalizeAmount(totalSupply, k.config.LPTDecimals)
	pricePerToken := totalValue.Div(totalSupplyDecimal)

	k.logger.Info().
		Str("totalValue", totalValue.String()).
		Str("totalSupply", totalSupplyDecimal.String()).
		Str("pricePerToken", pricePerToken.String()).
		Msg("LP token price calculated successfully")

	return pricePerToken.StringFixed(8), nil
}

// TVL returns the Total Value Locked in the protocol in USD cents (1 USD = 100 cents).
func (k *KodiakLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := k.getTotalValue(ctx)
	if err != nil {
		return "", err
	}

	k.logger.Info().
		Str("totalValue", totalValue.String()).
		Msg("TVL calculated successfully")

	return totalValue.StringFixed(8), nil
}

func (k *KodiakLPPriceProvider) GetConfig(ctx context.Context, address string, client *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}
	contract, err := sc.NewKodiakV1(common.HexToAddress(address), client)
	if err != nil {
		err = fmt.Errorf("failed to instatiate KodiakV1 smart contract, %v", err)
		return nil, err
	}

	kc := KodiakConfig{}
	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	// token0
	addr, err := contract.Token0(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token0 address for kodiak vault %s, %v", address, err)
		return nil, err
	}
	kc.Token0 = strings.ToLower(addr.Hex())

	// token1
	addr, err = contract.Token1(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token1 address for kodiak vault %s, %v", address, err)
		return nil, err
	}
	kc.Token1 = strings.ToLower(addr.Hex())

	// decimals
	decimals, err := contract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for LP token %s, %v", address, err)
		return nil, err
	}
	kc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(kc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func NormalizeAmount(amount *big.Int, decimals uint) decimal.Decimal {
	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	return decimal.NewFromBigInt(amount, 0).Div(decimal.NewFromBigInt(divisor, 0))
}
func (k *KodiakLPPriceProvider) getTotalValue(ctx context.Context) (decimal.Decimal, error) {
	var err error
	if len(k.prices) != 2 {
		err = fmt.Errorf("invalid price array, should have 2 elements, got %d instead", len(k.prices))
		k.logger.Error().Msg(err.Error())
		return decimal.Zero, err
	}

	// Fetch total supply
	totalSupply, err := k.getTotalSupply(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	// Avoid division by zero
	if totalSupply.Sign() == 0 {
		err := errors.New("totalSupply is zero, cannot calculate LP token price")
		k.logger.Error().Err(err).Msg("Invalid totalSupply")
		return decimal.Zero, err
	}

	// Fetch underlying balances
	amount0, amount1, err := k.getUnderlyingBalances(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	amount0Decimal := NormalizeAmount(amount0, k.prices[0].Decimals)
	amount1Decimal := NormalizeAmount(amount1, k.prices[1].Decimals)
	totalValue := amount0Decimal.Mul(k.prices[0].Price)
	totalValue.Add(amount1Decimal.Mul(k.prices[1].Price))
	totalSupplyDecimal := NormalizeAmount(totalSupply, k.config.LPTDecimals)
	return totalSupplyDecimal, nil
}
