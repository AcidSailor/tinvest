package rest

import "context"

type operationsServiceClient struct{ c *Client }

const (
	pathOperationsGetBrokerReport endpoint[
		*V1BrokerReportRequest,
		*V1BrokerReportResponse,
	] = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetBrokerReport"

	pathOperationsGetDividendsForeignIssuer endpoint[
		*V1GetDividendsForeignIssuerRequest,
		*V1GetDividendsForeignIssuerResponse,
	] = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetDividendsForeignIssuer"

	pathOperationsGetOperations endpoint[
		*V1OperationsRequest,
		*V1OperationsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetOperations"

	pathOperationsGetOperationsByCursor endpoint[
		*V1GetOperationsByCursorRequest,
		*V1GetOperationsByCursorResponse,
	] = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetOperationsByCursor"

	pathOperationsGetPortfolio endpoint[
		*V1PortfolioRequest,
		*V1PortfolioResponse,
	] = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetPortfolio"

	pathOperationsGetPositions endpoint[
		*V1PositionsRequest,
		*V1PositionsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetPositions"

	pathOperationsGetWithdrawLimits endpoint[
		*V1WithdrawLimitsRequest,
		*V1WithdrawLimitsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.OperationsService/GetWithdrawLimits"
)

func (s *operationsServiceClient) GetBrokerReport(
	ctx context.Context, req *V1BrokerReportRequest,
) (*V1BrokerReportResponse, error) {
	return call(ctx, s.c, pathOperationsGetBrokerReport, req)
}

func (s *operationsServiceClient) GetDividendsForeignIssuer(
	ctx context.Context, req *V1GetDividendsForeignIssuerRequest,
) (*V1GetDividendsForeignIssuerResponse, error) {
	return call(ctx, s.c, pathOperationsGetDividendsForeignIssuer, req)
}

func (s *operationsServiceClient) GetOperations(
	ctx context.Context, req *V1OperationsRequest,
) (*V1OperationsResponse, error) {
	return call(ctx, s.c, pathOperationsGetOperations, req)
}

func (s *operationsServiceClient) GetOperationsByCursor(
	ctx context.Context, req *V1GetOperationsByCursorRequest,
) (*V1GetOperationsByCursorResponse, error) {
	return call(ctx, s.c, pathOperationsGetOperationsByCursor, req)
}

func (s *operationsServiceClient) GetPortfolio(
	ctx context.Context, req *V1PortfolioRequest,
) (*V1PortfolioResponse, error) {
	return call(ctx, s.c, pathOperationsGetPortfolio, req)
}

func (s *operationsServiceClient) GetPositions(
	ctx context.Context, req *V1PositionsRequest,
) (*V1PositionsResponse, error) {
	return call(ctx, s.c, pathOperationsGetPositions, req)
}

func (s *operationsServiceClient) GetWithdrawLimits(
	ctx context.Context, req *V1WithdrawLimitsRequest,
) (*V1WithdrawLimitsResponse, error) {
	return call(ctx, s.c, pathOperationsGetWithdrawLimits, req)
}
