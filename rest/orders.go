package rest

import "context"

type ordersServiceClient struct{ c *Client }

const (
	pathOrdersCancelOrder endpoint[
		*V1CancelOrderRequest,
		*V1CancelOrderResponse,
	] = "/tinkoff.public.invest.api.contract.v1.OrdersService/CancelOrder"

	pathOrdersGetMaxLots endpoint[
		*V1GetMaxLotsRequest,
		*V1GetMaxLotsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.OrdersService/GetMaxLots"

	pathOrdersGetOrderPrice endpoint[
		*V1GetOrderPriceRequest,
		*V1GetOrderPriceResponse,
	] = "/tinkoff.public.invest.api.contract.v1.OrdersService/GetOrderPrice"

	pathOrdersGetOrderState endpoint[
		*V1GetOrderStateRequest,
		*Contractv1OrderState,
	] = "/tinkoff.public.invest.api.contract.v1.OrdersService/GetOrderState"

	pathOrdersGetOrders endpoint[
		*V1GetOrdersRequest,
		*V1GetOrdersResponse,
	] = "/tinkoff.public.invest.api.contract.v1.OrdersService/GetOrders"

	pathOrdersPostOrder endpoint[
		*V1PostOrderRequest,
		*V1PostOrderResponse,
	] = "/tinkoff.public.invest.api.contract.v1.OrdersService/PostOrder"

	pathOrdersPostOrderAsync endpoint[
		*V1PostOrderAsyncRequest,
		*V1PostOrderAsyncResponse,
	] = "/tinkoff.public.invest.api.contract.v1.OrdersService/PostOrderAsync"

	pathOrdersReplaceOrder endpoint[
		*V1ReplaceOrderRequest,
		*V1PostOrderResponse,
	] = "/tinkoff.public.invest.api.contract.v1.OrdersService/ReplaceOrder"
)

func (s *ordersServiceClient) CancelOrder(
	ctx context.Context, req *V1CancelOrderRequest,
) (*V1CancelOrderResponse, error) {
	return call(ctx, s.c, pathOrdersCancelOrder, req)
}

func (s *ordersServiceClient) GetMaxLots(
	ctx context.Context, req *V1GetMaxLotsRequest,
) (*V1GetMaxLotsResponse, error) {
	return call(ctx, s.c, pathOrdersGetMaxLots, req)
}

func (s *ordersServiceClient) GetOrderPrice(
	ctx context.Context, req *V1GetOrderPriceRequest,
) (*V1GetOrderPriceResponse, error) {
	return call(ctx, s.c, pathOrdersGetOrderPrice, req)
}

func (s *ordersServiceClient) GetOrderState(
	ctx context.Context, req *V1GetOrderStateRequest,
) (*Contractv1OrderState, error) {
	return call(ctx, s.c, pathOrdersGetOrderState, req)
}

func (s *ordersServiceClient) GetOrders(
	ctx context.Context, req *V1GetOrdersRequest,
) (*V1GetOrdersResponse, error) {
	return call(ctx, s.c, pathOrdersGetOrders, req)
}

func (s *ordersServiceClient) PostOrder(
	ctx context.Context, req *V1PostOrderRequest,
) (*V1PostOrderResponse, error) {
	return call(ctx, s.c, pathOrdersPostOrder, req)
}

func (s *ordersServiceClient) PostOrderAsync(
	ctx context.Context, req *V1PostOrderAsyncRequest,
) (*V1PostOrderAsyncResponse, error) {
	return call(ctx, s.c, pathOrdersPostOrderAsync, req)
}

func (s *ordersServiceClient) ReplaceOrder(
	ctx context.Context, req *V1ReplaceOrderRequest,
) (*V1PostOrderResponse, error) {
	return call(ctx, s.c, pathOrdersReplaceOrder, req)
}
