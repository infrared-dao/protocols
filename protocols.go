// Package protocols provides standardized protocol implementations
// for computing LP data, e.g. LP token prices and TVL values
package protocols

import (
	"context"
	"math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/infrared-dao/protocols/fetchers"
	"github.com/shopspring/decimal"
)

// Version information
const Version = "0.4.0"

// Protocol defines methods for querying DeFi protocol metrics.
// Implementations must be safe for concurrent use.
type Protocol interface {
	// GetConfig returns static configuration data needed for LP token price/TVL calculation.
	// Returns an error if the configuration cannot be prepared.
	GetConfig(ctx context.Context, address string, client bind.ContractBackend) ([]byte, error)

	// Initialize performs any needed setup for the data structures to call other functions
	// Returns any error which occurs in setup
	Initialize(ctx context.Context, client bind.ContractBackend, httpClient fetchers.HttpClient) error

	// LPTokenPrice returns the current price of the protocol's LP token
	// in USD.
	// Returns an error if the price cannot be determined.
	LPTokenPrice(ctx context.Context) (string, error)

	// TVL returns the Total Value Locked in the protocol in USD.
	// Returns an error if the TVL cannot be determined.
	TVL(ctx context.Context) (string, error)

	// TVLBreakdown returns the breakdown of TVL by underlying tokens.
	// The returned map keys are token addresses (lowercase) and values are TokenTVL structs
	// containing the token amount and its USD value contribution to the total TVL.
	// This method should be independent of price and retrievable by a specified block.
	// Returns an error if the breakdown cannot be determined.
	TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error)

	// UpdateBlock sets the internal block and priceMap state to different time
	// If the protocol doesn't need a price map the second param can be nil
	UpdateBlock(block *big.Int, prices map[string]Price)
}

// Useful helper functions common across all Protocol adapter implementations

const roundingDecimals = 8

// Type to track a token Price along with the number of decimals used representing token amounts
type Price struct {
	TokenName string
	Decimals  uint
	Price     decimal.Decimal
}

// TokenTVL represents the TVL contribution of a single token
type TokenTVL struct {
	TokenAddress string          // Token contract address (lowercase)
	TokenSymbol  string          // Token symbol for reference
	Amount       decimal.Decimal // Normalized token amount
	USDValue     decimal.Decimal // USD value of this token's contribution
	Ratio        decimal.Decimal // Ratio of this token's value to total TVL (0-1)
}

// NormalizeAmount returns a token amount as a big.Int and the number of decimals (uint) get a decimal.Decimal amount
func NormalizeAmount(amount *big.Int, decimals uint) decimal.Decimal {
	divisor := decimal.New(1, int32(decimals))
	return decimal.NewFromBigInt(amount, 0).Div(divisor)
}
