// Package protocols provides standardized protocol implementations
// for computing LP data, e.g. LP token prices and TVL values
package protocols

import (
	"context"
	"math/big"
)

// Version information
const Version = "0.1.0"

// Protocol defines methods for querying DeFi protocol metrics.
// Implementations must be safe for concurrent use.
type Protocol interface {
	// LPTokenPrice returns the current price of the protocol's LP token
	// in USD cents (1 USD = 100 cents).
	// Returns an error if the price cannot be determined.
	LPTokenPrice(ctx context.Context) (*big.Int, error)

	// TVL returns the Total Value Locked in the protocol in USD cents
	// (1 USD = 100 cents).
	// Returns an error if the TVL cannot be determined.
	TVL(ctx context.Context) (*big.Int, error)

	// LPTokensIssued returns the total number of LP tokens issued.
	// Returns an error if the number of total LP tokens issued cannot be
	// determined.
	LPTokensIssued(ctx context.Context) (*big.Int, error)
}
