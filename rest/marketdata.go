package rest

import "context"

type marketDataServiceClient struct{ c *Client }

const (
	pathMarketDataGetCandles         = "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetCandles"
	pathMarketDataGetClosePrices     = "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetClosePrices"
	pathMarketDataGetLastPrices      = "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetLastPrices"
	pathMarketDataGetLastTrades      = "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetLastTrades"
	pathMarketDataGetMarketValues    = "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetMarketValues"
	pathMarketDataGetOrderBook       = "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetOrderBook"
	pathMarketDataGetTechAnalysis    = "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetTechAnalysis"
	pathMarketDataGetTradingStatus   = "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetTradingStatus"
	pathMarketDataGetTradingStatuses = "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetTradingStatuses"
)

func (s *marketDataServiceClient) GetCandles(
	ctx context.Context, req *V1GetCandlesRequest,
) (*V1GetCandlesResponse, error) {
	return do[*V1GetCandlesResponse](ctx, s.c, pathMarketDataGetCandles, req)
}

func (s *marketDataServiceClient) GetClosePrices(
	ctx context.Context, req *V1GetClosePricesRequest,
) (*V1GetClosePricesResponse, error) {
	return do[*V1GetClosePricesResponse](
		ctx, s.c, pathMarketDataGetClosePrices, req,
	)
}

func (s *marketDataServiceClient) GetLastPrices(
	ctx context.Context, req *V1GetLastPricesRequest,
) (*V1GetLastPricesResponse, error) {
	return do[*V1GetLastPricesResponse](
		ctx, s.c, pathMarketDataGetLastPrices, req,
	)
}

func (s *marketDataServiceClient) GetLastTrades(
	ctx context.Context, req *V1GetLastTradesRequest,
) (*V1GetLastTradesResponse, error) {
	return do[*V1GetLastTradesResponse](
		ctx, s.c, pathMarketDataGetLastTrades, req,
	)
}

func (s *marketDataServiceClient) GetMarketValues(
	ctx context.Context, req *V1GetMarketValuesRequest,
) (*V1GetMarketValuesResponse, error) {
	return do[*V1GetMarketValuesResponse](
		ctx, s.c, pathMarketDataGetMarketValues, req,
	)
}

func (s *marketDataServiceClient) GetOrderBook(
	ctx context.Context, req *V1GetOrderBookRequest,
) (*V1GetOrderBookResponse, error) {
	return do[*V1GetOrderBookResponse](
		ctx,
		s.c,
		pathMarketDataGetOrderBook,
		req,
	)
}

func (s *marketDataServiceClient) GetTechAnalysis(
	ctx context.Context, req *V1GetTechAnalysisRequest,
) (*V1GetTechAnalysisResponse, error) {
	return do[*V1GetTechAnalysisResponse](
		ctx, s.c, pathMarketDataGetTechAnalysis, req,
	)
}

func (s *marketDataServiceClient) GetTradingStatus(
	ctx context.Context, req *V1GetTradingStatusRequest,
) (*V1GetTradingStatusResponse, error) {
	return do[*V1GetTradingStatusResponse](
		ctx, s.c, pathMarketDataGetTradingStatus, req,
	)
}

func (s *marketDataServiceClient) GetTradingStatuses(
	ctx context.Context, req *V1GetTradingStatusesRequest,
) (*V1GetTradingStatusesResponse, error) {
	return do[*V1GetTradingStatusesResponse](
		ctx, s.c, pathMarketDataGetTradingStatuses, req,
	)
}
