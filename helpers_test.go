package tinvest

import (
	"testing"

	"github.com/quagmt/udecimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pb "github.com/acidsailor/tinvest/pb"
)

func mustParse(s string) udecimal.Decimal {
	d, err := udecimal.Parse(s)
	if err != nil {
		panic(err)
	}
	return d
}

func TestQuotationToDecimal(t *testing.T) {
	tests := []struct {
		name     string
		q        *pb.Quotation
		expected string
		wantErr  bool
	}{
		{"positive", &pb.Quotation{Units: 114, Nano: 250000000}, "114.25", false},
		{"zero", &pb.Quotation{Units: 0, Nano: 0}, "0", false},
		{"negative", &pb.Quotation{Units: -200, Nano: -500000000}, "-200.5", false},
		{"negative zero units", &pb.Quotation{Units: 0, Nano: -100000000}, "-0.1", false},
		{"units only", &pb.Quotation{Units: 42, Nano: 0}, "42", false},
		{"nano only", &pb.Quotation{Units: 0, Nano: 100000000}, "0.1", false},
		{"small nano", &pb.Quotation{Units: 1, Nano: 1}, "1.000000001", false},
		{"nil returns error", nil, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := QuotationToDecimal(tt.q)
			if tt.wantErr {
				require.Error(t, err)
				assert.ErrorIs(t, err, ErrTInvestClient)
				assert.ErrorIs(t, err, ErrNil)
				return
			}
			require.NoError(t, err)
			assert.True(t, result.Equal(mustParse(tt.expected)))
		})
	}
}

func TestDecimalToQuotation(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedUnits int64
		expectedNano  int32
		wantErr       bool
	}{
		{"positive", "114.25", 114, 250000000, false},
		{"zero", "0", 0, 0, false},
		{"negative", "-200.5", -200, -500000000, false},
		{"integer", "42", 42, 0, false},
		{"negative zero units", "-0.5", 0, -500000000, false},
		{"small fraction", "1.000000001", 1, 1, false},
		{"sub-nano precision", "1.0000000009", 0, 0, true},
		{"units overflow", "9999999999999999999999999999", 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := mustParse(tt.input)
			q, err := DecimalToQuotation(d)
			if tt.wantErr {
				require.Error(t, err)
				assert.ErrorIs(t, err, ErrTInvestClient)
				assert.ErrorIs(t, err, ErrOverflow)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.expectedUnits, q.Units)
			assert.Equal(t, tt.expectedNano, q.Nano)
		})
	}
}

func TestMoneyValueToDecimal(t *testing.T) {
	tests := []struct {
		name     string
		m        *pb.MoneyValue
		expected string
		wantErr  bool
	}{
		{"positive", &pb.MoneyValue{Currency: "rub", Units: 100, Nano: 500000000}, "100.5", false},
		{"zero", &pb.MoneyValue{Currency: "rub", Units: 0, Nano: 0}, "0", false},
		{"negative", &pb.MoneyValue{Currency: "rub", Units: -50, Nano: -250000000}, "-50.25", false},
		{"negative zero units", &pb.MoneyValue{Currency: "rub", Units: 0, Nano: -100000000}, "-0.1", false},
		{"nano only", &pb.MoneyValue{Currency: "usd", Units: 0, Nano: 1}, "0.000000001", false},
		{"nil returns error", nil, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := MoneyValueToDecimal(tt.m)
			if tt.wantErr {
				require.Error(t, err)
				assert.ErrorIs(t, err, ErrTInvestClient)
				assert.ErrorIs(t, err, ErrNil)
				return
			}
			require.NoError(t, err)
			assert.True(t, result.Equal(mustParse(tt.expected)))
		})
	}
}

func TestDecimalToMoneyValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		currency string
		units    int64
		nano     int32
		wantErr  bool
	}{
		{"positive", "114.25", "rub", 114, 250000000, false},
		{"zero", "0", "rub", 0, 0, false},
		{"negative", "-50.25", "usd", -50, -250000000, false},
		{"empty currency", "1.5", "", 1, 500000000, false},
		{"sub-nano precision", "-0.0000000001", "rub", 0, 0, true},
		{"units overflow", "9999999999999999999999999999", "rub", 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := DecimalToMoneyValue(mustParse(tt.input), tt.currency)
			if tt.wantErr {
				require.Error(t, err)
				assert.ErrorIs(t, err, ErrTInvestClient)
				assert.ErrorIs(t, err, ErrOverflow)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.currency, m.Currency)
			assert.Equal(t, tt.units, m.Units)
			assert.Equal(t, tt.nano, m.Nano)
		})
	}
}

func TestRoundTrip_Quotation(t *testing.T) {
	tests := []struct {
		name string
		q    *pb.Quotation
	}{
		{"positive", &pb.Quotation{Units: 123, Nano: 456789000}},
		{"negative", &pb.Quotation{Units: -123, Nano: -456789000}},
		{"negative zero units", &pb.Quotation{Units: 0, Nano: -100000000}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := QuotationToDecimal(tt.q)
			require.NoError(t, err)

			result, err := DecimalToQuotation(d)
			require.NoError(t, err)

			assert.Equal(t, tt.q.Units, result.Units)
			assert.Equal(t, tt.q.Nano, result.Nano)
		})
	}
}

func TestRoundTrip_MoneyValue(t *testing.T) {
	tests := []struct {
		name string
		m    *pb.MoneyValue
	}{
		{"positive", &pb.MoneyValue{Currency: "rub", Units: 123, Nano: 456789000}},
		{"negative", &pb.MoneyValue{Currency: "rub", Units: -123, Nano: -456789000}},
		{"negative zero units", &pb.MoneyValue{Currency: "rub", Units: 0, Nano: -100000000}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := MoneyValueToDecimal(tt.m)
			require.NoError(t, err)

			result, err := DecimalToMoneyValue(d, tt.m.Currency)
			require.NoError(t, err)

			assert.Equal(t, tt.m.Currency, result.Currency)
			assert.Equal(t, tt.m.Units, result.Units)
			assert.Equal(t, tt.m.Nano, result.Nano)
		})
	}
}
