package protocols

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"

	"github.com/infrared-dao/protocols/internal/sc"
)

// BurrBearConfig defines the configuration for a BurrBear pool
type BurrBearConfig struct {
	PoolId      [32]byte `json:"pool_id"`
	LPTDecimals uint     `json:"lpt_decimals"`
}

// BurrBearLPPriceProvider defines the provider for BurrBear LP price and Pool TVL.
type BurrBearLPPriceProvider struct {
	vaultAddress common.Address
	poolAddress  common.Address
	logger       zerolog.Logger
	priceMap     map[string]Price
	configBytes  []byte
	config       *BurrBearConfig
	vault        *sc.BalancerVault
	lpPool       *sc.BalancerBasePool
}

// NewBurrBearLPPriceProvider creates a new instance of the BurrBearLPPriceProvider.
func NewBurrBearLPPriceProvider(vaultAddress common.Address, poolAddress common.Address, prices map[string]Price, logger zerolog.Logger, config []byte) *BurrBearLPPriceProvider {
	return &BurrBearLPPriceProvider{
		vaultAddress: vaultAddress,
		poolAddress:  poolAddress,
		priceMap:     prices,
		logger:       logger,
		configBytes:  config,
	}
}

// Initialize checks the configuration and instantiates the contracts
func (b *BurrBearLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	// Parse config
	config := &BurrBearConfig{}
	if err := json.Unmarshal(b.configBytes, config); err != nil {
		b.logger.Error().Err(err).Msg("failed to unmarshal config")
		return err
	}
	b.config = config

	// Create Balancer vault contract instance
	vault, err := sc.NewBalancerVault(b.vaultAddress, client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instantiate Balancer vault contract")
		return err
	}
	b.vault = vault

	// Create Balancer pool contract instance
	pool, err := sc.NewBalancerBasePool(b.poolAddress, client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instantiate Balancer pool contract")
		return err
	}
	b.lpPool = pool

	return nil
}

// LPTokenPrice returns the current price of the protocol's LP token in USD.
func (b *BurrBearLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	// Using GetActualSupply because this is how much is circulating for pools which lock up some LP tokens
	totalSupply, err := b.lpPool.BalancerBasePoolCaller.GetActualSupply(&bind.CallOpts{Context: ctx, Pending: false})
	if err != nil {
		return "", fmt.Errorf("failed to get total supply: %w", err)
	}

	// Avoid division by zero
	if totalSupply.Sign() == 0 {
		err := errors.New("totalSupply is zero, cannot calculate LP token price")
		b.logger.Error().Err(err).Msg("Invalid totalSupply")
		return "", err
	}

	// Calculate total value of the pool
	totalValue, err := b.totalValue(ctx)
	if err != nil {
		return "", err
	}

	totalSupplyDecimal := NormalizeAmount(totalSupply, b.config.LPTDecimals)

	pricePerToken := totalValue.Div(totalSupplyDecimal)

	b.logger.Info().
		Str("totalValue", totalValue.String()).
		Str("totalSupply", totalSupplyDecimal.String()).
		Str("pricePerToken", pricePerToken.String()).
		Msg("LP token price calculated successfully")

	return pricePerToken.StringFixed(roundingDecimals), nil
}

// TVL returns the Total Value Locked in the protocol LP pool in USD cents (1 USD = 100 cents).
func (b *BurrBearLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := b.totalValue(ctx)
	if err != nil {
		return "", err
	}

	b.logger.Info().
		Str("totalValue", totalValue.String()).
		Msg("TVL calculated successfully")

	return totalValue.StringFixed(roundingDecimals), nil
}

func (b *BurrBearLPPriceProvider) GetConfig(ctx context.Context, poolAddress string, client *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(poolAddress) {
		err = fmt.Errorf("invalid smart contract address, '%s'", poolAddress)
		return nil, err
	}

	poolContract, err := sc.NewBalancerBasePool(common.HexToAddress(poolAddress), client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instatiate Balancer Base Pool contract on LP Token")
		return nil, err
	}

	bbc := BurrBearConfig{}
	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	// returns as [32]byte
	poolId, err := poolContract.BalancerBasePoolCaller.GetPoolId(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain poolID for bex pool %s, %v", poolAddress, err)
		return nil, err
	}
	bbc.PoolId = poolId

	// decimals is uint8
	decimals, err := poolContract.BalancerBasePoolCaller.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for LP token %s, %v", poolAddress, err)
		return nil, err
	}
	bbc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(bbc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// totalValue calculates the total value of all assets in the pool
func (b *BurrBearLPPriceProvider) totalValue(ctx context.Context) (decimal.Decimal, error) {

	// Get pool tokens and balances
	poolTokens, err := b.vault.BalancerVaultCaller.GetPoolTokens(&bind.CallOpts{Context: ctx, Pending: false}, b.config.PoolId)
	if err != nil {
		return decimal.Zero, fmt.Errorf("failed to get pool tokens: %w", err)
	}

	totalValue := decimal.Zero
	for i, tokenAddress := range poolTokens.Tokens {
		token := strings.ToLower(tokenAddress.Hex())

		//verify this is always sound for all pool types
		if token == strings.ToLower(b.poolAddress.Hex()) {
			// ignore when some of the LP token is locked in pool itself
			// this is why should use actualSupply instead of totalSupply
			continue
		}

		price, err := b.getPrice(token)
		if err != nil {
			return decimal.Zero, err
		}

		// Calculate token value
		balance := poolTokens.Balances[i]

		amount := NormalizeAmount(balance, price.Decimals)

		value := amount.Mul(price.Price)

		totalValue = totalValue.Add(value)

	}

	return totalValue, nil
}

// getPrice gets the price for a token from the price map
func (b *BurrBearLPPriceProvider) getPrice(tokenAddress string) (*Price, error) {
	price, ok := b.priceMap[tokenAddress]
	if !ok {
		return nil, fmt.Errorf("price not found for token %s", tokenAddress)
	}
	return &price, nil
}
