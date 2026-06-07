package rest

import "context"

type operationsService struct{ c *Client }

const (
	pathOperationsGetBrokerReport           = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetBrokerReport"
	pathOperationsGetDividendsForeignIssuer = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetDividendsForeignIssuer"
	pathOperationsGetOperations             = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetOperations"
	pathOperationsGetOperationsByCursor     = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetOperationsByCursor"
	pathOperationsGetPortfolio              = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetPortfolio"
	pathOperationsGetPositions              = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetPositions"
	pathOperationsGetWithdrawLimits         = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetWithdrawLimits"
)

func (s operationsService) GetBrokerReport(
	ctx context.Context, req *V1BrokerReportRequest,
) (*V1BrokerReportResponse, error) {
	return do[V1BrokerReportResponse](
		ctx, s.c, pathOperationsGetBrokerReport, req,
	)
}

func (s operationsService) GetDividendsForeignIssuer(
	ctx context.Context, req *V1GetDividendsForeignIssuerRequest,
) (*V1GetDividendsForeignIssuerResponse, error) {
	return do[V1GetDividendsForeignIssuerResponse](
		ctx, s.c, pathOperationsGetDividendsForeignIssuer, req,
	)
}

func (s operationsService) GetOperations(
	ctx context.Context, req *V1OperationsRequest,
) (*V1OperationsResponse, error) {
	return do[V1OperationsResponse](
		ctx, s.c, pathOperationsGetOperations, req,
	)
}

func (s operationsService) GetOperationsByCursor(
	ctx context.Context, req *V1GetOperationsByCursorRequest,
) (*V1GetOperationsByCursorResponse, error) {
	return do[V1GetOperationsByCursorResponse](
		ctx, s.c, pathOperationsGetOperationsByCursor, req,
	)
}

func (s operationsService) GetPortfolio(
	ctx context.Context, req *V1PortfolioRequest,
) (*V1PortfolioResponse, error) {
	return do[V1PortfolioResponse](ctx, s.c, pathOperationsGetPortfolio, req)
}

func (s operationsService) GetPositions(
	ctx context.Context, req *V1PositionsRequest,
) (*V1PositionsResponse, error) {
	return do[V1PositionsResponse](ctx, s.c, pathOperationsGetPositions, req)
}

func (s operationsService) GetWithdrawLimits(
	ctx context.Context, req *V1WithdrawLimitsRequest,
) (*V1WithdrawLimitsResponse, error) {
	return do[V1WithdrawLimitsResponse](
		ctx, s.c, pathOperationsGetWithdrawLimits, req,
	)
}
