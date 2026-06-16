package grpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pb "github.com/acidsailor/tinvest/grpc/pb"
	"github.com/acidsailor/tinvest/money"
)

func TestFuturesPointValue(t *testing.T) {
	tests := []struct {
		name    string
		margin  *pb.GetFuturesMarginResponse
		units   int64
		nano    int32
		wantErr bool
	}{
		{
			"step=1, stepValue=0.2 → 0.2",
			&pb.GetFuturesMarginResponse{
				MinPriceIncrement:       &pb.Quotation{Units: 1},
				MinPriceIncrementAmount: &pb.Quotation{Nano: 200000000},
			},
			0,
			200000000,
			false,
		},
		{
			"nil margin",
			nil,
			0,
			0,
			true,
		},
		{
			"missing step",
			&pb.GetFuturesMarginResponse{
				MinPriceIncrementAmount: &pb.Quotation{Units: 1},
			},
			0,
			0,
			true,
		},
		{
			"zero step",
			&pb.GetFuturesMarginResponse{
				MinPriceIncrement:       &pb.Quotation{},
				MinPriceIncrementAmount: &pb.Quotation{Units: 1},
			},
			0,
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q, err := FuturesPointValue(tt.margin)
			if tt.wantErr {
				require.Error(t, err)
				m := tt.margin
				if m == nil ||
					m.GetMinPriceIncrement() == nil ||
					m.GetMinPriceIncrementAmount() == nil {
					assert.ErrorContains(t, err, "tinvest:")
				} else {
					assert.ErrorIs(t, err, money.ErrConversion)
				}
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.units, q.Units)
			assert.Equal(t, tt.nano, q.Nano)
		})
	}
}
