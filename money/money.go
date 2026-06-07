// Package money provides protobuf-free conversions between T-Invest units/nano
// money values (int64 units + int32 nano-billionths) and udecimal.Decimal, plus
// sign handling and display formatting. It is the single home of this math so
// both the proto-typed helpers in package tinvest and JSON-typed callers (the
// MCP) share one implementation.
package money

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/quagmt/udecimal"
)

// Err is the sentinel every error from this package wraps.
var Err = errors.New("tinvest money")

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
)

// nanoFactor is 10^9. Panic is intentional: a compile-time constant.
var nanoFactor = func() udecimal.Decimal {
	d, err := udecimal.NewFromInt64(1_000_000_000, 0)
	if err != nil {
		panic("tinvest/money: invalid nano factor: " + err.Error())
	}
	return d
}()

// ValidateSigns rejects mixed-sign (units, nano) pairs, which the API never
// emits and which arithmetic would silently coerce.
func ValidateSigns(units int64, nano int32) error {
	if (units > 0 && nano < 0) || (units < 0 && nano > 0) {
		return fmt.Errorf(
			"%w: mixed sign (units=%d, nano=%d): %w",
			Err, units, nano, ErrConversion,
		)
	}
	return nil
}

// NormalizeSign validates signs and returns the sign prefix ("-" or "") with
// the absolute units and nano.
func NormalizeSign(
	units int64, nano int32,
) (string, int64, int32, error) {
	if err := ValidateSigns(units, nano); err != nil {
		return "", 0, 0, err
	}
	if units < 0 || nano < 0 {
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
			"%w: units: %w: %w", Err, ErrConversion, err,
		)
	}
	n, err := udecimal.NewFromInt64(int64(nano), nanoPrecision)
	if err != nil {
		return udecimal.Decimal{}, fmt.Errorf(
			"%w: nano: %w: %w", Err, ErrConversion, err,
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
		return 0, 0, fmt.Errorf(
			"%w: units: %w: %w", Err, ErrOverflow, err,
		)
	}
	n, err := nanoDecimal.Int64()
	if err != nil {
		return 0, 0, fmt.Errorf(
			"%w: nano: %w: %w", Err, ErrOverflow, err,
		)
	}
	if n > math.MaxInt32 || n < math.MinInt32 {
		return 0, 0, fmt.Errorf(
			"%w: nano value %d exceeds int32 range: %w",
			Err, n, ErrOverflow,
		)
	}
	nano := int32(n)

	reconstructed, err := UnitsNanoToDecimal(u, nano)
	if err != nil {
		return 0, 0, err
	}
	if !reconstructed.Equal(d) {
		return 0, 0, fmt.Errorf(
			"%w: decimal precision exceeds 9 fractional digits: %w",
			Err, ErrOverflow,
		)
	}
	return u, nano, nil
}

// FormatMoney renders units/nano + currency as "250.50 RUB" (2 dp, truncated).
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

// FormatQuotation renders units/nano as a decimal string; 2 dp only when nano
// is non-zero.
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
