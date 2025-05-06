package protocols

import (
	"errors"
)

// ErrPriceNotReadyYet define a custom error indicating that the price is not ready yet.
var ErrPriceNotReadyYet = errors.New("price is not ready yet")
