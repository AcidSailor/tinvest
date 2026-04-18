package tinvest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pb "github.com/acidsailor/tinvest/pb"
)

func TestNewInstrumentRequest(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		wantType  pb.InstrumentIdType
		wantID    string
		wantClass string
	}{
		{
			"uuid → uid",
			"11111111-2222-3333-4444-555555555555",
			pb.InstrumentIdType_INSTRUMENT_ID_TYPE_UID,
			"11111111-2222-3333-4444-555555555555",
			"",
		},
		{
			"ticker_class → ticker",
			"SBER_TQBR",
			pb.InstrumentIdType_INSTRUMENT_ID_TYPE_TICKER,
			"SBER",
			"TQBR",
		},
		{
			"figi",
			"BBG004730RP0",
			pb.InstrumentIdType_INSTRUMENT_ID_TYPE_FIGI,
			"BBG004730RP0",
			"",
		},
		{
			"trailing underscore → figi",
			"SBER_",
			pb.InstrumentIdType_INSTRUMENT_ID_TYPE_FIGI,
			"SBER_",
			"",
		},
		{
			"leading underscore → figi",
			"_TQBR",
			pb.InstrumentIdType_INSTRUMENT_ID_TYPE_FIGI,
			"_TQBR",
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := NewInstrumentRequest(tt.id)
			assert.Equal(t, tt.wantType, req.GetIdType())
			assert.Equal(t, tt.wantID, req.GetId())
			if tt.wantClass == "" {
				assert.Nil(t, req.ClassCode)
			} else {
				assert.Equal(t, tt.wantClass, req.GetClassCode())
			}
		})
	}
}

func TestPerLotMoney(t *testing.T) {
	tests := []struct {
		name     string
		price    *pb.Quotation
		inst     *pb.Instrument
		units    int64
		nano     int32
		currency string
		wantErr  bool
	}{
		{
			"positive",
			&pb.Quotation{Units: 516, Nano: 550000000},
			&pb.Instrument{Lot: 10, Currency: "rub"},
			5165,
			500000000,
			"rub",
			false,
		},
		{
			"nil instrument",
			&pb.Quotation{Units: 1},
			nil,
			0,
			0,
			"",
			true,
		},
		{
			"zero lot",
			&pb.Quotation{Units: 1},
			&pb.Instrument{Lot: 0, Currency: "rub"},
			0,
			0,
			"",
			true,
		},
		{
			"negative lot",
			&pb.Quotation{Units: 1},
			&pb.Instrument{Lot: -1, Currency: "rub"},
			0,
			0,
			"",
			true,
		},
		{
			"nil price",
			nil,
			&pb.Instrument{Lot: 1, Currency: "rub"},
			0,
			0,
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mv, err := PerLotMoney(tt.price, tt.inst)
			if tt.wantErr {
				require.Error(t, err)
				assert.ErrorIs(t, err, ErrClient)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.units, mv.Units)
			assert.Equal(t, tt.nano, mv.Nano)
			assert.Equal(t, tt.currency, mv.Currency)
		})
	}
}
