package protocols

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/infrared-dao/protocols/internal/sc"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

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
func (k *KodiakLPPriceProvider) LPTokenPrice(ctx context.Context) (uint64, error) {
	k.logger.Info().Msg("Calculating LP token price")

	// Fetch total supply
	totalSupply, err := k.getTotalSupply(ctx)
	if err != nil {
		return 0, err
	}

	// Avoid division by zero
	if totalSupply.Sign() == 0 {
		err := errors.New("totalSupply is zero, cannot calculate LP token price")
		k.logger.Error().Err(err).Msg("Invalid totalSupply")
		return 0, err
	}

	// Fetch underlying balances
	amount0, amount1, err := k.getUnderlyingBalances(ctx)
	if err != nil {
		return 0, err
	}

	// Calculate total value in USD using token prices
	amount0Decimal := decimal.NewFromBigInt(amount0, 0)
	amount1Decimal := decimal.NewFromBigInt(amount1, 0)

	totalValue := amount0Decimal.Mul(k.tokenPrices[0]).Add(amount1Decimal.Mul(k.tokenPrices[1]))

	// Calculate price per LP token
	totalSupplyDecimal := decimal.NewFromBigInt(totalSupply, 0)
	pricePerToken := totalValue.Div(totalSupplyDecimal)

	// Convert to USD cents
	priceInCents := pricePerToken.Mul(decimal.NewFromInt(100)).Round(0).BigInt()

	k.logger.Info().
		Str("totalValue", totalValue.String()).
		Str("totalSupply", totalSupplyDecimal.String()).
		Str("pricePerToken", pricePerToken.String()).
		Msg("LP token price calculated successfully")

	return priceInCents.Uint64(), nil
}

// TVL returns the Total Value Locked in the protocol in USD cents (1 USD = 100 cents).
func (k *KodiakLPPriceProvider) TVL(ctx context.Context) (uint64, error) {
	k.logger.Info().Msg("Calculating TVL")

	// Fetch underlying balances
	amount0, amount1, err := k.getUnderlyingBalances(ctx)
	if err != nil {
		return 0, err
	}

	// Calculate total value in USD using token prices
	amount0Decimal := decimal.NewFromBigInt(amount0, 0)
	amount1Decimal := decimal.NewFromBigInt(amount1, 0)

	totalValue := amount0Decimal.Mul(k.tokenPrices[0]).Add(amount1Decimal.Mul(k.tokenPrices[1]))

	// Divide by 1e18 to normalize the value
	totalValue = totalValue.Div(decimal.NewFromInt(1e18))

	// Convert to USD cents
	totalValueInCents := totalValue.Mul(decimal.NewFromInt(100)).Round(0).BigInt()

	k.logger.Info().
		Str("totalValue", totalValue.String()).
		Msg("TVL calculated successfully")

	return totalValueInCents.Uint64(), nil
}
