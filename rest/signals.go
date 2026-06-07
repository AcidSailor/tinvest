package rest

import "context"

type signalsService struct{ c *Client }

const (
	pathSignalGetSignals    = "/tinkoff.public.invest.api.contract.v1.SignalService/GetSignals"
	pathSignalGetStrategies = "/tinkoff.public.invest.api.contract.v1.SignalService/GetStrategies"
)

func (s signalsService) GetSignals(
	ctx context.Context, req *V1GetSignalsRequest,
) (*V1GetSignalsResponse, error) {
	return do[V1GetSignalsResponse](ctx, s.c, pathSignalGetSignals, req)
}

func (s signalsService) GetStrategies(
	ctx context.Context, req *V1GetStrategiesRequest,
) (*V1GetStrategiesResponse, error) {
	return do[V1GetStrategiesResponse](ctx, s.c, pathSignalGetStrategies, req)
}
