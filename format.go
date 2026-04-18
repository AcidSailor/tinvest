package tinvest

import (
	"fmt"
	"strings"

	pb "github.com/acidsailor/tinvest/pb"
)

// nanoPerHundredth is the nano-unit value representing one hundredth of a
// currency unit (1e9 / 100).
const nanoPerHundredth = 10_000_000

// FormatMoney formats a MoneyValue as "250.50 RUB". Shows 2 decimal places
// (nano is truncated to centesimal precision). Returns an error if moneyValue
// is nil or units and nano have mixed signs.
func FormatMoney(moneyValue *pb.MoneyValue) (string, error) {
	if moneyValue == nil {
		return "", fmt.Errorf("%w: moneyValue: %w", ErrClient, ErrNil)
	}
	sign, units, nano, err := normalizeSign(
		moneyValue.GetUnits(),
		moneyValue.GetNano(),
	)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"%s%d.%02d %s",
		sign,
		units,
		nano/nanoPerHundredth,
		strings.ToUpper(moneyValue.GetCurrency()),
	), nil
}

// FormatQuotation formats a Quotation as a decimal string. Shows 2 decimal
// places only when Nano is non-zero (truncated to centesimal precision).
// Returns an error if quotation is nil or units and nano have mixed signs.
func FormatQuotation(quotation *pb.Quotation) (string, error) {
	if quotation == nil {
		return "", fmt.Errorf("%w: quotation: %w", ErrClient, ErrNil)
	}
	sign, units, nano, err := normalizeSign(
		quotation.GetUnits(),
		quotation.GetNano(),
	)
	if err != nil {
		return "", err
	}
	if nano == 0 {
		return fmt.Sprintf("%s%d", sign, units), nil
	}
	return fmt.Sprintf("%s%d.%02d", sign, units, nano/nanoPerHundredth), nil
}

// validateSigns rejects (units, nano) pairs with mixed signs, which the
// T-Invest API never emits and which would otherwise be silently coerced
// into an unintended value by arithmetic combination.
func validateSigns(units int64, nano int32) error {
	if (units > 0 && nano < 0) || (units < 0 && nano > 0) {
		return fmt.Errorf(
			"%w: mixed sign (units=%d, nano=%d): %w",
			ErrClient,
			units,
			nano,
			ErrConversion,
		)
	}
	return nil
}

// normalizeSign validates signs and returns the sign prefix ("-" or "")
// along with the absolute values of units and nano.
func normalizeSign(units int64, nano int32) (string, int64, int32, error) {
	if err := validateSigns(units, nano); err != nil {
		return "", 0, 0, err
	}
	if units < 0 || nano < 0 {
		return "-", -units, -nano, nil
	}
	return "", units, nano, nil
}
