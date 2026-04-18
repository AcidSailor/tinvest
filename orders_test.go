package tinvest

import (
	"testing"

	"github.com/stretchr/testify/assert"

	pb "github.com/acidsailor/tinvest/pb"
)

func TestOrderDirectionFromStop(t *testing.T) {
	tests := []struct {
		name string
		in   pb.StopOrderDirection
		want pb.OrderDirection
	}{
		{
			"buy",
			pb.StopOrderDirection_STOP_ORDER_DIRECTION_BUY,
			pb.OrderDirection_ORDER_DIRECTION_BUY,
		},
		{
			"sell",
			pb.StopOrderDirection_STOP_ORDER_DIRECTION_SELL,
			pb.OrderDirection_ORDER_DIRECTION_SELL,
		},
		{
			"unspecified",
			pb.StopOrderDirection_STOP_ORDER_DIRECTION_UNSPECIFIED,
			pb.OrderDirection_ORDER_DIRECTION_UNSPECIFIED,
		},
		{
			"unknown",
			pb.StopOrderDirection(99),
			pb.OrderDirection_ORDER_DIRECTION_UNSPECIFIED,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, OrderDirectionFromStop(tt.in))
		})
	}
}
