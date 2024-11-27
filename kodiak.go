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

type KodiakConfig struct {
	Token0      string `json:"token0"`
	Token1      string `json:"token1"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// KodiakLPPriceProvider defines the provider for Kodiak LP price and TVL.
type KodiakLPPriceProvider struct {
	address     common.Address
	logger      zerolog.Logger
	tokenPrices []decimal.Decimal // Prices of the two underlying tokens
	contract    *sc.KodiakV1
}

// NewKodiakLPPriceProvider creates a new instance of the KodiakLPPriceProvider.
func NewKodiakLPPriceProvider(address common.Address, prices []decimal.Decimal, logger zerolog.Logger) *KodiakLPPriceProvider {
	return &KodiakLPPriceProvider{
		address:     address,
		logger:      logger,
		tokenPrices: prices,
	}
}

// Initialize instantiates the KodiakV1 smart contract.
func (k *KodiakLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	k.logger.Info().Msg("Initializing KodiakLPPriceProvider")
	var err error
	if len(k.tokenPrices) != 2 {
		err = fmt.Errorf("invalid price array, should have 2 elements, got %d instead", len(k.tokenPrices))
		k.logger.Error().Msg(err.Error())
		return err
	}

	k.contract, err = sc.NewKodiakV1(k.address, client)
	if err != nil {
		k.logger.Error().Err(err).Msg("failed to instatiate KodiakV1 smart contract")
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
	k.logger.Info().Msg("Calculating LP token price")
	var err error
	if len(k.tokenPrices) != 2 {
		err = fmt.Errorf("invalid price array, should have 2 elements, got %d instead", len(k.tokenPrices))
		k.logger.Error().Msg(err.Error())
		return "", err
	}

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

	// Fetch underlying balances
	amount0, amount1, err := k.getUnderlyingBalances(ctx)
	if err != nil {
		return "", err
	}

	// TODO:use decimals provided by caller to calculate the correct price [Issue DEV-866] (@kuma)
	// Calculate total value in USD using token prices
	amount0Decimal := decimal.NewFromBigInt(amount0, 0)
	amount1Decimal := decimal.NewFromBigInt(amount1, 0)

	totalValue := amount0Decimal.Mul(k.tokenPrices[0]).Add(amount1Decimal.Mul(k.tokenPrices[1]))

	// Calculate price per LP token
	totalSupplyDecimal := decimal.NewFromBigInt(totalSupply, 0)
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
	k.logger.Info().Msg("Calculating TVL")
	var err error
	if len(k.tokenPrices) != 2 {
		err = fmt.Errorf("invalid price array, should have 2 elements, got %d instead", len(k.tokenPrices))
		k.logger.Error().Msg(err.Error())
		return "", err
	}

	// Fetch underlying balances
	amount0, amount1, err := k.getUnderlyingBalances(ctx)
	if err != nil {
		return "", err
	}

	// Calculate total value in USD using token prices
	amount0Decimal := decimal.NewFromBigInt(amount0, 0)
	amount1Decimal := decimal.NewFromBigInt(amount1, 0)

	totalValue := amount0Decimal.Mul(k.tokenPrices[0]).Add(amount1Decimal.Mul(k.tokenPrices[1]))

	// TODO: do not rely on 18 decimals, use decimals provided by caller [Issue DEV-866] (@kuma)
	// Divide by 1e18 to normalize the value
	totalValue = totalValue.Div(decimal.NewFromInt(1e18))

	k.logger.Info().
		Str("totalValue", totalValue.String()).
		Msg("TVL calculated successfully")

	return totalValue.StringFixed(8), nil
}

func (k *KodiakLPPriceProvider) GetConfig(ctx context.Context) ([]byte, error) {
	kc := KodiakConfig{}
	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	// token0
	addr, err := k.contract.Token0(opts)
	if err != nil {
		k.logger.Error().Msgf("failed to obtain token0 address for kodiak vault %s, %v", k.address.String(), err)
		return nil, err
	}
	kc.Token0 = strings.ToLower(addr.Hex())

	// token1
	addr, err = k.contract.Token1(opts)
	if err != nil {
		k.logger.Error().Msgf("failed to obtain token1 address for kodiak vault %s, %v", k.address.String(), err)
		return nil, err
	}
	kc.Token1 = strings.ToLower(addr.Hex())

	// decimals
	decimals, err := k.contract.Decimals(opts)
	if err != nil {
		k.logger.Error().Msgf("failed to obtain number of decmals for LP token %s, %v", k.address.String(), err)
		return nil, err
	}
	kc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(kc)
	if err != nil {
		return nil, err
	}

	return body, nil
}
