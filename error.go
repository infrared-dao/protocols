package protocols

import (
	"errors"
)

// ErrPriceNotReadyYet define a custom error indicating that the price is not ready yet.
var ErrPriceNotReadyYet = errors.New("price is not ready yet")

// ErrTVLBreakdownNotImplemented define a custom error indicating that TVLBreakdown is not implemented for a provider.
var ErrTVLBreakdownNotImplemented = errors.New("TVLBreakdown not implemented")
