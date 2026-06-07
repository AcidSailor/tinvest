package rest

import "context"

type stopOrdersServiceClient struct{ c *Client }

const (
	pathStopOrdersCancelStopOrder = "/tinkoff.public.invest.api.contract.v1.StopOrdersService/CancelStopOrder"
	pathStopOrdersGetStopOrders   = "/tinkoff.public.invest.api.contract.v1.StopOrdersService/GetStopOrders"
	pathStopOrdersPostStopOrder   = "/tinkoff.public.invest.api.contract.v1.StopOrdersService/PostStopOrder"
)

func (s stopOrdersServiceClient) CancelStopOrder(
	ctx context.Context, req *V1CancelStopOrderRequest,
) (*V1CancelStopOrderResponse, error) {
	return do[V1CancelStopOrderResponse](
		ctx, s.c, pathStopOrdersCancelStopOrder, req,
	)
}

func (s stopOrdersServiceClient) GetStopOrders(
	ctx context.Context, req *V1GetStopOrdersRequest,
) (*V1GetStopOrdersResponse, error) {
	return do[V1GetStopOrdersResponse](
		ctx, s.c, pathStopOrdersGetStopOrders, req,
	)
}

func (s stopOrdersServiceClient) PostStopOrder(
	ctx context.Context, req *V1PostStopOrderRequest,
) (*V1PostStopOrderResponse, error) {
	return do[V1PostStopOrderResponse](
		ctx, s.c, pathStopOrdersPostStopOrder, req,
	)
}
