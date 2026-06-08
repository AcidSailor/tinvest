package grpc

import (
	"errors"

	pb "github.com/acidsailor/tinvest/grpc/pb"
	"github.com/acidsailor/tinvest/money"
)

// FormatMoney formats a MoneyValue as "250.50 RUB". Shows 2 decimal places
// (nano is truncated to centesimal precision). Returns an error if moneyValue
// is nil or units and nano have mixed signs.
func FormatMoney(moneyValue *pb.MoneyValue) (string, error) {
	if moneyValue == nil {
		return "", errors.New("tinvest: moneyValue is nil")
	}
	return money.FormatMoney(
		moneyValue.GetUnits(), moneyValue.GetNano(), moneyValue.GetCurrency(),
	)
}

// FormatQuotation formats a Quotation as a decimal string. Shows 2 decimal
// places only when Nano is non-zero (truncated to centesimal precision).
// Returns an error if quotation is nil or units and nano have mixed signs.
func FormatQuotation(quotation *pb.Quotation) (string, error) {
	if quotation == nil {
		return "", errors.New("tinvest: quotation is nil")
	}
	return money.FormatQuotation(quotation.GetUnits(), quotation.GetNano())
}
