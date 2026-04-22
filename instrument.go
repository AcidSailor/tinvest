package tinvest

import (
	"fmt"
	"strings"

	"github.com/google/uuid"

	pb "github.com/acidsailor/tinvest/pb"
)

// NewInstrumentRequest builds an InstrumentRequest with IdType inferred from
// the shape of id:
//   - valid UUID → INSTRUMENT_ID_TYPE_UID
//   - "<TICKER>_<CLASS_CODE>" (at least one char on each side of the last
//     underscore) → INSTRUMENT_ID_TYPE_TICKER with class_code split off
//   - otherwise → INSTRUMENT_ID_TYPE_FIGI
func NewInstrumentRequest(id string) *pb.InstrumentRequest {
	if _, err := uuid.Parse(id); err == nil {
		return &pb.InstrumentRequest{
			IdType: pb.InstrumentIdType_INSTRUMENT_ID_TYPE_UID,
			Id:     id,
		}
	}
	if i := strings.LastIndex(id, "_"); i > 0 && i < len(id)-1 {
		classCode := id[i+1:]
		return &pb.InstrumentRequest{
			IdType:    pb.InstrumentIdType_INSTRUMENT_ID_TYPE_TICKER,
			Id:        id[:i],
			ClassCode: &classCode,
		}
	}
	return &pb.InstrumentRequest{
		IdType: pb.InstrumentIdType_INSTRUMENT_ID_TYPE_FIGI,
		Id:     id,
	}
}

// PerLotMoney returns price * instrument.Lot as a MoneyValue in the
// instrument's currency. Returns an error if instrument or price is nil, if
// lot is non-positive, or if the arithmetic overflows MoneyValue precision.
func PerLotMoney(
	price *pb.Quotation,
	instrument *pb.Instrument,
) (*pb.MoneyValue, error) {
	if instrument == nil {
		return nil, fmt.Errorf("%w: instrument: %w",
			ErrClient, ErrNil)
	}
	lot := instrument.GetLot()
	if lot <= 0 {
		return nil, fmt.Errorf(
			"%w: instrument %s has non-positive lot %d: %w",
			ErrClient,
			instrument.GetUid(),
			lot,
			ErrConversion,
		)
	}
	d, err := QuotationToDecimal(price)
	if err != nil {
		return nil, err
	}
	return DecimalToMoneyValue(d.Mul64(uint64(lot)), instrument.GetCurrency())
}
