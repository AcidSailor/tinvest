package tinvest

import (
	pb "github.com/acidsailor/tinvest/pb"
)

// OrderDirectionFromStop maps a StopOrderDirection to the equivalent
// OrderDirection. Unknown values map to ORDER_DIRECTION_UNSPECIFIED.
//
// Used when feeding a stop order's direction into endpoints that take the
// regular OrderDirection enum (e.g. GetOrderPrice).
func OrderDirectionFromStop(d pb.StopOrderDirection) pb.OrderDirection {
	switch d {
	case pb.StopOrderDirection_STOP_ORDER_DIRECTION_BUY:
		return pb.OrderDirection_ORDER_DIRECTION_BUY
	case pb.StopOrderDirection_STOP_ORDER_DIRECTION_SELL:
		return pb.OrderDirection_ORDER_DIRECTION_SELL
	default:
		return pb.OrderDirection_ORDER_DIRECTION_UNSPECIFIED
	}
}
