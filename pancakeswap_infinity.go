package protocols

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/infrared-dao/protocols/fetchers"
	"github.com/infrared-dao/protocols/internal/sc"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

// CLPoolManager contract address on BNB Chain
const CLPoolManagerAddressBSC = "0xa0FfB9c1CE1Fe56963B0321B32E7A0302114058b"

// WBNB address on BNB Chain - used to map native BNB (zero address) for price lookups
const WBNBAddressBSC = "0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c"

// zeroAddress represents native token in PancakeSwap Infinity pools
const zeroAddress = "0x0000000000000000000000000000000000000000"

// Q96 is 2^96, used for sqrtPriceX96 calculations
var Q96 = new(big.Int).Exp(big.NewInt(2), big.NewInt(96), nil)

// enforce interface adherence
var _ Protocol = &PancakeSwapInfinityLPPriceProvider{}

// PancakeSwapInfinityConfig defines the configuration for PancakeSwap Infinity CLMM pools.
type PancakeSwapInfinityConfig struct {
	PoolId            string `json:"pool_id"`
	Token0            string `json:"token0"`
	Token1            string `json:"token1"`
	Fee               uint32 `json:"fee"`
	CLPoolManagerAddr string `json:"cl_pool_manager"`
}

// PancakeSwapInfinityLPPriceProvider defines the provider for PancakeSwap Infinity CLMM pool TVL.
type PancakeSwapInfinityLPPriceProvider struct {
	address     common.Address // kept for interface compatibility, not used
	poolId      [32]byte
	block       *big.Int
	priceMap    map[string]Price
	logger      zerolog.Logger
	configBytes []byte
	config      *PancakeSwapInfinityConfig
	contract    *sc.CLPoolManager
}

// NewPancakeSwapInfinityLPPriceProvider creates a new instance of the PancakeSwapInfinityLPPriceProvider.
// Note: For Infinity pools, the poolId is parsed from the config, not the address parameter.
// The address parameter is kept for interface compatibility with other adapters.
func NewPancakeSwapInfinityLPPriceProvider(
	address common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *PancakeSwapInfinityLPPriceProvider {
	p := &PancakeSwapInfinityLPPriceProvider{
		address:     address,
		block:       block,
		priceMap:    prices,
		logger:      logger,
		configBytes: config,
	}
	return p
}

// Initialize checks the configuration/data provided and instantiates the CLPoolManager contract.
func (p *PancakeSwapInfinityLPPriceProvider) Initialize(ctx context.Context, client bind.ContractBackend, httpClient fetchers.HttpClient) error {
	var err error

	p.config = &PancakeSwapInfinityConfig{}
	err = json.Unmarshal(p.configBytes, p.config)
	if err != nil {
		p.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	// Parse poolId from config
	p.poolId, err = hexToBytes32(p.config.PoolId)
	if err != nil {
		p.logger.Error().Err(err).Str("poolId", p.config.PoolId).Msg("failed to parse pool ID from config")
		return fmt.Errorf("invalid pool ID in config: %w", err)
	}

	_, ok := p.priceMap[p.config.Token0]
	if !ok {
		err = fmt.Errorf("no price data found for token0 (%s)", p.config.Token0)
		p.logger.Error().Msg(err.Error())
		return err
	}
	_, ok = p.priceMap[p.config.Token1]
	if !ok {
		err = fmt.Errorf("no price data found for token1 (%s)", p.config.Token1)
		p.logger.Error().Msg(err.Error())
		return err
	}

	// Use configured CLPoolManager address or default to BSC
	clPoolManagerAddr := p.config.CLPoolManagerAddr
	if clPoolManagerAddr == "" {
		clPoolManagerAddr = CLPoolManagerAddressBSC
	}

	p.contract, err = sc.NewCLPoolManager(common.HexToAddress(clPoolManagerAddr), client)
	if err != nil {
		p.logger.Error().Err(err).Msg("failed to instantiate CLPoolManager contract")
		return err
	}

	return nil
}

// LPTokenPrice returns "0" as Infinity CLMM pools don't have traditional LP tokens.
// Use TVL() and TVLBreakdown() instead for pool value metrics.
func (p *PancakeSwapInfinityLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	// CLMM pools don't have a single LP token price like V2 pools
	// Each position is an NFT with its own value based on tick range
	return "0", nil
}

// TVL returns the Total Value Locked in the pool in USD.
func (p *PancakeSwapInfinityLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := p.totalValue(ctx)
	if err != nil {
		return "", err
	}

	p.logger.Debug().
		Str("totalValue", totalValue.String()).
		Msg("TVL calculated successfully")

	return totalValue.StringFixed(roundingDecimals), nil
}

// GetConfig fetches and returns the configuration for the PancakeSwap Infinity pool.
func (p *PancakeSwapInfinityLPPriceProvider) GetConfig(ctx context.Context, address string, client bind.ContractBackend) ([]byte, error) {
	// For Infinity pools, address is the PoolId (bytes32 hex string)
	poolIdBytes, err := hexToBytes32(address)
	if err != nil {
		return nil, fmt.Errorf("invalid pool ID format: %w", err)
	}

	// Use default CLPoolManager address for BSC
	contract, err := sc.NewCLPoolManager(common.HexToAddress(CLPoolManagerAddressBSC), client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate CLPoolManager contract: %w", err)
	}

	opts := &bind.CallOpts{
		Context: ctx,
	}

	// Get pool key to extract token addresses
	poolKey, err := contract.PoolIdToPoolKey(opts, poolIdBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to get pool key for pool %s: %w", address, err)
	}

	config := PancakeSwapInfinityConfig{
		PoolId:            address,
		Token0:            mapNativeToWrapped(strings.ToLower(poolKey.Currency0.Hex())),
		Token1:            mapNativeToWrapped(strings.ToLower(poolKey.Currency1.Hex())),
		Fee:               uint32(poolKey.Fee.Uint64()),
		CLPoolManagerAddr: CLPoolManagerAddressBSC,
	}

	body, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// UpdateBlock sets the internal block and priceMap state to different time.
func (p *PancakeSwapInfinityLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	p.block = block
	if prices != nil {
		p.priceMap = prices
	}
}

// TVLBreakdown returns the breakdown of TVL by underlying tokens.
func (p *PancakeSwapInfinityLPPriceProvider) TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error) {
	// Get pool state
	amount0, amount1, err := p.getVirtualReserves(ctx)
	if err != nil {
		return nil, err
	}

	price0, err := p.getPrice(p.config.Token0)
	if err != nil {
		return nil, err
	}
	amount0Decimal := NormalizeAmount(amount0, price0.Decimals)
	token0USDValue := amount0Decimal.Mul(price0.Price)

	price1, err := p.getPrice(p.config.Token1)
	if err != nil {
		return nil, err
	}
	amount1Decimal := NormalizeAmount(amount1, price1.Decimals)
	token1USDValue := amount1Decimal.Mul(price1.Price)

	totalValue := token0USDValue.Add(token1USDValue)

	// Calculate ratios (handle zero TVL case)
	var token0Ratio, token1Ratio decimal.Decimal
	if totalValue.IsZero() {
		token0Ratio = decimal.Zero
		token1Ratio = decimal.Zero
	} else {
		token0Ratio = token0USDValue.Div(totalValue)
		token1Ratio = token1USDValue.Div(totalValue)
	}

	breakdown := map[string]TokenTVL{
		p.config.Token0: {
			TokenAddress: p.config.Token0,
			TokenSymbol:  price0.TokenName,
			Amount:       amount0Decimal,
			USDValue:     token0USDValue,
			Ratio:        token0Ratio,
		},
		p.config.Token1: {
			TokenAddress: p.config.Token1,
			TokenSymbol:  price1.TokenName,
			Amount:       amount1Decimal,
			USDValue:     token1USDValue,
			Ratio:        token1Ratio,
		},
	}

	p.logger.Debug().
		Str("token0", p.config.Token0).
		Str("token0Amount", amount0Decimal.String()).
		Str("token0USDValue", token0USDValue.String()).
		Str("token0Ratio", token0Ratio.String()).
		Str("token1", p.config.Token1).
		Str("token1Amount", amount1Decimal.String()).
		Str("token1USDValue", token1USDValue.String()).
		Str("token1Ratio", token1Ratio.String()).
		Msg("TVL breakdown calculated successfully")

	return breakdown, nil
}

// Internal Helper methods

// totalValue calculates the total value of the pool.
func (p *PancakeSwapInfinityLPPriceProvider) totalValue(ctx context.Context) (decimal.Decimal, error) {
	amount0, amount1, err := p.getVirtualReserves(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	price0, err := p.getPrice(p.config.Token0)
	if err != nil {
		return decimal.Zero, err
	}
	amount0Decimal := NormalizeAmount(amount0, price0.Decimals)

	price1, err := p.getPrice(p.config.Token1)
	if err != nil {
		return decimal.Zero, err
	}
	amount1Decimal := NormalizeAmount(amount1, price1.Decimals)

	totalValue := amount0Decimal.Mul(price0.Price).Add(amount1Decimal.Mul(price1.Price))
	return totalValue, nil
}

func (p *PancakeSwapInfinityLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := p.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		p.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getVirtualReserves calculates the virtual token reserves from liquidity and sqrtPriceX96.
// For CLMM pools, this uses the formula:
//   - amount0 = L * 2^96 / sqrtPriceX96
//   - amount1 = L * sqrtPriceX96 / 2^96
//
// Note: This is an approximation based on active liquidity at the current tick.
// Works for all PancakeSwap Infinity CLMM pools (not LBAMM/BinPool).
func (p *PancakeSwapInfinityLPPriceProvider) getVirtualReserves(ctx context.Context) (*big.Int, *big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: p.block,
	}

	// Get liquidity
	liquidity, err := p.contract.GetLiquidity(opts, p.poolId)
	if err != nil {
		p.logger.Error().Err(err).Msg("failed to get liquidity")
		return nil, nil, fmt.Errorf("failed to get liquidity: %w", err)
	}

	// Check for zero liquidity
	if liquidity == nil || liquidity.Sign() == 0 {
		p.logger.Warn().Msg("pool has zero liquidity")
		return big.NewInt(0), big.NewInt(0), nil
	}

	// Get slot0 for sqrtPriceX96
	slot0, err := p.contract.GetSlot0(opts, p.poolId)
	if err != nil {
		p.logger.Error().Err(err).Msg("failed to get slot0")
		return nil, nil, fmt.Errorf("failed to get slot0: %w", err)
	}

	sqrtPriceX96 := slot0.SqrtPriceX96

	// Check for zero or invalid sqrtPriceX96
	if sqrtPriceX96 == nil || sqrtPriceX96.Sign() == 0 {
		p.logger.Warn().Msg("pool has zero sqrtPriceX96 - pool may not be initialized")
		return big.NewInt(0), big.NewInt(0), nil
	}

	// Calculate virtual reserves using CLMM math
	// amount0 = L * 2^96 / sqrtPriceX96
	// amount1 = L * sqrtPriceX96 / 2^96

	// amount0 = L * 2^96 / sqrtPriceX96
	amount0 := new(big.Int).Mul(liquidity, Q96)
	amount0.Div(amount0, sqrtPriceX96)

	// amount1 = L * sqrtPriceX96 / 2^96
	amount1 := new(big.Int).Mul(liquidity, sqrtPriceX96)
	amount1.Div(amount1, Q96)

	p.logger.Debug().
		Str("liquidity", liquidity.String()).
		Str("sqrtPriceX96", sqrtPriceX96.String()).
		Str("amount0", amount0.String()).
		Str("amount1", amount1.String()).
		Str("tick", slot0.Tick.String()).
		Msg("calculated virtual reserves")

	return amount0, amount1, nil
}

// mapNativeToWrapped maps the zero address (native BNB) to WBNB for price lookups.
func mapNativeToWrapped(addr string) string {
	if addr == zeroAddress {
		return WBNBAddressBSC
	}
	return addr
}

// hexToBytes32 converts a hex string (with or without 0x prefix) to [32]byte.
func hexToBytes32(hexStr string) ([32]byte, error) {
	var result [32]byte

	// Remove 0x prefix if present
	hexStr = strings.TrimPrefix(hexStr, "0x")
	hexStr = strings.TrimPrefix(hexStr, "0X")

	if len(hexStr) != 64 {
		return result, fmt.Errorf("hex string must be 64 characters (32 bytes), got %d", len(hexStr))
	}

	bytes := common.FromHex("0x" + hexStr)
	if len(bytes) != 32 {
		return result, fmt.Errorf("failed to decode hex string")
	}

	copy(result[:], bytes)
	return result, nil
}
