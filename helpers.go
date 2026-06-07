package tinvest

import (
	"fmt"

	"github.com/quagmt/udecimal"

	"github.com/acidsailor/tinvest/money"
	pb "github.com/acidsailor/tinvest/pb"
)

// QuotationToDecimal converts a proto Quotation to udecimal.Decimal.
// Returns an error if quotation is nil.
func QuotationToDecimal(quotation *pb.Quotation) (udecimal.Decimal, error) {
	if quotation == nil {
		return udecimal.Decimal{}, fmt.Errorf(
			"%w: quotation: %w",
			ErrClient,
			ErrNil,
		)
	}
	d, err := money.UnitsNanoToDecimal(quotation.Units, quotation.Nano)
	if err != nil {
		return udecimal.Decimal{}, fmt.Errorf("%w: %w", ErrClient, err)
	}
	return d, nil
}

// DecimalToQuotation converts a udecimal.Decimal to a proto Quotation.
func DecimalToQuotation(d udecimal.Decimal) (*pb.Quotation, error) {
	units, nano, err := money.DecimalToUnitsNano(d)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrClient, err)
	}
	return &pb.Quotation{Units: units, Nano: nano}, nil
}

// MoneyValueToDecimal converts a proto MoneyValue to udecimal.Decimal.
// The currency field is dropped. Returns an error if moneyValue is nil.
func MoneyValueToDecimal(moneyValue *pb.MoneyValue) (udecimal.Decimal, error) {
	if moneyValue == nil {
		return udecimal.Decimal{}, fmt.Errorf(
			"%w: moneyValue: %w",
			ErrClient,
			ErrNil,
		)
	}
	d, err := money.UnitsNanoToDecimal(moneyValue.Units, moneyValue.Nano)
	if err != nil {
		return udecimal.Decimal{}, fmt.Errorf("%w: %w", ErrClient, err)
	}
	return d, nil
}

// DecimalToMoneyValue converts a udecimal.Decimal and currency to a proto MoneyValue.
func DecimalToMoneyValue(
	d udecimal.Decimal,
	currency string,
) (*pb.MoneyValue, error) {
	units, nano, err := money.DecimalToUnitsNano(d)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrClient, err)
	}
	return &pb.MoneyValue{Currency: currency, Units: units, Nano: nano}, nil
}

// MoneyValueToQuotation converts a proto MoneyValue to a Quotation, dropping
// the currency. Returns an error if moneyValue is nil or has mixed signs.
func MoneyValueToQuotation(moneyValue *pb.MoneyValue) (*pb.Quotation, error) {
	if moneyValue == nil {
		return nil, fmt.Errorf("%w: moneyValue: %w", ErrClient, ErrNil)
	}
	if err := money.ValidateSigns(moneyValue.Units, moneyValue.Nano); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrClient, err)
	}
	return &pb.Quotation{Units: moneyValue.Units, Nano: moneyValue.Nano}, nil
}
