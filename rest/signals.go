package rest

import "context"

type signalsServiceClient struct{ c *Client }

const (
	pathSignalGetSignals endpoint[
		*V1GetSignalsRequest,
		*V1GetSignalsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SignalService/GetSignals"

	pathSignalGetStrategies endpoint[
		*V1GetStrategiesRequest,
		*V1GetStrategiesResponse,
	] = "/tinkoff.public.invest.api.contract.v1.SignalService/GetStrategies"
)

func (s *signalsServiceClient) GetSignals(
	ctx context.Context, req *V1GetSignalsRequest,
) (*V1GetSignalsResponse, error) {
	return call(ctx, s.c, pathSignalGetSignals, req)
}

func (s *signalsServiceClient) GetStrategies(
	ctx context.Context, req *V1GetStrategiesRequest,
) (*V1GetStrategiesResponse, error) {
	return call(ctx, s.c, pathSignalGetStrategies, req)
}
