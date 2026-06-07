package rest

import "context"

type ordersService struct{ c *Client }

const (
	pathOrdersCancelOrder    = "/tinkoff.public.invest.api.contract.v1.OrdersService/CancelOrder"
	pathOrdersGetMaxLots     = "/tinkoff.public.invest.api.contract.v1.OrdersService/GetMaxLots"
	pathOrdersGetOrderPrice  = "/tinkoff.public.invest.api.contract.v1.OrdersService/GetOrderPrice"
	pathOrdersGetOrderState  = "/tinkoff.public.invest.api.contract.v1.OrdersService/GetOrderState"
	pathOrdersGetOrders      = "/tinkoff.public.invest.api.contract.v1.OrdersService/GetOrders"
	pathOrdersPostOrder      = "/tinkoff.public.invest.api.contract.v1.OrdersService/PostOrder"
	pathOrdersPostOrderAsync = "/tinkoff.public.invest.api.contract.v1.OrdersService/PostOrderAsync"
	pathOrdersReplaceOrder   = "/tinkoff.public.invest.api.contract.v1.OrdersService/ReplaceOrder"
)

func (s ordersService) CancelOrder(
	ctx context.Context, req *V1CancelOrderRequest,
) (*V1CancelOrderResponse, error) {
	return do[V1CancelOrderResponse](ctx, s.c, pathOrdersCancelOrder, req)
}

func (s ordersService) GetMaxLots(
	ctx context.Context, req *V1GetMaxLotsRequest,
) (*V1GetMaxLotsResponse, error) {
	return do[V1GetMaxLotsResponse](ctx, s.c, pathOrdersGetMaxLots, req)
}

func (s ordersService) GetOrderPrice(
	ctx context.Context, req *V1GetOrderPriceRequest,
) (*V1GetOrderPriceResponse, error) {
	return do[V1GetOrderPriceResponse](ctx, s.c, pathOrdersGetOrderPrice, req)
}

func (s ordersService) GetOrderState(
	ctx context.Context, req *V1GetOrderStateRequest,
) (*Contractv1OrderState, error) {
	return do[Contractv1OrderState](ctx, s.c, pathOrdersGetOrderState, req)
}

func (s ordersService) GetOrders(
	ctx context.Context, req *V1GetOrdersRequest,
) (*V1GetOrdersResponse, error) {
	return do[V1GetOrdersResponse](ctx, s.c, pathOrdersGetOrders, req)
}

func (s ordersService) PostOrder(
	ctx context.Context, req *V1PostOrderRequest,
) (*V1PostOrderResponse, error) {
	return do[V1PostOrderResponse](ctx, s.c, pathOrdersPostOrder, req)
}

func (s ordersService) PostOrderAsync(
	ctx context.Context, req *V1PostOrderAsyncRequest,
) (*V1PostOrderAsyncResponse, error) {
	return do[V1PostOrderAsyncResponse](ctx, s.c, pathOrdersPostOrderAsync, req)
}

func (s ordersService) ReplaceOrder(
	ctx context.Context, req *V1ReplaceOrderRequest,
) (*V1PostOrderResponse, error) {
	return do[V1PostOrderResponse](ctx, s.c, pathOrdersReplaceOrder, req)
}
