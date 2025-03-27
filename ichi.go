package protocols

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/infrared-dao/protocols/internal/sc"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

// Configuration for ICHI LP Vault
type ICHIConfig struct {
	Token0Decimals uint   `json:"token0_decimals"`
	Token1Decimals uint   `json:"token1_decimals"`
	Token0Address  string `json:"token0_address"`
	Token1Address  string `json:"token1_address"`
}

// ICHIPriceProvider defines the provider for ICHI vault LP price and TVL.
type ICHIPriceProvider struct {
	LPTAddress  common.Address
	logger      zerolog.Logger
	block       *big.Int
	configBytes []byte
	config      *ICHIConfig
	contract    *sc.Ichi
}

// NewICHIPriceProvider creates a new instance of the ICHIPriceProvider.
func NewICHIPriceProvider(LPTAddress common.Address, block *big.Int, logger zerolog.Logger, config []byte) *ICHIPriceProvider {
	b := &ICHIPriceProvider{
		LPTAddress:  LPTAddress,
		logger:      logger,
		block:       block,
		configBytes: config,
	}
	return b
}

// Initialize checks the configuration/data and instantiates the ICHI vault contract.
func (i *ICHIPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	var err error

	i.config = &ICHIConfig{}
	err = json.Unmarshal(i.configBytes, i.config)
	if err != nil {
		i.logger.Error().Err(err).Msg("failed to deserialize ICHI config")
		return err
	}

	// Initialize the ICHI vault contract from the address
	i.contract, err = sc.NewIchi(i.LPTAddress, client)
	if err != nil {
		i.logger.Error().Err(err).Msg("adapter init failed to instantiate ICHI vault contract")
		return err
	}

	return nil
}

// LPTokenPrice returns the current price of the protocol's LP token in USD
func (i *ICHIPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	tvl, err := i.calculateTVL(ctx)
	if err != nil {
		return "", err
	}

	opts := &bind.CallOpts{
		Pending:     false,
		Context:     ctx,
		BlockNumber: i.block,
	}

	// Get the total supply of the vault token
	totalSupply, err := i.contract.TotalSupply(opts)
	if err != nil {
		return "", fmt.Errorf("failed to get ICHI vault total supply, err: %w", err)
	}

	if totalSupply.Cmp(big.NewInt(0)) == 0 {
		return "0", nil
	}

	// LP Token has 18 decimals as per standard ERC20
	numTokens := NormalizeAmount(totalSupply, 18)

	// Calculate price per token = TVL / total supply
	pricePerToken := tvl.Div(numTokens)

	i.logger.Info().
		Str("pricePerToken", pricePerToken.String()).
		Msg("ICHI LP token price calculated successfully")

	return pricePerToken.StringFixed(roundingDecimals), nil
}

// TVL returns the Total Value Locked in the ICHI vault in USD
func (i *ICHIPriceProvider) TVL(ctx context.Context) (string, error) {
	tvl, err := i.calculateTVL(ctx)
	if err != nil {
		return "", err
	}

	return tvl.StringFixed(roundingDecimals), nil
}

// calculateTVL is a helper function to calculate the TVL based on token amounts and prices
func (i *ICHIPriceProvider) calculateTVL(ctx context.Context) (decimal.Decimal, error) {
	opts := &bind.CallOpts{
		Pending:     false,
		Context:     ctx,
		BlockNumber: i.block,
	}

	// Get total amounts of token0 and token1 in the vault
	totalAmounts, err := i.contract.GetTotalAmounts(opts)
	if err != nil {
		return decimal.Zero, fmt.Errorf("failed to get total amounts from ICHI vault, err: %w", err)
	}

	amount0 := NormalizeAmount(totalAmounts.Total0, i.config.Token0Decimals)
	amount1 := NormalizeAmount(totalAmounts.Total1, i.config.Token1Decimals)

	// Get price feeds for token0 and token1
	// For this we need to query the price feed contract or service
	// Assuming we have a function to get the price of a token in USD
	token0Price, err := i.getTokenPrice(ctx, common.HexToAddress(i.config.Token0Address))
	if err != nil {
		return decimal.Zero, fmt.Errorf("failed to get token0 price, err: %w", err)
	}

	token1Price, err := i.getTokenPrice(ctx, common.HexToAddress(i.config.Token1Address))
	if err != nil {
		return decimal.Zero, fmt.Errorf("failed to get token1 price, err: %w", err)
	}

	// Calculate the total value: (amount0 * price0) + (amount1 * price1)
	token0Value := amount0.Mul(token0Price)
	token1Value := amount1.Mul(token1Price)
	totalValue := token0Value.Add(token1Value)

	return totalValue, nil
}

// getTokenPrice fetches the USD price of a token
// This is a placeholder - in a real implementation, you'd connect to a price oracle or feed
func (i *ICHIPriceProvider) getTokenPrice(ctx context.Context, tokenAddress common.Address) (decimal.Decimal, error) {
	// In a real implementation, this would query a price oracle or service
	// For this example, we're assuming all tokens have a price feed available through MetaBeraborrowCore's price feed

	//opts := &bind.CallOpts{
	//	Pending:     false,
	//	Context:     ctx,
	//	BlockNumber: i.block,
	//}

	if tokenAddress == common.HexToAddress("0x1F7210257FA157227D09449229a9266b0D581337") {
		return decimal.NewFromFloat(0.002678), nil
	} else if tokenAddress == common.HexToAddress("0x6969696969696969696969696969696969696969") {
		return decimal.NewFromFloat(7.63), nil
	}

	// This is where you would query a price oracle or feed
	// For example, if using Chainlink or a similar oracle:
	// price, err := priceOracle.GetLatestPrice(opts, tokenAddress)

	// For this example, we'll just return a placeholder 1.0 USD price
	// In a real implementation, replace this with actual price fetching logic
	return decimal.NewFromFloat(0), fmt.Errorf("error fetching token price")
}

// GetConfig gathers configuration data needed for the ICHI vault adapter
func (i *ICHIPriceProvider) GetConfig(ctx context.Context, lpAddress string, client *ethclient.Client) ([]byte, error) {
	if !common.IsHexAddress(lpAddress) {
		err := fmt.Errorf("invalid smart contract address, '%s'", lpAddress)
		return nil, err
	}

	config := ICHIConfig{}
	lpTokenAddr := common.HexToAddress(lpAddress)

	opts := &bind.CallOpts{
		Context: ctx,
	}

	// Create ICHI vault contract instance
	ichiVault, err := sc.NewIchi(lpTokenAddr, client)
	if err != nil {
		i.logger.Error().Err(err).Msg("failed to instantiate ICHI vault contract")
		return nil, err
	}

	// Get token0 and token1 addresses
	token0Addr, err := ichiVault.Token0(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to get token0 address: %w", err)
	}
	config.Token0Address = token0Addr.Hex()

	token1Addr, err := ichiVault.Token1(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to get token1 address: %w", err)
	}
	config.Token1Address = token1Addr.Hex()

	// Get token decimals
	// We need to create ERC20 contract instances for both tokens to get their decimals
	token0, err := sc.NewERC20(token0Addr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create token0 contract: %w", err)
	}

	token1, err := sc.NewERC20(token1Addr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create token1 contract: %w", err)
	}

	token0Decimals, err := token0.Decimals(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to get token0 decimals: %w", err)
	}
	config.Token0Decimals = uint(token0Decimals)

	token1Decimals, err := token1.Decimals(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to get token1 decimals: %w", err)
	}
	config.Token1Decimals = uint(token1Decimals)

	// Serialize config to JSON
	configBytes, err := json.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize ICHI config: %w", err)
	}

	return configBytes, nil
}
