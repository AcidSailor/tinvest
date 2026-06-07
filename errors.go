package tinvest

import (
	"errors"
)

// ErrClient is the sentinel error for the tinvest package.
// Errors returned by this package's own validation and conversion functions wrap it,
// enabling callers to detect them with errors.Is(err, ErrClient).
// Errors from gRPC RPC calls are passed through unwrapped.
var ErrClient = errors.New("tinvest client")

// Sub-sentinel errors for finer-grained errors.Is matching. Call sites join
// them with ErrClient (e.g. fmt.Errorf("%w: ...: %w", ErrClient, ErrNil)), so
// errors.Is(err, ErrClient) still matches alongside the specific sentinel.
var (
	// ErrNil indicates a required argument was nil.
	ErrNil = errors.New("nil")
	// ErrInvalidConfig indicates a configuration value is missing or invalid.
	ErrInvalidConfig = errors.New("invalid config")
)
