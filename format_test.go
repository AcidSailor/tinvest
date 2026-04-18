package tinvest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pb "github.com/acidsailor/tinvest/pb"
)

func TestFormatMoney(t *testing.T) {
	tests := []struct {
		name    string
		input   *pb.MoneyValue
		want    string
		wantErr bool
	}{
		{
			"positive",
			&pb.MoneyValue{Currency: "rub", Units: 250, Nano: 500000000},
			"250.50 RUB",
			false,
		},
		{
			"zero nano",
			&pb.MoneyValue{Currency: "usd", Units: 100, Nano: 0},
			"100.00 USD",
			false,
		},
		{"nil", nil, "", true},
		{
			"negative",
			&pb.MoneyValue{Currency: "rub", Units: -1, Nano: -500000000},
			"-1.50 RUB",
			false,
		},
		{
			"negative zero units",
			&pb.MoneyValue{Currency: "eur", Units: 0, Nano: -500000000},
			"-0.50 EUR",
			false,
		},
		{
			"one cent",
			&pb.MoneyValue{Currency: "usd", Units: 0, Nano: 10000000},
			"0.01 USD",
			false,
		},
		{
			"mixed sign positive units negative nano",
			&pb.MoneyValue{Currency: "rub", Units: 1, Nano: -500000000},
			"",
			true,
		},
		{
			"mixed sign negative units positive nano",
			&pb.MoneyValue{Currency: "rub", Units: -1, Nano: 500000000},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FormatMoney(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				assert.ErrorIs(t, err, ErrClient)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFormatQuotation(t *testing.T) {
	tests := []struct {
		name    string
		input   *pb.Quotation
		want    string
		wantErr bool
	}{
		{
			"positive",
			&pb.Quotation{Units: 250, Nano: 500000000},
			"250.50",
			false,
		},
		{"zero nano", &pb.Quotation{Units: 100, Nano: 0}, "100", false},
		{"nil", nil, "", true},
		{
			"negative",
			&pb.Quotation{Units: -250, Nano: -500000000},
			"-250.50",
			false,
		},
		{
			"negative zero units",
			&pb.Quotation{Units: 0, Nano: -500000000},
			"-0.50",
			false,
		},
		{"one cent", &pb.Quotation{Units: 0, Nano: 10000000}, "0.01", false},
		{"mixed sign", &pb.Quotation{Units: 1, Nano: -500000000}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FormatQuotation(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				assert.ErrorIs(t, err, ErrClient)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
