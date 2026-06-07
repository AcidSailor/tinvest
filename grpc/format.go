package grpc

import (
	"fmt"

	"github.com/acidsailor/tinvest"

	"github.com/acidsailor/tinvest/money"
	pb "github.com/acidsailor/tinvest/pb"
)

// FormatMoney formats a MoneyValue as "250.50 RUB". Shows 2 decimal places
// (nano is truncated to centesimal precision). Returns an error if moneyValue
// is nil or units and nano have mixed signs.
func FormatMoney(moneyValue *pb.MoneyValue) (string, error) {
	if moneyValue == nil {
		return "", fmt.Errorf(
			"%w: moneyValue: %w",
			tinvest.ErrClient,
			tinvest.ErrNil,
		)
	}
	s, err := money.FormatMoney(
		moneyValue.GetUnits(), moneyValue.GetNano(), moneyValue.GetCurrency(),
	)
	if err != nil {
		return "", fmt.Errorf("%w: %w", tinvest.ErrClient, err)
	}
	return s, nil
}

// FormatQuotation formats a Quotation as a decimal string. Shows 2 decimal
// places only when Nano is non-zero (truncated to centesimal precision).
// Returns an error if quotation is nil or units and nano have mixed signs.
func FormatQuotation(quotation *pb.Quotation) (string, error) {
	if quotation == nil {
		return "", fmt.Errorf(
			"%w: quotation: %w",
			tinvest.ErrClient,
			tinvest.ErrNil,
		)
	}
	s, err := money.FormatQuotation(quotation.GetUnits(), quotation.GetNano())
	if err != nil {
		return "", fmt.Errorf("%w: %w", tinvest.ErrClient, err)
	}
	return s, nil
}
