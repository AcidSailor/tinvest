package grpc

import (
	"errors"

	"github.com/quagmt/udecimal"

	pb "github.com/acidsailor/tinvest/grpc/pb"
	"github.com/acidsailor/tinvest/money"
)

// QuotationToDecimal converts a proto Quotation to udecimal.Decimal.
// Returns an error if quotation is nil.
func QuotationToDecimal(quotation *pb.Quotation) (udecimal.Decimal, error) {
	if quotation == nil {
		return udecimal.Decimal{}, errors.New("tinvest: quotation is nil")
	}
	return money.UnitsNanoToDecimal(quotation.Units, quotation.Nano)
}

// DecimalToQuotation converts a udecimal.Decimal to a proto Quotation.
func DecimalToQuotation(d udecimal.Decimal) (*pb.Quotation, error) {
	units, nano, err := money.DecimalToUnitsNano(d)
	if err != nil {
		return nil, err
	}
	return &pb.Quotation{Units: units, Nano: nano}, nil
}

// MoneyValueToDecimal converts a proto MoneyValue to udecimal.Decimal.
// The currency field is dropped. Returns an error if moneyValue is nil.
func MoneyValueToDecimal(moneyValue *pb.MoneyValue) (udecimal.Decimal, error) {
	if moneyValue == nil {
		return udecimal.Decimal{}, errors.New("tinvest: moneyValue is nil")
	}
	return money.UnitsNanoToDecimal(moneyValue.Units, moneyValue.Nano)
}

// DecimalToMoneyValue converts a udecimal.Decimal and currency to a proto MoneyValue.
func DecimalToMoneyValue(
	d udecimal.Decimal,
	currency string,
) (*pb.MoneyValue, error) {
	units, nano, err := money.DecimalToUnitsNano(d)
	if err != nil {
		return nil, err
	}
	return &pb.MoneyValue{Currency: currency, Units: units, Nano: nano}, nil
}

// MoneyValueToQuotation converts a proto MoneyValue to a Quotation, dropping
// the currency. Returns an error if moneyValue is nil or has mixed signs.
func MoneyValueToQuotation(moneyValue *pb.MoneyValue) (*pb.Quotation, error) {
	if moneyValue == nil {
		return nil, errors.New("tinvest: moneyValue is nil")
	}
	if err := money.ValidateSigns(
		moneyValue.Units,
		moneyValue.Nano,
	); err != nil {
		return nil, err
	}
	return &pb.Quotation{Units: moneyValue.Units, Nano: moneyValue.Nano}, nil
}
