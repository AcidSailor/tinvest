package grpc

import (
	"errors"
	"fmt"

	pb "github.com/acidsailor/tinvest/grpc/pb"
	"github.com/acidsailor/tinvest/money"
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
		return nil, errors.New("tinvest: margin is nil")
	}
	step := margin.GetMinPriceIncrement()
	stepValue := margin.GetMinPriceIncrementAmount()
	if step == nil || stepValue == nil {
		return nil, errors.New(
			"tinvest: futures margin missing price increment",
		)
	}
	stepDec, err := QuotationToDecimal(step)
	if err != nil {
		return nil, err
	}
	if stepDec.IsZero() {
		return nil, fmt.Errorf("futures step is zero: %w", money.ErrConversion)
	}
	stepValDec, err := QuotationToDecimal(stepValue)
	if err != nil {
		return nil, err
	}
	pv, err := stepValDec.Div(stepDec)
	if err != nil {
		return nil, fmt.Errorf(
			"point value div: %w: %w",
			money.ErrConversion,
			err,
		)
	}
	return DecimalToQuotation(pv)
}
