package rest

import "context"

type stopOrdersServiceClient struct{ c *Client }

const (
	pathStopOrdersCancelStopOrder endpoint[
		*V1CancelStopOrderRequest,
		*V1CancelStopOrderResponse,
	] = "/tinkoff.public.invest.api.contract.v1.StopOrdersService/CancelStopOrder"

	pathStopOrdersGetStopOrders endpoint[
		*V1GetStopOrdersRequest,
		*V1GetStopOrdersResponse,
	] = "/tinkoff.public.invest.api.contract.v1.StopOrdersService/GetStopOrders"

	pathStopOrdersPostStopOrder endpoint[
		*V1PostStopOrderRequest,
		*V1PostStopOrderResponse,
	] = "/tinkoff.public.invest.api.contract.v1.StopOrdersService/PostStopOrder"
)

func (s *stopOrdersServiceClient) CancelStopOrder(
	ctx context.Context, req *V1CancelStopOrderRequest,
) (*V1CancelStopOrderResponse, error) {
	return call(ctx, s.c, pathStopOrdersCancelStopOrder, req)
}

func (s *stopOrdersServiceClient) GetStopOrders(
	ctx context.Context, req *V1GetStopOrdersRequest,
) (*V1GetStopOrdersResponse, error) {
	return call(ctx, s.c, pathStopOrdersGetStopOrders, req)
}

func (s *stopOrdersServiceClient) PostStopOrder(
	ctx context.Context, req *V1PostStopOrderRequest,
) (*V1PostStopOrderResponse, error) {
	return call(ctx, s.c, pathStopOrdersPostStopOrder, req)
}
