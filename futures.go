package tinvest

import (
	"fmt"

	"github.com/quagmt/udecimal"

	pb "github.com/acidsailor/tinvest/pb"
)

// FuturesPointValue returns the value of a single price point for a futures
// contract, derived as min_price_increment_amount / min_price_increment from
// the given GetFuturesMargin response. The result is a unitless multiplier
// (price-point value), not a currency amount. Returns an error if margin is
// nil, either increment field is missing, or the step is zero.
func FuturesPointValue(
	margin *pb.GetFuturesMarginResponse,
) (*pb.Quotation, error) {
	f := func() (udecimal.Decimal, error) {
		if margin == nil {
			return udecimal.Decimal{}, fmt.Errorf("margin: %w", ErrNil)
		}
		step := margin.GetMinPriceIncrement()
		stepValue := margin.GetMinPriceIncrementAmount()
		if step == nil || stepValue == nil {
			return udecimal.Decimal{}, fmt.Errorf(
				"futures margin missing price increment: %w",
				ErrNil,
			)
		}
		stepDec, err := QuotationToDecimal(step)
		if err != nil {
			return udecimal.Decimal{}, err
		}
		if stepDec.IsZero() {
			return udecimal.Decimal{}, fmt.Errorf(
				"futures step is zero: %w",
				ErrConversion,
			)
		}
		stepValDec, err := QuotationToDecimal(stepValue)
		if err != nil {
			return udecimal.Decimal{}, err
		}
		pv, err := stepValDec.Div(stepDec)
		if err != nil {
			return udecimal.Decimal{}, fmt.Errorf(
				"point value div: %w: %w",
				ErrConversion,
				err,
			)
		}
		return pv, nil
	}
	pv, err := f()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrClient, err)
	}
	return DecimalToQuotation(pv)
}
