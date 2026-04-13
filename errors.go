package tinvest

import "errors"

// ErrTInvestClient is the sentinel error for the tinvest package.
// Errors returned by this package's own validation and conversion functions wrap it,
// enabling callers to detect them with errors.Is(err, ErrTInvestClient).
// Errors from gRPC RPC calls are passed through unwrapped.
var ErrTInvestClient = errors.New("tinvest client")

// Sub-sentinel errors for finer-grained errors.Is matching.
// All wrap ErrTInvestClient, so errors.Is(err, ErrTInvestClient) still works.
var (
	// ErrNil indicates a required argument was nil.
	ErrNil = errors.New("nil")
	// ErrInvalidConfig indicates a configuration value is missing or invalid.
	ErrInvalidConfig = errors.New("invalid config")
	// ErrOverflow indicates a value does not fit the target representation
	// (int64 units, int32 nano, or 9-digit fractional precision).
	ErrOverflow = errors.New("overflow")
	// ErrConversion indicates a failure in decimal ↔ units/nano conversion.
	ErrConversion = errors.New("conversion")
)
