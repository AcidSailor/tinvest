package tinvest

import (
	"fmt"

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
	if margin == nil {
		return nil, fmt.Errorf("%w: margin: %w", ErrClient, ErrNil)
	}
	step := margin.GetMinPriceIncrement()
	stepValue := margin.GetMinPriceIncrementAmount()
	if step == nil || stepValue == nil {
		return nil, fmt.Errorf(
			"%w: futures margin missing price increment: %w",
			ErrClient,
			ErrNil,
		)
	}
	stepDec, err := QuotationToDecimal(step)
	if err != nil {
		return nil, err
	}
	if stepDec.IsZero() {
		return nil, fmt.Errorf(
			"%w: futures step is zero: %w",
			ErrClient,
			ErrConversion,
		)
	}
	stepValDec, err := QuotationToDecimal(stepValue)
	if err != nil {
		return nil, err
	}
	pv, err := stepValDec.Div(stepDec)
	if err != nil {
		return nil, fmt.Errorf(
			"%w: point value div: %w: %w",
			ErrClient,
			ErrConversion,
			err,
		)
	}
	return DecimalToQuotation(pv)
}
