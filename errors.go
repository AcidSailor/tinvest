package tinvest

import (
	"errors"
)

// Sentinel errors returned by this library's validation and conversion
// functions. Each names a specific failure condition and is matched directly
// with errors.Is — there is no broad package-level sentinel that they all wrap.
// Errors from gRPC RPC calls are passed through unwrapped.
var (
	// ErrInvalidConfig indicates a configuration value is missing or invalid.
	ErrInvalidConfig = errors.New("tinvest: invalid config")
)
