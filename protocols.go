// Package protocols provides standardized protocol implementations
// for computing LP data, e.g. LP token prices and TVL values
package protocols

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

// Version information
const Version = "0.4.0"

// Protocol defines methods for querying DeFi protocol metrics.
// Implementations must be safe for concurrent use.
type Protocol interface {
	// GetConfig returns static configuration data needed for LP token price/TVL calculation.
	// Returns an error if the configuration cannot be prepared.
	GetConfig(ctx context.Context, address string, ethClient *ethclient.Client) ([]byte, error)

	// Initialize performs any needed setup for the data structures to call other functions
	// Returns any error which occurs in setup
	Initialize(ctx context.Context, client *ethclient.Client) error

	// LPTokenPrice returns the current price of the protocol's LP token in USD.
	// Returns an error if the price cannot be determined.
	LPTokenPrice(ctx context.Context) (string, error)

	// TVL returns the Total Value Locked in the protocol in USD.
	// Returns an error if the TVL cannot be determined.
	TVL(ctx context.Context) (string, error)
}

// Useful helper functions common across all Protocol adapter implementations

// Type to track a token Price along with the number of decimals used representing token amounts
type Price struct {
	TokenName string
	Decimals  uint
	Price     decimal.Decimal
}

// Given a token amount as a big.Int and the number of decimals (uint) get a decimal.Decimal amount
func NormalizeAmount(amount *big.Int, decimals uint) decimal.Decimal {
	divisor := pow10(decimals)
	return decimal.NewFromBigInt(amount, 0).Div(divisor)
}

// Helper to raise 10 to the power passed in and produce a decimal.Decimal for normalization
func pow10(n uint) decimal.Decimal {
	exp := int64(n)
	base := decimal.NewFromInt(10)
	result := base.Pow(decimal.NewFromInt(exp))

	return result
}
