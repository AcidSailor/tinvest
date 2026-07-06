// Package money provides protobuf-free conversions between T-Invest units/nano
// money values (int64 units + int32 nano-billionths) and udecimal.Decimal, plus
// sign handling and display formatting. It is the single home of this math so
// both the proto-typed helpers in package grpc and JSON-typed callers (the
// MCP) share one implementation.
package money

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/quagmt/udecimal"
)

// Sentinel errors returned by this package. Each names a specific failure
// condition and is matched directly with errors.Is — there is no broad
// package-level sentinel that they all wrap.
var (
	// ErrConversion indicates a failure (or invalid input) converting between
	// units/nano and decimal.
	ErrConversion = errors.New("conversion")
	// ErrOverflow indicates a value does not fit the target representation
	// (int64 units, int32 nano, or 9-digit fractional precision).
	ErrOverflow = errors.New("overflow")
)

const (
	nanoPrecision    = 9
	nanoPerHundredth = 10_000_000 // 1e9 / 100
	// nanoLimit is 10^9: nano carries only the fractional billionths, so its
	// magnitude must stay strictly below one whole unit. |nano| >= nanoLimit is
	// not a valid split — the whole part belongs in units.
	nanoLimit = 1_000_000_000
)

// nanoFactor is 10^9 as a udecimal. It is built once at package init from a
// fixed, in-range literal, so udecimal.NewFromInt64 cannot fail here; the panic
// guards that impossible error at init rather than any runtime input.
var nanoFactor = func() udecimal.Decimal {
	d, err := udecimal.NewFromInt64(1_000_000_000, 0)
	if err != nil {
		panic(
			fmt.Errorf("invalid nano factor: %w", err),
		)
	}
	return d
}()

// ValidateSigns rejects invalid (units, nano) pairs: a mixed-sign pair (which
// the API never emits and which arithmetic would silently coerce), or a nano
// whose magnitude reaches one whole unit (|nano| >= 1e9). nano carries only the
// fractional billionths, so an out-of-range nano is not a valid split — the
// whole part belongs in units — and combining it would silently produce a wrong
// value (e.g. units=1, nano=1.5e9 would yield 2.5).
func ValidateSigns(units int64, nano int32) error {
	if (units > 0 && nano < 0) || (units < 0 && nano > 0) {
		return fmt.Errorf(
			"mixed sign (units=%d, nano=%d): %w",
			units, nano, ErrConversion,
		)
	}
	if nano <= -nanoLimit || nano >= nanoLimit {
		return fmt.Errorf(
			"nano %d outside valid range (-1e9, 1e9): %w", nano, ErrOverflow,
		)
	}
	return nil
}

// NormalizeSign validates signs and returns the sign prefix ("-" or "") with
// the absolute units and nano. It errors if units or nano is its integer
// minimum, which cannot be negated in two's complement.
func NormalizeSign(
	units int64, nano int32,
) (string, int64, int32, error) {
	if err := ValidateSigns(units, nano); err != nil {
		return "", 0, 0, err
	}
	if units < 0 || nano < 0 {
		if units == math.MinInt64 || nano == math.MinInt32 {
			return "", 0, 0, fmt.Errorf(
				"value at integer minimum cannot be negated: %w",
				ErrOverflow,
			)
		}
		return "-", -units, -nano, nil
	}
	return "", units, nano, nil
}

// UnitsNanoToDecimal combines units + nano/1e9 into a decimal.
func UnitsNanoToDecimal(
	units int64, nano int32,
) (udecimal.Decimal, error) {
	if err := ValidateSigns(units, nano); err != nil {
		return udecimal.Decimal{}, err
	}
	u, err := udecimal.NewFromInt64(units, 0)
	if err != nil {
		return udecimal.Decimal{}, fmt.Errorf(
			"units: %w: %w",
			ErrConversion,
			err,
		)
	}
	n, err := udecimal.NewFromInt64(int64(nano), nanoPrecision)
	if err != nil {
		return udecimal.Decimal{}, fmt.Errorf(
			"nano: %w: %w",
			ErrConversion,
			err,
		)
	}
	return u.Add(n), nil
}

// DecimalToUnitsNano splits a decimal into units + nano, rejecting values with
// more than 9 fractional digits or that overflow the target integer types.
func DecimalToUnitsNano(d udecimal.Decimal) (int64, int32, error) {
	units := d.Trunc(0)
	frac := d.Sub(units)
	nanoDecimal := frac.Mul(nanoFactor).Trunc(0)

	u, err := units.Int64()
	if err != nil {
		return 0, 0, fmt.Errorf("units: %w: %w", ErrOverflow, err)
	}
	n, err := nanoDecimal.Int64()
	if err != nil {
		return 0, 0, fmt.Errorf("nano: %w: %w", ErrOverflow, err)
	}
	if n > math.MaxInt32 || n < math.MinInt32 {
		return 0, 0, fmt.Errorf(
			"nano value %d exceeds int32 range: %w", n, ErrOverflow,
		)
	}
	nano := int32(n)

	reconstructed, err := UnitsNanoToDecimal(u, nano)
	if err != nil {
		return 0, 0, err
	}
	if !reconstructed.Equal(d) {
		return 0, 0, fmt.Errorf(
			"decimal precision exceeds 9 fractional digits: %w",
			ErrOverflow,
		)
	}
	return u, nano, nil
}

// FormatMoney renders units/nano + currency as "250.50 RUB" (2 dp, truncated
// toward zero). A non-zero value below 0.01 therefore renders as "0.00"; when
// sub-cent exactness matters, format [UnitsNanoToDecimal] instead.
func FormatMoney(
	units int64, nano int32, currency string,
) (string, error) {
	sign, u, n, err := NormalizeSign(units, nano)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"%s%d.%02d %s",
		sign, u, n/nanoPerHundredth, strings.ToUpper(currency),
	), nil
}

// FormatQuotation renders units/nano as a decimal string; 2 dp (truncated
// toward zero) only when nano is non-zero. A non-zero nano below 0.01 truncates
// to ".00"; when sub-cent exactness matters, format [UnitsNanoToDecimal]
// instead.
func FormatQuotation(units int64, nano int32) (string, error) {
	sign, u, n, err := NormalizeSign(units, nano)
	if err != nil {
		return "", err
	}
	if n == 0 {
		return fmt.Sprintf("%s%d", sign, u), nil
	}
	return fmt.Sprintf(
		"%s%d.%02d", sign, u, n/nanoPerHundredth,
	), nil
}
