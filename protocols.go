// Package protocols provides standardized protocol implementations
// for computing LP data, e.g. LP token prices and TVL values
package protocols

import (
	"context"
)

// Version information
const Version = "0.2.0"

// Protocol defines methods for querying DeFi protocol metrics.
// Implementations must be safe for concurrent use.
type Protocol interface {
	// LPTokenPrice returns the current price of the protocol's LP token
	// in USD.
	// Returns an error if the price cannot be determined.
	LPTokenPrice(ctx context.Context) (string, error)

	// TVL returns the Total Value Locked in the protocol in USD.
	// Returns an error if the TVL cannot be determined.
	TVL(ctx context.Context) (string, error)
}
