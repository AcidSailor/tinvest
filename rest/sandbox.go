package rest

import "context"

type sandboxServiceClient struct{ c *Client }

const (
	pathSandboxCancelSandboxOrder endpoint[
		*V1CancelOrderRequest,
		*V1CancelOrderResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SandboxService/CancelSandboxOrder"

	pathSandboxCancelSandboxStopOrder endpoint[
		*V1CancelStopOrderRequest,
		*V1CancelStopOrderResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SandboxService/CancelSandboxStopOrder"

	pathSandboxCloseSandboxAccount endpoint[
		*V1CloseSandboxAccountRequest,
		*V1CloseSandboxAccountResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SandboxService/CloseSandboxAccount"

	pathSandboxGetSandboxAccounts endpoint[
		*V1GetAccountsRequest,
		*V1GetAccountsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxAccounts"

	pathSandboxGetSandboxMaxLots endpoint[
		*V1GetMaxLotsRequest,
		*V1GetMaxLotsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxMaxLots"

	pathSandboxGetSandboxOperations endpoint[
		*V1OperationsRequest,
		*V1OperationsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxOperations"

	pathSandboxGetSandboxOperationsByCursor endpoint[
		*V1GetOperationsByCursorRequest,
		*V1GetOperationsByCursorResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxOperationsByCursor"

	pathSandboxGetSandboxOrderPrice endpoint[
		*V1GetOrderPriceRequest,
		*V1GetOrderPriceResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxOrderPrice"

	pathSandboxGetSandboxOrderState endpoint[
		*V1GetOrderStateRequest,
		*Contractv1OrderState,
	] = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxOrderState"

	pathSandboxGetSandboxOrders endpoint[
		*V1GetOrdersRequest,
		*V1GetOrdersResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxOrders"

	pathSandboxGetSandboxPortfolio endpoint[
		*V1PortfolioRequest,
		*V1PortfolioResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxPortfolio"

	pathSandboxGetSandboxPositions endpoint[
		*V1PositionsRequest,
		*V1PositionsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxPositions"

	pathSandboxGetSandboxStopOrders endpoint[
		*V1GetStopOrdersRequest,
		*V1GetStopOrdersResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxStopOrders"

	pathSandboxGetSandboxWithdrawLimits endpoint[
		*V1WithdrawLimitsRequest,
		*V1WithdrawLimitsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxWithdrawLimits"

	pathSandboxOpenSandboxAccount endpoint[
		*V1OpenSandboxAccountRequest,
		*V1OpenSandboxAccountResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SandboxService/OpenSandboxAccount"

	pathSandboxPostSandboxOrder endpoint[
		*V1PostOrderRequest,
		*V1PostOrderResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SandboxService/PostSandboxOrder"

	pathSandboxPostSandboxOrderAsync endpoint[
		*V1PostOrderAsyncRequest,
		*V1PostOrderAsyncResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SandboxService/PostSandboxOrderAsync"

	pathSandboxPostSandboxStopOrder endpoint[
		*V1PostStopOrderRequest,
		*V1PostStopOrderResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SandboxService/PostSandboxStopOrder"

	pathSandboxReplaceSandboxOrder endpoint[
		*V1ReplaceOrderRequest,
		*V1PostOrderResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SandboxService/ReplaceSandboxOrder"

	pathSandboxSandboxPayIn endpoint[
		*V1SandboxPayInRequest,
		*V1SandboxPayInResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SandboxService/SandboxPayIn"
)

func (s *sandboxServiceClient) CancelSandboxOrder(
	ctx context.Context, req *V1CancelOrderRequest,
) (*V1CancelOrderResponse, error) {
	return call(ctx, s.c, pathSandboxCancelSandboxOrder, req)
}

func (s *sandboxServiceClient) CancelSandboxStopOrder(
	ctx context.Context, req *V1CancelStopOrderRequest,
) (*V1CancelStopOrderResponse, error) {
	return call(ctx, s.c, pathSandboxCancelSandboxStopOrder, req)
}

func (s *sandboxServiceClient) CloseSandboxAccount(
	ctx context.Context, req *V1CloseSandboxAccountRequest,
) (*V1CloseSandboxAccountResponse, error) {
	return call(ctx, s.c, pathSandboxCloseSandboxAccount, req)
}

func (s *sandboxServiceClient) GetSandboxAccounts(
	ctx context.Context, req *V1GetAccountsRequest,
) (*V1GetAccountsResponse, error) {
	return call(ctx, s.c, pathSandboxGetSandboxAccounts, req)
}

func (s *sandboxServiceClient) GetSandboxMaxLots(
	ctx context.Context, req *V1GetMaxLotsRequest,
) (*V1GetMaxLotsResponse, error) {
	return call(ctx, s.c, pathSandboxGetSandboxMaxLots, req)
}

func (s *sandboxServiceClient) GetSandboxOperations(
	ctx context.Context, req *V1OperationsRequest,
) (*V1OperationsResponse, error) {
	return call(ctx, s.c, pathSandboxGetSandboxOperations, req)
}

func (s *sandboxServiceClient) GetSandboxOperationsByCursor(
	ctx context.Context, req *V1GetOperationsByCursorRequest,
) (*V1GetOperationsByCursorResponse, error) {
	return call(ctx, s.c, pathSandboxGetSandboxOperationsByCursor, req)
}

func (s *sandboxServiceClient) GetSandboxOrderPrice(
	ctx context.Context, req *V1GetOrderPriceRequest,
) (*V1GetOrderPriceResponse, error) {
	return call(ctx, s.c, pathSandboxGetSandboxOrderPrice, req)
}

func (s *sandboxServiceClient) GetSandboxOrderState(
	ctx context.Context, req *V1GetOrderStateRequest,
) (*Contractv1OrderState, error) {
	return call(ctx, s.c, pathSandboxGetSandboxOrderState, req)
}

func (s *sandboxServiceClient) GetSandboxOrders(
	ctx context.Context, req *V1GetOrdersRequest,
) (*V1GetOrdersResponse, error) {
	return call(ctx, s.c, pathSandboxGetSandboxOrders, req)
}

func (s *sandboxServiceClient) GetSandboxPortfolio(
	ctx context.Context, req *V1PortfolioRequest,
) (*V1PortfolioResponse, error) {
	return call(ctx, s.c, pathSandboxGetSandboxPortfolio, req)
}

func (s *sandboxServiceClient) GetSandboxPositions(
	ctx context.Context, req *V1PositionsRequest,
) (*V1PositionsResponse, error) {
	return call(ctx, s.c, pathSandboxGetSandboxPositions, req)
}

func (s *sandboxServiceClient) GetSandboxStopOrders(
	ctx context.Context, req *V1GetStopOrdersRequest,
) (*V1GetStopOrdersResponse, error) {
	return call(ctx, s.c, pathSandboxGetSandboxStopOrders, req)
}

func (s *sandboxServiceClient) GetSandboxWithdrawLimits(
	ctx context.Context, req *V1WithdrawLimitsRequest,
) (*V1WithdrawLimitsResponse, error) {
	return call(ctx, s.c, pathSandboxGetSandboxWithdrawLimits, req)
}

func (s *sandboxServiceClient) OpenSandboxAccount(
	ctx context.Context, req *V1OpenSandboxAccountRequest,
) (*V1OpenSandboxAccountResponse, error) {
	return call(ctx, s.c, pathSandboxOpenSandboxAccount, req)
}

func (s *sandboxServiceClient) PostSandboxOrder(
	ctx context.Context, req *V1PostOrderRequest,
) (*V1PostOrderResponse, error) {
	return call(ctx, s.c, pathSandboxPostSandboxOrder, req)
}

func (s *sandboxServiceClient) PostSandboxOrderAsync(
	ctx context.Context, req *V1PostOrderAsyncRequest,
) (*V1PostOrderAsyncResponse, error) {
	return call(ctx, s.c, pathSandboxPostSandboxOrderAsync, req)
}

func (s *sandboxServiceClient) PostSandboxStopOrder(
	ctx context.Context, req *V1PostStopOrderRequest,
) (*V1PostStopOrderResponse, error) {
	return call(ctx, s.c, pathSandboxPostSandboxStopOrder, req)
}

func (s *sandboxServiceClient) ReplaceSandboxOrder(
	ctx context.Context, req *V1ReplaceOrderRequest,
) (*V1PostOrderResponse, error) {
	return call(ctx, s.c, pathSandboxReplaceSandboxOrder, req)
}

func (s *sandboxServiceClient) SandboxPayIn(
	ctx context.Context, req *V1SandboxPayInRequest,
) (*V1SandboxPayInResponse, error) {
	return call(ctx, s.c, pathSandboxSandboxPayIn, req)
}
