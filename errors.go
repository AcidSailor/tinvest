package tinvest

import (
	"errors"

	"github.com/acidsailor/tinvest/money"
)

// ErrClient is the sentinel error for the tinvest package.
// Errors returned by this package's own validation and conversion functions wrap it,
// enabling callers to detect them with errors.Is(err, ErrClient).
// Errors from gRPC RPC calls are passed through unwrapped.
var ErrClient = errors.New("tinvest client")

// Sub-sentinel errors for finer-grained errors.Is matching.
// All wrap ErrClient, so errors.Is(err, ErrClient) still works.
var (
	// ErrNil indicates a required argument was nil.
	ErrNil = errors.New("nil")
	// ErrInvalidConfig indicates a configuration value is missing or invalid.
	ErrInvalidConfig = errors.New("invalid config")
)

// ErrOverflow and ErrConversion now live in package money; re-exported so
// existing errors.Is(err, tinvest.ErrConversion) call sites keep matching.
var (
	ErrOverflow   = money.ErrOverflow
	ErrConversion = money.ErrConversion
)
