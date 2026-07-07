package rest

import "context"

type sandboxServiceClient struct{ c *Client }

const (
	pathSandboxCancelSandboxOrder           = "/tinkoff.public.invest.api.contract.v1.SandboxService/CancelSandboxOrder"
	pathSandboxCancelSandboxStopOrder       = "/tinkoff.public.invest.api.contract.v1.SandboxService/CancelSandboxStopOrder"
	pathSandboxCloseSandboxAccount          = "/tinkoff.public.invest.api.contract.v1.SandboxService/CloseSandboxAccount"
	pathSandboxGetSandboxAccounts           = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxAccounts"
	pathSandboxGetSandboxMaxLots            = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxMaxLots"
	pathSandboxGetSandboxOperations         = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxOperations"
	pathSandboxGetSandboxOperationsByCursor = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxOperationsByCursor"
	pathSandboxGetSandboxOrderPrice         = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxOrderPrice"
	pathSandboxGetSandboxOrderState         = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxOrderState"
	pathSandboxGetSandboxOrders             = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxOrders"
	pathSandboxGetSandboxPortfolio          = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxPortfolio"
	pathSandboxGetSandboxPositions          = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxPositions"
	pathSandboxGetSandboxStopOrders         = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxStopOrders"
	pathSandboxGetSandboxWithdrawLimits     = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxWithdrawLimits"
	pathSandboxOpenSandboxAccount           = "/tinkoff.public.invest.api.contract.v1.SandboxService/OpenSandboxAccount"
	pathSandboxPostSandboxOrder             = "/tinkoff.public.invest.api.contract.v1.SandboxService/PostSandboxOrder"
	pathSandboxPostSandboxOrderAsync        = "/tinkoff.public.invest.api.contract.v1.SandboxService/PostSandboxOrderAsync"
	pathSandboxPostSandboxStopOrder         = "/tinkoff.public.invest.api.contract.v1.SandboxService/PostSandboxStopOrder"
	pathSandboxReplaceSandboxOrder          = "/tinkoff.public.invest.api.contract.v1.SandboxService/ReplaceSandboxOrder"
	pathSandboxSandboxPayIn                 = "/tinkoff.public.invest.api.contract.v1.SandboxService/SandboxPayIn"
)

func (s *sandboxServiceClient) CancelSandboxOrder(
	ctx context.Context, req *V1CancelOrderRequest,
) (*V1CancelOrderResponse, error) {
	return do[*V1CancelOrderResponse](
		ctx, s.c, pathSandboxCancelSandboxOrder, req,
	)
}

func (s *sandboxServiceClient) CancelSandboxStopOrder(
	ctx context.Context, req *V1CancelStopOrderRequest,
) (*V1CancelStopOrderResponse, error) {
	return do[*V1CancelStopOrderResponse](
		ctx, s.c, pathSandboxCancelSandboxStopOrder, req,
	)
}

func (s *sandboxServiceClient) CloseSandboxAccount(
	ctx context.Context, req *V1CloseSandboxAccountRequest,
) (*V1CloseSandboxAccountResponse, error) {
	return do[*V1CloseSandboxAccountResponse](
		ctx, s.c, pathSandboxCloseSandboxAccount, req,
	)
}

func (s *sandboxServiceClient) GetSandboxAccounts(
	ctx context.Context, req *V1GetAccountsRequest,
) (*V1GetAccountsResponse, error) {
	return do[*V1GetAccountsResponse](
		ctx, s.c, pathSandboxGetSandboxAccounts, req,
	)
}

func (s *sandboxServiceClient) GetSandboxMaxLots(
	ctx context.Context, req *V1GetMaxLotsRequest,
) (*V1GetMaxLotsResponse, error) {
	return do[*V1GetMaxLotsResponse](
		ctx, s.c, pathSandboxGetSandboxMaxLots, req,
	)
}

func (s *sandboxServiceClient) GetSandboxOperations(
	ctx context.Context, req *V1OperationsRequest,
) (*V1OperationsResponse, error) {
	return do[*V1OperationsResponse](
		ctx, s.c, pathSandboxGetSandboxOperations, req,
	)
}

func (s *sandboxServiceClient) GetSandboxOperationsByCursor(
	ctx context.Context, req *V1GetOperationsByCursorRequest,
) (*V1GetOperationsByCursorResponse, error) {
	return do[*V1GetOperationsByCursorResponse](
		ctx, s.c, pathSandboxGetSandboxOperationsByCursor, req,
	)
}

func (s *sandboxServiceClient) GetSandboxOrderPrice(
	ctx context.Context, req *V1GetOrderPriceRequest,
) (*V1GetOrderPriceResponse, error) {
	return do[*V1GetOrderPriceResponse](
		ctx, s.c, pathSandboxGetSandboxOrderPrice, req,
	)
}

func (s *sandboxServiceClient) GetSandboxOrderState(
	ctx context.Context, req *V1GetOrderStateRequest,
) (*Contractv1OrderState, error) {
	return do[*Contractv1OrderState](
		ctx, s.c, pathSandboxGetSandboxOrderState, req,
	)
}

func (s *sandboxServiceClient) GetSandboxOrders(
	ctx context.Context, req *V1GetOrdersRequest,
) (*V1GetOrdersResponse, error) {
	return do[*V1GetOrdersResponse](ctx, s.c, pathSandboxGetSandboxOrders, req)
}

func (s *sandboxServiceClient) GetSandboxPortfolio(
	ctx context.Context, req *V1PortfolioRequest,
) (*V1PortfolioResponse, error) {
	return do[*V1PortfolioResponse](
		ctx, s.c, pathSandboxGetSandboxPortfolio, req,
	)
}

func (s *sandboxServiceClient) GetSandboxPositions(
	ctx context.Context, req *V1PositionsRequest,
) (*V1PositionsResponse, error) {
	return do[*V1PositionsResponse](
		ctx, s.c, pathSandboxGetSandboxPositions, req,
	)
}

func (s *sandboxServiceClient) GetSandboxStopOrders(
	ctx context.Context, req *V1GetStopOrdersRequest,
) (*V1GetStopOrdersResponse, error) {
	return do[*V1GetStopOrdersResponse](
		ctx, s.c, pathSandboxGetSandboxStopOrders, req,
	)
}

func (s *sandboxServiceClient) GetSandboxWithdrawLimits(
	ctx context.Context, req *V1WithdrawLimitsRequest,
) (*V1WithdrawLimitsResponse, error) {
	return do[*V1WithdrawLimitsResponse](
		ctx, s.c, pathSandboxGetSandboxWithdrawLimits, req,
	)
}

func (s *sandboxServiceClient) OpenSandboxAccount(
	ctx context.Context, req *V1OpenSandboxAccountRequest,
) (*V1OpenSandboxAccountResponse, error) {
	return do[*V1OpenSandboxAccountResponse](
		ctx, s.c, pathSandboxOpenSandboxAccount, req,
	)
}

func (s *sandboxServiceClient) PostSandboxOrder(
	ctx context.Context, req *V1PostOrderRequest,
) (*V1PostOrderResponse, error) {
	return do[*V1PostOrderResponse](ctx, s.c, pathSandboxPostSandboxOrder, req)
}

func (s *sandboxServiceClient) PostSandboxOrderAsync(
	ctx context.Context, req *V1PostOrderAsyncRequest,
) (*V1PostOrderAsyncResponse, error) {
	return do[*V1PostOrderAsyncResponse](
		ctx, s.c, pathSandboxPostSandboxOrderAsync, req,
	)
}

func (s *sandboxServiceClient) PostSandboxStopOrder(
	ctx context.Context, req *V1PostStopOrderRequest,
) (*V1PostStopOrderResponse, error) {
	return do[*V1PostStopOrderResponse](
		ctx, s.c, pathSandboxPostSandboxStopOrder, req,
	)
}

func (s *sandboxServiceClient) ReplaceSandboxOrder(
	ctx context.Context, req *V1ReplaceOrderRequest,
) (*V1PostOrderResponse, error) {
	return do[*V1PostOrderResponse](
		ctx, s.c, pathSandboxReplaceSandboxOrder, req,
	)
}

func (s *sandboxServiceClient) SandboxPayIn(
	ctx context.Context, req *V1SandboxPayInRequest,
) (*V1SandboxPayInResponse, error) {
	return do[*V1SandboxPayInResponse](ctx, s.c, pathSandboxSandboxPayIn, req)
}
