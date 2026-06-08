package rest

import "context"

type operationsServiceClient struct{ c *Client }

const (
	pathOperationsGetBrokerReport           = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetBrokerReport"
	pathOperationsGetDividendsForeignIssuer = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetDividendsForeignIssuer"
	pathOperationsGetOperations             = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetOperations"
	pathOperationsGetOperationsByCursor     = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetOperationsByCursor"
	pathOperationsGetPortfolio              = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetPortfolio"
	pathOperationsGetPositions              = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetPositions"
	pathOperationsGetWithdrawLimits         = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetWithdrawLimits"
)

func (s operationsServiceClient) GetBrokerReport(
	ctx context.Context, req *V1BrokerReportRequest,
) (*V1BrokerReportResponse, error) {
	return do[*V1BrokerReportResponse](
		ctx, s.c, pathOperationsGetBrokerReport, req,
	)
}

func (s operationsServiceClient) GetDividendsForeignIssuer(
	ctx context.Context, req *V1GetDividendsForeignIssuerRequest,
) (*V1GetDividendsForeignIssuerResponse, error) {
	return do[*V1GetDividendsForeignIssuerResponse](
		ctx, s.c, pathOperationsGetDividendsForeignIssuer, req,
	)
}

func (s operationsServiceClient) GetOperations(
	ctx context.Context, req *V1OperationsRequest,
) (*V1OperationsResponse, error) {
	return do[*V1OperationsResponse](
		ctx, s.c, pathOperationsGetOperations, req,
	)
}

func (s operationsServiceClient) GetOperationsByCursor(
	ctx context.Context, req *V1GetOperationsByCursorRequest,
) (*V1GetOperationsByCursorResponse, error) {
	return do[*V1GetOperationsByCursorResponse](
		ctx, s.c, pathOperationsGetOperationsByCursor, req,
	)
}

func (s operationsServiceClient) GetPortfolio(
	ctx context.Context, req *V1PortfolioRequest,
) (*V1PortfolioResponse, error) {
	return do[*V1PortfolioResponse](ctx, s.c, pathOperationsGetPortfolio, req)
}

func (s operationsServiceClient) GetPositions(
	ctx context.Context, req *V1PositionsRequest,
) (*V1PositionsResponse, error) {
	return do[*V1PositionsResponse](ctx, s.c, pathOperationsGetPositions, req)
}

func (s operationsServiceClient) GetWithdrawLimits(
	ctx context.Context, req *V1WithdrawLimitsRequest,
) (*V1WithdrawLimitsResponse, error) {
	return do[*V1WithdrawLimitsResponse](
		ctx, s.c, pathOperationsGetWithdrawLimits, req,
	)
}
