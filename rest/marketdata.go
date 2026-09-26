package rest

import "context"

type marketDataServiceClient struct{ c *Client }

const (
	pathMarketDataGetCandles endpoint[
		*V1GetCandlesRequest,
		*V1GetCandlesResponse,
	] = "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetCandles"

	pathMarketDataGetClosePrices endpoint[
		*V1GetClosePricesRequest,
		*V1GetClosePricesResponse,
	] = "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetClosePrices"

	pathMarketDataGetLastPrices endpoint[
		*V1GetLastPricesRequest,
		*V1GetLastPricesResponse,
	] = "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetLastPrices"

	pathMarketDataGetLastTrades endpoint[
		*V1GetLastTradesRequest,
		*V1GetLastTradesResponse,
	] = "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetLastTrades"

	pathMarketDataGetMarketValues endpoint[
		*V1GetMarketValuesRequest,
		*V1GetMarketValuesResponse,
	] = "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetMarketValues"

	pathMarketDataGetOrderBook endpoint[
		*V1GetOrderBookRequest,
		*V1GetOrderBookResponse,
	] = "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetOrderBook"

	pathMarketDataGetTechAnalysis endpoint[
		*V1GetTechAnalysisRequest,
		*V1GetTechAnalysisResponse,
	] = "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetTechAnalysis"

	pathMarketDataGetTradingStatus endpoint[
		*V1GetTradingStatusRequest,
		*V1GetTradingStatusResponse,
	] = "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetTradingStatus"

	pathMarketDataGetTradingStatuses endpoint[
		*V1GetTradingStatusesRequest,
		*V1GetTradingStatusesResponse,
	] = "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetTradingStatuses"
)

func (s *marketDataServiceClient) GetCandles(
	ctx context.Context, req *V1GetCandlesRequest,
) (*V1GetCandlesResponse, error) {
	return call(ctx, s.c, pathMarketDataGetCandles, req)
}

func (s *marketDataServiceClient) GetClosePrices(
	ctx context.Context, req *V1GetClosePricesRequest,
) (*V1GetClosePricesResponse, error) {
	return call(ctx, s.c, pathMarketDataGetClosePrices, req)
}

func (s *marketDataServiceClient) GetLastPrices(
	ctx context.Context, req *V1GetLastPricesRequest,
) (*V1GetLastPricesResponse, error) {
	return call(ctx, s.c, pathMarketDataGetLastPrices, req)
}

func (s *marketDataServiceClient) GetLastTrades(
	ctx context.Context, req *V1GetLastTradesRequest,
) (*V1GetLastTradesResponse, error) {
	return call(ctx, s.c, pathMarketDataGetLastTrades, req)
}

func (s *marketDataServiceClient) GetMarketValues(
	ctx context.Context, req *V1GetMarketValuesRequest,
) (*V1GetMarketValuesResponse, error) {
	return call(ctx, s.c, pathMarketDataGetMarketValues, req)
}

func (s *marketDataServiceClient) GetOrderBook(
	ctx context.Context, req *V1GetOrderBookRequest,
) (*V1GetOrderBookResponse, error) {
	return call(ctx, s.c, pathMarketDataGetOrderBook, req)
}

func (s *marketDataServiceClient) GetTechAnalysis(
	ctx context.Context, req *V1GetTechAnalysisRequest,
) (*V1GetTechAnalysisResponse, error) {
	return call(ctx, s.c, pathMarketDataGetTechAnalysis, req)
}

func (s *marketDataServiceClient) GetTradingStatus(
	ctx context.Context, req *V1GetTradingStatusRequest,
) (*V1GetTradingStatusResponse, error) {
	return call(ctx, s.c, pathMarketDataGetTradingStatus, req)
}

func (s *marketDataServiceClient) GetTradingStatuses(
	ctx context.Context, req *V1GetTradingStatusesRequest,
) (*V1GetTradingStatusesResponse, error) {
	return call(ctx, s.c, pathMarketDataGetTradingStatuses, req)
}
