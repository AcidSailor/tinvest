package rest

import "context"

type signalsServiceClient struct{ c *Client }

const (
	pathSignalGetSignals    = "/tinkoff.public.invest.api.contract.v1.SignalService/GetSignals"
	pathSignalGetStrategies = "/tinkoff.public.invest.api.contract.v1.SignalService/GetStrategies"
)

func (s signalsServiceClient) GetSignals(
	ctx context.Context, req *V1GetSignalsRequest,
) (*V1GetSignalsResponse, error) {
	return do[*V1GetSignalsResponse](ctx, s.c, pathSignalGetSignals, req)
}

func (s signalsServiceClient) GetStrategies(
	ctx context.Context, req *V1GetStrategiesRequest,
) (*V1GetStrategiesResponse, error) {
	return do[*V1GetStrategiesResponse](ctx, s.c, pathSignalGetStrategies, req)
}
