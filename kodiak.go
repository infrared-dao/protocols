package protocols

import (
	"context"
	"errors"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

// KodiakLPPriceProvider defines the provider for Kodiak LP price and TVL.
type KodiakLPPriceProvider struct {
	address     common.Address
	abiPath     string
	abi         abi.ABI
	logger      zerolog.Logger
	tokenPrices [2]decimal.Decimal // Prices of the two underlying tokens
}

// NewKodiakLPPriceProvider creates a new instance of the KodiakLPPriceProvider.
func NewKodiakLPPriceProvider(address common.Address, abiPath string, prices [2]decimal.Decimal, logger zerolog.Logger) *KodiakLPPriceProvider {
	return &KodiakLPPriceProvider{
		address:     address,
		abiPath:     abiPath,
		logger:      logger,
		tokenPrices: prices,
	}
}

// Initialize loads the ABI file and prepares the provider.
func (k *KodiakLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	k.logger.Info().Msg("Initializing KodiakLPPriceProvider")

	// Load ABI file
	abiData, err := os.ReadFile(k.abiPath)
	if err != nil {
		k.logger.Error().Err(err).Msg("Failed to read ABI file")
		return err
	}

	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(string(abiData)))
	if err != nil {
		k.logger.Error().Err(err).Msg("Failed to parse ABI")
		return err
	}

	k.abi = parsedABI
	k.logger.Info().Msg("KodiakLPPriceProvider initialized successfully")
	return nil
}

// getTotalSupply fetches the total supply of the LP token.
func (k *KodiakLPPriceProvider) getTotalSupply(ctx context.Context, client *ethclient.Client) (*big.Int, error) {
	// Call totalSupply
	totalSupplyData, err := k.abi.Pack("totalSupply")
	if err != nil {
		k.logger.Error().Err(err).Msg("Failed to pack totalSupply call")
		return nil, err
	}

	totalSupplyRes, err := client.CallContract(ctx, ethereum.CallMsg{
		To:   &k.address,
		Data: totalSupplyData,
	}, nil)
	if err != nil {
		k.logger.Error().Err(err).Msg("Failed to call totalSupply")
		return nil, err
	}

	var totalSupply *big.Int
	if err := k.abi.UnpackIntoInterface(&totalSupply, "totalSupply", totalSupplyRes); err != nil {
		k.logger.Error().Err(err).Msg("Failed to unpack totalSupply result")
		return nil, err
	}

	return totalSupply, nil
}

// getUnderlyingBalances fetches the current token balances.
func (k *KodiakLPPriceProvider) getUnderlyingBalances(ctx context.Context, client *ethclient.Client) (amount0, amount1 *big.Int, err error) {
	// Call getUnderlyingBalances
	getUnderlyingBalancesData, err := k.abi.Pack("getUnderlyingBalances")
	if err != nil {
		k.logger.Error().Err(err).Msg("Failed to pack getUnderlyingBalances call")
		return nil, nil, err
	}

	underlyingBalancesRes, err := client.CallContract(ctx, ethereum.CallMsg{
		To:   &k.address,
		Data: getUnderlyingBalancesData,
	}, nil)
	if err != nil {
		k.logger.Error().Err(err).Msg("Failed to call getUnderlyingBalances")
		return nil, nil, err
	}

	var amount0Result, amount1Result *big.Int
	if err := k.abi.UnpackIntoInterface(&[]interface{}{&amount0Result, &amount1Result}, "getUnderlyingBalances", underlyingBalancesRes); err != nil {
		k.logger.Error().Err(err).Msg("Failed to unpack getUnderlyingBalances result")
		return nil, nil, err
	}

	k.logger.Info().Msgf("underlying balance 0: %v", amount0Result)
	k.logger.Info().Msgf("underlying balance 1: %v", amount1Result)
	return amount0Result, amount1Result, nil
}

// LPTokenPrice returns the current price of the protocol's LP token in USD cents (1 USD = 100 cents).
func (k *KodiakLPPriceProvider) LPTokenPrice(ctx context.Context, client *ethclient.Client) (uint64, error) {
	k.logger.Info().Msg("Calculating LP token price")

	// Fetch total supply
	totalSupply, err := k.getTotalSupply(ctx, client)
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
	amount0, amount1, err := k.getUnderlyingBalances(ctx, client)
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
func (k *KodiakLPPriceProvider) TVL(ctx context.Context, client *ethclient.Client) (uint64, error) {
	k.logger.Info().Msg("Calculating TVL")

	// Fetch underlying balances
	amount0, amount1, err := k.getUnderlyingBalances(ctx, client)
	if err != nil {
		return 0, err
	}

	// Calculate total value in USD using token prices
	amount0Decimal := decimal.NewFromBigInt(amount0, 0)
	amount1Decimal := decimal.NewFromBigInt(amount1, 0)
	totalValue := amount0Decimal.Mul(k.tokenPrices[0]).Add(amount1Decimal.Mul(k.tokenPrices[1]))

	// Convert to USD cents
	totalValueInCents := totalValue.Mul(decimal.NewFromInt(100)).Round(0).BigInt()

	k.logger.Info().
		Str("amount0Current", amount0Decimal.String()).
		Str("amount1Current", amount1Decimal.String()).
		Str("totalValue", totalValue.String()).
		Msg("TVL calculated successfully")

	return totalValueInCents.Uint64(), nil
}
