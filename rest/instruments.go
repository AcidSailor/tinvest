package rest

import "context"

type instrumentsServiceClient struct{ c *Client }

const (
	pathInstrumentsBondBy endpoint[
		*V1InstrumentRequest,
		*V1BondResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/BondBy"

	pathInstrumentsBonds endpoint[
		*V1InstrumentsRequest,
		*V1BondsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Bonds"

	pathInstrumentsCreateFavoriteGroup endpoint[
		*V1CreateFavoriteGroupRequest,
		*V1CreateFavoriteGroupResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/CreateFavoriteGroup"

	pathInstrumentsCurrencies endpoint[
		*V1InstrumentsRequest,
		*V1CurrenciesResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Currencies"

	pathInstrumentsCurrencyBy endpoint[
		*V1InstrumentRequest,
		*V1CurrencyResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/CurrencyBy"

	pathInstrumentsDeleteFavoriteGroup endpoint[
		*V1DeleteFavoriteGroupRequest,
		*V1DeleteFavoriteGroupResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/DeleteFavoriteGroup"

	pathInstrumentsDfaBy endpoint[
		*V1InstrumentRequest,
		*V1DfaResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/DfaBy"

	pathInstrumentsDfas endpoint[
		*V1DfasRequest,
		*V1DfasResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Dfas"

	pathInstrumentsEditFavorites endpoint[
		*V1EditFavoritesRequest,
		*V1EditFavoritesResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/EditFavorites"

	pathInstrumentsEtfBy endpoint[
		*V1InstrumentRequest,
		*V1EtfResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/EtfBy"

	pathInstrumentsEtfs endpoint[
		*V1InstrumentsRequest,
		*V1EtfsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Etfs"

	pathInstrumentsFindInstrument endpoint[
		*V1FindInstrumentRequest,
		*V1FindInstrumentResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/FindInstrument"

	pathInstrumentsFutureBy endpoint[
		*V1InstrumentRequest,
		*V1FutureResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/FutureBy"

	pathInstrumentsFutures endpoint[
		*V1InstrumentsRequest,
		*V1FuturesResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Futures"

	pathInstrumentsGetAccruedInterests endpoint[
		*V1GetAccruedInterestsRequest,
		*V1GetAccruedInterestsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetAccruedInterests"

	pathInstrumentsGetAssetBy endpoint[
		*V1AssetRequest,
		*V1AssetResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetAssetBy"

	pathInstrumentsGetAssetFundamentals endpoint[
		*V1GetAssetFundamentalsRequest,
		*V1GetAssetFundamentalsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetAssetFundamentals"

	pathInstrumentsGetAssetReports endpoint[
		*V1GetAssetReportsRequest,
		*V1GetAssetReportsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetAssetReports"

	pathInstrumentsGetAssets endpoint[
		*V1AssetsRequest,
		*V1AssetsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetAssets"

	pathInstrumentsGetBondCoupons endpoint[
		*V1GetBondCouponsRequest,
		*V1GetBondCouponsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetBondCoupons"

	pathInstrumentsGetBondEvents endpoint[
		*V1GetBondEventsRequest,
		*V1GetBondEventsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetBondEvents"

	pathInstrumentsGetBrandBy endpoint[
		*V1GetBrandRequest,
		*V1Brand,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetBrandBy"

	pathInstrumentsGetBrands endpoint[
		*V1GetBrandsRequest,
		*V1GetBrandsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetBrands"

	pathInstrumentsGetConsensusForecasts endpoint[
		*V1GetConsensusForecastsRequest,
		*V1GetConsensusForecastsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetConsensusForecasts"

	pathInstrumentsGetCountries endpoint[
		*V1GetCountriesRequest,
		*V1GetCountriesResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetCountries"

	pathInstrumentsGetDividends endpoint[
		*V1GetDividendsRequest,
		*V1GetDividendsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetDividends"

	pathInstrumentsGetFavoriteGroups endpoint[
		*V1GetFavoriteGroupsRequest,
		*V1GetFavoriteGroupsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetFavoriteGroups"

	pathInstrumentsGetFavorites endpoint[
		*V1GetFavoritesRequest,
		*V1GetFavoritesResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetFavorites"

	pathInstrumentsGetForecastBy endpoint[
		*V1GetForecastRequest,
		*V1GetForecastResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetForecastBy"

	pathInstrumentsGetFuturesMargin endpoint[
		*V1GetFuturesMarginRequest,
		*V1GetFuturesMarginResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetFuturesMargin"

	pathInstrumentsGetInsiderDeals endpoint[
		*V1GetInsiderDealsRequest,
		*V1GetInsiderDealsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetInsiderDeals"

	pathInstrumentsGetInstrumentBy endpoint[
		*V1InstrumentRequest,
		*V1InstrumentResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetInstrumentBy"

	pathInstrumentsGetRiskRates endpoint[
		*V1RiskRatesRequest,
		*V1RiskRatesResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetRiskRates"

	pathInstrumentsIndicatives endpoint[
		*V1IndicativesRequest,
		*V1IndicativesResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Indicatives"

	pathInstrumentsNews endpoint[
		*V1NewsRequest,
		*V1NewsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/News"

	pathInstrumentsOptionBy endpoint[
		*V1InstrumentRequest,
		*V1OptionResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/OptionBy"

	pathInstrumentsOptions endpoint[
		*V1InstrumentsRequest,
		*V1OptionsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Options"

	pathInstrumentsOptionsBy endpoint[
		*V1FilterOptionsRequest,
		*V1OptionsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/OptionsBy"

	pathInstrumentsShareBy endpoint[
		*V1InstrumentRequest,
		*V1ShareResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/ShareBy"

	pathInstrumentsShares endpoint[
		*V1InstrumentsRequest,
		*V1SharesResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Shares"

	pathInstrumentsStructuredNoteBy endpoint[
		*V1InstrumentRequest,
		*V1StructuredNoteResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/StructuredNoteBy"

	pathInstrumentsStructuredNotes endpoint[
		*V1InstrumentsRequest,
		*V1StructuredNotesResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/StructuredNotes"

	pathInstrumentsTradingSchedules endpoint[
		*V1TradingSchedulesRequest,
		*V1TradingSchedulesResponse,
	] = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/TradingSchedules"
)

func (s *instrumentsServiceClient) BondBy(
	ctx context.Context, req *V1InstrumentRequest,
) (*V1BondResponse, error) {
	return call(ctx, s.c, pathInstrumentsBondBy, req)
}

func (s *instrumentsServiceClient) Bonds(
	ctx context.Context, req *V1InstrumentsRequest,
) (*V1BondsResponse, error) {
	return call(ctx, s.c, pathInstrumentsBonds, req)
}

func (s *instrumentsServiceClient) CreateFavoriteGroup(
	ctx context.Context, req *V1CreateFavoriteGroupRequest,
) (*V1CreateFavoriteGroupResponse, error) {
	return call(ctx, s.c, pathInstrumentsCreateFavoriteGroup, req)
}

func (s *instrumentsServiceClient) Currencies(
	ctx context.Context, req *V1InstrumentsRequest,
) (*V1CurrenciesResponse, error) {
	return call(ctx, s.c, pathInstrumentsCurrencies, req)
}

func (s *instrumentsServiceClient) CurrencyBy(
	ctx context.Context, req *V1InstrumentRequest,
) (*V1CurrencyResponse, error) {
	return call(ctx, s.c, pathInstrumentsCurrencyBy, req)
}

func (s *instrumentsServiceClient) DeleteFavoriteGroup(
	ctx context.Context, req *V1DeleteFavoriteGroupRequest,
) (*V1DeleteFavoriteGroupResponse, error) {
	return call(ctx, s.c, pathInstrumentsDeleteFavoriteGroup, req)
}

func (s *instrumentsServiceClient) DfaBy(
	ctx context.Context, req *V1InstrumentRequest,
) (*V1DfaResponse, error) {
	return call(ctx, s.c, pathInstrumentsDfaBy, req)
}

func (s *instrumentsServiceClient) Dfas(
	ctx context.Context, req *V1DfasRequest,
) (*V1DfasResponse, error) {
	return call(ctx, s.c, pathInstrumentsDfas, req)
}

func (s *instrumentsServiceClient) EditFavorites(
	ctx context.Context, req *V1EditFavoritesRequest,
) (*V1EditFavoritesResponse, error) {
	return call(ctx, s.c, pathInstrumentsEditFavorites, req)
}

func (s *instrumentsServiceClient) EtfBy(
	ctx context.Context, req *V1InstrumentRequest,
) (*V1EtfResponse, error) {
	return call(ctx, s.c, pathInstrumentsEtfBy, req)
}

func (s *instrumentsServiceClient) Etfs(
	ctx context.Context, req *V1InstrumentsRequest,
) (*V1EtfsResponse, error) {
	return call(ctx, s.c, pathInstrumentsEtfs, req)
}

func (s *instrumentsServiceClient) FindInstrument(
	ctx context.Context, req *V1FindInstrumentRequest,
) (*V1FindInstrumentResponse, error) {
	return call(ctx, s.c, pathInstrumentsFindInstrument, req)
}

func (s *instrumentsServiceClient) FutureBy(
	ctx context.Context, req *V1InstrumentRequest,
) (*V1FutureResponse, error) {
	return call(ctx, s.c, pathInstrumentsFutureBy, req)
}

func (s *instrumentsServiceClient) Futures(
	ctx context.Context, req *V1InstrumentsRequest,
) (*V1FuturesResponse, error) {
	return call(ctx, s.c, pathInstrumentsFutures, req)
}

func (s *instrumentsServiceClient) GetAccruedInterests(
	ctx context.Context, req *V1GetAccruedInterestsRequest,
) (*V1GetAccruedInterestsResponse, error) {
	return call(ctx, s.c, pathInstrumentsGetAccruedInterests, req)
}

func (s *instrumentsServiceClient) GetAssetBy(
	ctx context.Context, req *V1AssetRequest,
) (*V1AssetResponse, error) {
	return call(ctx, s.c, pathInstrumentsGetAssetBy, req)
}

func (s *instrumentsServiceClient) GetAssetFundamentals(
	ctx context.Context, req *V1GetAssetFundamentalsRequest,
) (*V1GetAssetFundamentalsResponse, error) {
	return call(ctx, s.c, pathInstrumentsGetAssetFundamentals, req)
}

func (s *instrumentsServiceClient) GetAssetReports(
	ctx context.Context, req *V1GetAssetReportsRequest,
) (*V1GetAssetReportsResponse, error) {
	return call(ctx, s.c, pathInstrumentsGetAssetReports, req)
}

func (s *instrumentsServiceClient) GetAssets(
	ctx context.Context, req *V1AssetsRequest,
) (*V1AssetsResponse, error) {
	return call(ctx, s.c, pathInstrumentsGetAssets, req)
}

func (s *instrumentsServiceClient) GetBondCoupons(
	ctx context.Context, req *V1GetBondCouponsRequest,
) (*V1GetBondCouponsResponse, error) {
	return call(ctx, s.c, pathInstrumentsGetBondCoupons, req)
}

func (s *instrumentsServiceClient) GetBondEvents(
	ctx context.Context, req *V1GetBondEventsRequest,
) (*V1GetBondEventsResponse, error) {
	return call(ctx, s.c, pathInstrumentsGetBondEvents, req)
}

func (s *instrumentsServiceClient) GetBrandBy(
	ctx context.Context, req *V1GetBrandRequest,
) (*V1Brand, error) {
	return call(ctx, s.c, pathInstrumentsGetBrandBy, req)
}

func (s *instrumentsServiceClient) GetBrands(
	ctx context.Context, req *V1GetBrandsRequest,
) (*V1GetBrandsResponse, error) {
	return call(ctx, s.c, pathInstrumentsGetBrands, req)
}

func (s *instrumentsServiceClient) GetConsensusForecasts(
	ctx context.Context, req *V1GetConsensusForecastsRequest,
) (*V1GetConsensusForecastsResponse, error) {
	return call(ctx, s.c, pathInstrumentsGetConsensusForecasts, req)
}

func (s *instrumentsServiceClient) GetCountries(
	ctx context.Context, req *V1GetCountriesRequest,
) (*V1GetCountriesResponse, error) {
	return call(ctx, s.c, pathInstrumentsGetCountries, req)
}

func (s *instrumentsServiceClient) GetDividends(
	ctx context.Context, req *V1GetDividendsRequest,
) (*V1GetDividendsResponse, error) {
	return call(ctx, s.c, pathInstrumentsGetDividends, req)
}

func (s *instrumentsServiceClient) GetFavoriteGroups(
	ctx context.Context, req *V1GetFavoriteGroupsRequest,
) (*V1GetFavoriteGroupsResponse, error) {
	return call(ctx, s.c, pathInstrumentsGetFavoriteGroups, req)
}

func (s *instrumentsServiceClient) GetFavorites(
	ctx context.Context, req *V1GetFavoritesRequest,
) (*V1GetFavoritesResponse, error) {
	return call(ctx, s.c, pathInstrumentsGetFavorites, req)
}

func (s *instrumentsServiceClient) GetForecastBy(
	ctx context.Context, req *V1GetForecastRequest,
) (*V1GetForecastResponse, error) {
	return call(ctx, s.c, pathInstrumentsGetForecastBy, req)
}

func (s *instrumentsServiceClient) GetFuturesMargin(
	ctx context.Context, req *V1GetFuturesMarginRequest,
) (*V1GetFuturesMarginResponse, error) {
	return call(ctx, s.c, pathInstrumentsGetFuturesMargin, req)
}

func (s *instrumentsServiceClient) GetInsiderDeals(
	ctx context.Context, req *V1GetInsiderDealsRequest,
) (*V1GetInsiderDealsResponse, error) {
	return call(ctx, s.c, pathInstrumentsGetInsiderDeals, req)
}

func (s *instrumentsServiceClient) GetInstrumentBy(
	ctx context.Context, req *V1InstrumentRequest,
) (*V1InstrumentResponse, error) {
	return call(ctx, s.c, pathInstrumentsGetInstrumentBy, req)
}

func (s *instrumentsServiceClient) GetRiskRates(
	ctx context.Context, req *V1RiskRatesRequest,
) (*V1RiskRatesResponse, error) {
	return call(ctx, s.c, pathInstrumentsGetRiskRates, req)
}

func (s *instrumentsServiceClient) Indicatives(
	ctx context.Context, req *V1IndicativesRequest,
) (*V1IndicativesResponse, error) {
	return call(ctx, s.c, pathInstrumentsIndicatives, req)
}

func (s *instrumentsServiceClient) News(
	ctx context.Context, req *V1NewsRequest,
) (*V1NewsResponse, error) {
	return call(ctx, s.c, pathInstrumentsNews, req)
}

func (s *instrumentsServiceClient) OptionBy(
	ctx context.Context, req *V1InstrumentRequest,
) (*V1OptionResponse, error) {
	return call(ctx, s.c, pathInstrumentsOptionBy, req)
}

func (s *instrumentsServiceClient) Options(
	ctx context.Context, req *V1InstrumentsRequest,
) (*V1OptionsResponse, error) {
	return call(ctx, s.c, pathInstrumentsOptions, req)
}

func (s *instrumentsServiceClient) OptionsBy(
	ctx context.Context, req *V1FilterOptionsRequest,
) (*V1OptionsResponse, error) {
	return call(ctx, s.c, pathInstrumentsOptionsBy, req)
}

func (s *instrumentsServiceClient) ShareBy(
	ctx context.Context, req *V1InstrumentRequest,
) (*V1ShareResponse, error) {
	return call(ctx, s.c, pathInstrumentsShareBy, req)
}

func (s *instrumentsServiceClient) Shares(
	ctx context.Context, req *V1InstrumentsRequest,
) (*V1SharesResponse, error) {
	return call(ctx, s.c, pathInstrumentsShares, req)
}

func (s *instrumentsServiceClient) StructuredNoteBy(
	ctx context.Context, req *V1InstrumentRequest,
) (*V1StructuredNoteResponse, error) {
	return call(ctx, s.c, pathInstrumentsStructuredNoteBy, req)
}

func (s *instrumentsServiceClient) StructuredNotes(
	ctx context.Context, req *V1InstrumentsRequest,
) (*V1StructuredNotesResponse, error) {
	return call(ctx, s.c, pathInstrumentsStructuredNotes, req)
}

func (s *instrumentsServiceClient) TradingSchedules(
	ctx context.Context, req *V1TradingSchedulesRequest,
) (*V1TradingSchedulesResponse, error) {
	return call(ctx, s.c, pathInstrumentsTradingSchedules, req)
}
