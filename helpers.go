package tinvest

import (
	"fmt"
	"math"

	"github.com/quagmt/udecimal"

	pb "github.com/acidsailor/tinvest/pb"
)

const nanoPrecision = 9

// nanoFactor is 10^9 for units/nano ↔ decimal conversion. Panic is intentional —
// the value is a compile-time constant that must never fail to parse.
var nanoFactor = func() udecimal.Decimal {
	d, err := udecimal.NewFromInt64(1_000_000_000, 0)
	if err != nil {
		panic("tinvest: invalid nano factor: " + err.Error())
	}
	return d
}()

// QuotationToDecimal converts a proto Quotation to udecimal.Decimal.
// Returns an error if q is nil.
func QuotationToDecimal(q *pb.Quotation) (udecimal.Decimal, error) {
	if q == nil {
		return udecimal.Decimal{}, fmt.Errorf("%w: %w: quotation", ErrTInvestClient, ErrNil)
	}
	return unitsNanoToDecimal(q.Units, q.Nano)
}

// DecimalToQuotation converts a udecimal.Decimal to a proto Quotation.
func DecimalToQuotation(d udecimal.Decimal) (*pb.Quotation, error) {
	units, nano, err := decimalToUnitsNano(d)
	if err != nil {
		return nil, err
	}
	return &pb.Quotation{Units: units, Nano: nano}, nil
}

// MoneyValueToDecimal converts a proto MoneyValue to udecimal.Decimal.
// The currency field is dropped. Returns an error if m is nil.
func MoneyValueToDecimal(m *pb.MoneyValue) (udecimal.Decimal, error) {
	if m == nil {
		return udecimal.Decimal{}, fmt.Errorf("%w: %w: money value", ErrTInvestClient, ErrNil)
	}
	return unitsNanoToDecimal(m.Units, m.Nano)
}

// DecimalToMoneyValue converts a udecimal.Decimal and currency to a proto MoneyValue.
func DecimalToMoneyValue(d udecimal.Decimal, currency string) (*pb.MoneyValue, error) {
	units, nano, err := decimalToUnitsNano(d)
	if err != nil {
		return nil, err
	}
	return &pb.MoneyValue{Currency: currency, Units: units, Nano: nano}, nil
}

func unitsNanoToDecimal(units int64, nano int32) (udecimal.Decimal, error) {
	u, err := udecimal.NewFromInt64(units, 0)
	if err != nil {
		return udecimal.Decimal{}, fmt.Errorf("%w: %w: units: %w", ErrTInvestClient, ErrConversion, err)
	}
	n, err := udecimal.NewFromInt64(int64(nano), nanoPrecision)
	if err != nil {
		return udecimal.Decimal{}, fmt.Errorf("%w: %w: nano: %w", ErrTInvestClient, ErrConversion, err)
	}
	return u.Add(n), nil
}

func decimalToUnitsNano(d udecimal.Decimal) (int64, int32, error) {
	units := d.Trunc(0)

	frac := d.Sub(units)
	nanoDecimal := frac.Mul(nanoFactor).Trunc(0)

	u, err := units.Int64()
	if err != nil {
		return 0, 0, fmt.Errorf("%w: %w: units: %w", ErrTInvestClient, ErrOverflow, err)
	}
	n, err := nanoDecimal.Int64()
	if err != nil {
		return 0, 0, fmt.Errorf("%w: %w: nano: %w", ErrTInvestClient, ErrOverflow, err)
	}

	if n > math.MaxInt32 || n < math.MinInt32 {
		return 0, 0, fmt.Errorf("%w: %w: nano value %d exceeds int32 range", ErrTInvestClient, ErrOverflow, n)
	}
	nano := int32(n)

	// Round-trip check: reject decimals whose precision exceeds 9 fractional digits,
	// since the units/nano representation would silently truncate them.
	reconstructed, err := unitsNanoToDecimal(u, nano)
	if err != nil {
		return 0, 0, err
	}
	// udecimal.Equal compares by value (1.50 == 1.5), which is intentional here.
	if !reconstructed.Equal(d) {
		return 0, 0, fmt.Errorf("%w: %w: decimal precision exceeds 9 fractional digits", ErrTInvestClient, ErrOverflow)
	}

	return u, nano, nil
}
