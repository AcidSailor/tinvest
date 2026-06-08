package rest

import "context"

type instrumentsServiceClient struct{ c *Client }

const (
	pathInstrumentsBondBy                = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/BondBy"
	pathInstrumentsBonds                 = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Bonds"
	pathInstrumentsCreateFavoriteGroup   = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/CreateFavoriteGroup"
	pathInstrumentsCurrencies            = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Currencies"
	pathInstrumentsCurrencyBy            = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/CurrencyBy"
	pathInstrumentsDeleteFavoriteGroup   = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/DeleteFavoriteGroup"
	pathInstrumentsDfaBy                 = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/DfaBy"
	pathInstrumentsDfas                  = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Dfas"
	pathInstrumentsEditFavorites         = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/EditFavorites"
	pathInstrumentsEtfBy                 = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/EtfBy"
	pathInstrumentsEtfs                  = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Etfs"
	pathInstrumentsFindInstrument        = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/FindInstrument"
	pathInstrumentsFutureBy              = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/FutureBy"
	pathInstrumentsFutures               = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Futures"
	pathInstrumentsGetAccruedInterests   = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetAccruedInterests"
	pathInstrumentsGetAssetBy            = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetAssetBy"
	pathInstrumentsGetAssetFundamentals  = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetAssetFundamentals"
	pathInstrumentsGetAssetReports       = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetAssetReports"
	pathInstrumentsGetAssets             = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetAssets"
	pathInstrumentsGetBondCoupons        = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetBondCoupons"
	pathInstrumentsGetBondEvents         = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetBondEvents"
	pathInstrumentsGetBrandBy            = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetBrandBy"
	pathInstrumentsGetBrands             = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetBrands"
	pathInstrumentsGetConsensusForecasts = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetConsensusForecasts"
	pathInstrumentsGetCountries          = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetCountries"
	pathInstrumentsGetDividends          = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetDividends"
	pathInstrumentsGetFavoriteGroups     = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetFavoriteGroups"
	pathInstrumentsGetFavorites          = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetFavorites"
	pathInstrumentsGetForecastBy         = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetForecastBy"
	pathInstrumentsGetFuturesMargin      = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetFuturesMargin"
	pathInstrumentsGetInsiderDeals       = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetInsiderDeals"
	pathInstrumentsGetInstrumentBy       = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetInstrumentBy"
	pathInstrumentsGetRiskRates          = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetRiskRates"
	pathInstrumentsIndicatives           = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Indicatives"
	pathInstrumentsNews                  = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/News"
	pathInstrumentsOptionBy              = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/OptionBy"
	pathInstrumentsOptions               = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Options"
	pathInstrumentsOptionsBy             = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/OptionsBy"
	pathInstrumentsShareBy               = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/ShareBy"
	pathInstrumentsShares                = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Shares"
	pathInstrumentsStructuredNoteBy      = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/StructuredNoteBy"
	pathInstrumentsStructuredNotes       = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/StructuredNotes"
	pathInstrumentsTradingSchedules      = "/tinkoff.public.invest.api.contract.v1.InstrumentsService/TradingSchedules"
)

func (s instrumentsServiceClient) BondBy(
	ctx context.Context, req *V1InstrumentRequest,
) (*V1BondResponse, error) {
	return do[*V1BondResponse](ctx, s.c, pathInstrumentsBondBy, req)
}

func (s instrumentsServiceClient) Bonds(
	ctx context.Context, req *V1InstrumentsRequest,
) (*V1BondsResponse, error) {
	return do[*V1BondsResponse](ctx, s.c, pathInstrumentsBonds, req)
}

func (s instrumentsServiceClient) CreateFavoriteGroup(
	ctx context.Context, req *V1CreateFavoriteGroupRequest,
) (*V1CreateFavoriteGroupResponse, error) {
	return do[*V1CreateFavoriteGroupResponse](
		ctx, s.c, pathInstrumentsCreateFavoriteGroup, req,
	)
}

func (s instrumentsServiceClient) Currencies(
	ctx context.Context, req *V1InstrumentsRequest,
) (*V1CurrenciesResponse, error) {
	return do[*V1CurrenciesResponse](ctx, s.c, pathInstrumentsCurrencies, req)
}

func (s instrumentsServiceClient) CurrencyBy(
	ctx context.Context, req *V1InstrumentRequest,
) (*V1CurrencyResponse, error) {
	return do[*V1CurrencyResponse](ctx, s.c, pathInstrumentsCurrencyBy, req)
}

func (s instrumentsServiceClient) DeleteFavoriteGroup(
	ctx context.Context, req *V1DeleteFavoriteGroupRequest,
) (*V1DeleteFavoriteGroupResponse, error) {
	return do[*V1DeleteFavoriteGroupResponse](
		ctx, s.c, pathInstrumentsDeleteFavoriteGroup, req,
	)
}

func (s instrumentsServiceClient) DfaBy(
	ctx context.Context, req *V1InstrumentRequest,
) (*V1DfaResponse, error) {
	return do[*V1DfaResponse](ctx, s.c, pathInstrumentsDfaBy, req)
}

func (s instrumentsServiceClient) Dfas(
	ctx context.Context, req *V1DfasRequest,
) (*V1DfasResponse, error) {
	return do[*V1DfasResponse](ctx, s.c, pathInstrumentsDfas, req)
}

func (s instrumentsServiceClient) EditFavorites(
	ctx context.Context, req *V1EditFavoritesRequest,
) (*V1EditFavoritesResponse, error) {
	return do[*V1EditFavoritesResponse](
		ctx,
		s.c,
		pathInstrumentsEditFavorites,
		req,
	)
}

func (s instrumentsServiceClient) EtfBy(
	ctx context.Context, req *V1InstrumentRequest,
) (*V1EtfResponse, error) {
	return do[*V1EtfResponse](ctx, s.c, pathInstrumentsEtfBy, req)
}

func (s instrumentsServiceClient) Etfs(
	ctx context.Context, req *V1InstrumentsRequest,
) (*V1EtfsResponse, error) {
	return do[*V1EtfsResponse](ctx, s.c, pathInstrumentsEtfs, req)
}

func (s instrumentsServiceClient) FindInstrument(
	ctx context.Context, req *V1FindInstrumentRequest,
) (*V1FindInstrumentResponse, error) {
	return do[*V1FindInstrumentResponse](
		ctx, s.c, pathInstrumentsFindInstrument, req,
	)
}

func (s instrumentsServiceClient) FutureBy(
	ctx context.Context, req *V1InstrumentRequest,
) (*V1FutureResponse, error) {
	return do[*V1FutureResponse](ctx, s.c, pathInstrumentsFutureBy, req)
}

func (s instrumentsServiceClient) Futures(
	ctx context.Context, req *V1InstrumentsRequest,
) (*V1FuturesResponse, error) {
	return do[*V1FuturesResponse](ctx, s.c, pathInstrumentsFutures, req)
}

func (s instrumentsServiceClient) GetAccruedInterests(
	ctx context.Context, req *V1GetAccruedInterestsRequest,
) (*V1GetAccruedInterestsResponse, error) {
	return do[*V1GetAccruedInterestsResponse](
		ctx, s.c, pathInstrumentsGetAccruedInterests, req,
	)
}

func (s instrumentsServiceClient) GetAssetBy(
	ctx context.Context, req *V1AssetRequest,
) (*V1AssetResponse, error) {
	return do[*V1AssetResponse](ctx, s.c, pathInstrumentsGetAssetBy, req)
}

func (s instrumentsServiceClient) GetAssetFundamentals(
	ctx context.Context, req *V1GetAssetFundamentalsRequest,
) (*V1GetAssetFundamentalsResponse, error) {
	return do[*V1GetAssetFundamentalsResponse](
		ctx, s.c, pathInstrumentsGetAssetFundamentals, req,
	)
}

func (s instrumentsServiceClient) GetAssetReports(
	ctx context.Context, req *V1GetAssetReportsRequest,
) (*V1GetAssetReportsResponse, error) {
	return do[*V1GetAssetReportsResponse](
		ctx, s.c, pathInstrumentsGetAssetReports, req,
	)
}

func (s instrumentsServiceClient) GetAssets(
	ctx context.Context, req *V1AssetsRequest,
) (*V1AssetsResponse, error) {
	return do[*V1AssetsResponse](ctx, s.c, pathInstrumentsGetAssets, req)
}

func (s instrumentsServiceClient) GetBondCoupons(
	ctx context.Context, req *V1GetBondCouponsRequest,
) (*V1GetBondCouponsResponse, error) {
	return do[*V1GetBondCouponsResponse](
		ctx, s.c, pathInstrumentsGetBondCoupons, req,
	)
}

func (s instrumentsServiceClient) GetBondEvents(
	ctx context.Context, req *V1GetBondEventsRequest,
) (*V1GetBondEventsResponse, error) {
	return do[*V1GetBondEventsResponse](
		ctx, s.c, pathInstrumentsGetBondEvents, req,
	)
}

func (s instrumentsServiceClient) GetBrandBy(
	ctx context.Context, req *V1GetBrandRequest,
) (*V1Brand, error) {
	return do[*V1Brand](ctx, s.c, pathInstrumentsGetBrandBy, req)
}

func (s instrumentsServiceClient) GetBrands(
	ctx context.Context, req *V1GetBrandsRequest,
) (*V1GetBrandsResponse, error) {
	return do[*V1GetBrandsResponse](ctx, s.c, pathInstrumentsGetBrands, req)
}

func (s instrumentsServiceClient) GetConsensusForecasts(
	ctx context.Context, req *V1GetConsensusForecastsRequest,
) (*V1GetConsensusForecastsResponse, error) {
	return do[*V1GetConsensusForecastsResponse](
		ctx, s.c, pathInstrumentsGetConsensusForecasts, req,
	)
}

func (s instrumentsServiceClient) GetCountries(
	ctx context.Context, req *V1GetCountriesRequest,
) (*V1GetCountriesResponse, error) {
	return do[*V1GetCountriesResponse](
		ctx,
		s.c,
		pathInstrumentsGetCountries,
		req,
	)
}

func (s instrumentsServiceClient) GetDividends(
	ctx context.Context, req *V1GetDividendsRequest,
) (*V1GetDividendsResponse, error) {
	return do[*V1GetDividendsResponse](
		ctx,
		s.c,
		pathInstrumentsGetDividends,
		req,
	)
}

func (s instrumentsServiceClient) GetFavoriteGroups(
	ctx context.Context, req *V1GetFavoriteGroupsRequest,
) (*V1GetFavoriteGroupsResponse, error) {
	return do[*V1GetFavoriteGroupsResponse](
		ctx, s.c, pathInstrumentsGetFavoriteGroups, req,
	)
}

func (s instrumentsServiceClient) GetFavorites(
	ctx context.Context, req *V1GetFavoritesRequest,
) (*V1GetFavoritesResponse, error) {
	return do[*V1GetFavoritesResponse](
		ctx,
		s.c,
		pathInstrumentsGetFavorites,
		req,
	)
}

func (s instrumentsServiceClient) GetForecastBy(
	ctx context.Context, req *V1GetForecastRequest,
) (*V1GetForecastResponse, error) {
	return do[*V1GetForecastResponse](
		ctx,
		s.c,
		pathInstrumentsGetForecastBy,
		req,
	)
}

func (s instrumentsServiceClient) GetFuturesMargin(
	ctx context.Context, req *V1GetFuturesMarginRequest,
) (*V1GetFuturesMarginResponse, error) {
	return do[*V1GetFuturesMarginResponse](
		ctx, s.c, pathInstrumentsGetFuturesMargin, req,
	)
}

func (s instrumentsServiceClient) GetInsiderDeals(
	ctx context.Context, req *V1GetInsiderDealsRequest,
) (*V1GetInsiderDealsResponse, error) {
	return do[*V1GetInsiderDealsResponse](
		ctx, s.c, pathInstrumentsGetInsiderDeals, req,
	)
}

func (s instrumentsServiceClient) GetInstrumentBy(
	ctx context.Context, req *V1InstrumentRequest,
) (*V1InstrumentResponse, error) {
	return do[*V1InstrumentResponse](
		ctx, s.c, pathInstrumentsGetInstrumentBy, req,
	)
}

func (s instrumentsServiceClient) GetRiskRates(
	ctx context.Context, req *V1RiskRatesRequest,
) (*V1RiskRatesResponse, error) {
	return do[*V1RiskRatesResponse](ctx, s.c, pathInstrumentsGetRiskRates, req)
}

func (s instrumentsServiceClient) Indicatives(
	ctx context.Context, req *V1IndicativesRequest,
) (*V1IndicativesResponse, error) {
	return do[*V1IndicativesResponse](ctx, s.c, pathInstrumentsIndicatives, req)
}

func (s instrumentsServiceClient) News(
	ctx context.Context, req *V1NewsRequest,
) (*V1NewsResponse, error) {
	return do[*V1NewsResponse](ctx, s.c, pathInstrumentsNews, req)
}

func (s instrumentsServiceClient) OptionBy(
	ctx context.Context, req *V1InstrumentRequest,
) (*V1OptionResponse, error) {
	return do[*V1OptionResponse](ctx, s.c, pathInstrumentsOptionBy, req)
}

func (s instrumentsServiceClient) Options(
	ctx context.Context, req *V1InstrumentsRequest,
) (*V1OptionsResponse, error) {
	return do[*V1OptionsResponse](ctx, s.c, pathInstrumentsOptions, req)
}

func (s instrumentsServiceClient) OptionsBy(
	ctx context.Context, req *V1FilterOptionsRequest,
) (*V1OptionsResponse, error) {
	return do[*V1OptionsResponse](ctx, s.c, pathInstrumentsOptionsBy, req)
}

func (s instrumentsServiceClient) ShareBy(
	ctx context.Context, req *V1InstrumentRequest,
) (*V1ShareResponse, error) {
	return do[*V1ShareResponse](ctx, s.c, pathInstrumentsShareBy, req)
}

func (s instrumentsServiceClient) Shares(
	ctx context.Context, req *V1InstrumentsRequest,
) (*V1SharesResponse, error) {
	return do[*V1SharesResponse](ctx, s.c, pathInstrumentsShares, req)
}

func (s instrumentsServiceClient) StructuredNoteBy(
	ctx context.Context, req *V1InstrumentRequest,
) (*V1StructuredNoteResponse, error) {
	return do[*V1StructuredNoteResponse](
		ctx, s.c, pathInstrumentsStructuredNoteBy, req,
	)
}

func (s instrumentsServiceClient) StructuredNotes(
	ctx context.Context, req *V1InstrumentsRequest,
) (*V1StructuredNotesResponse, error) {
	return do[*V1StructuredNotesResponse](
		ctx, s.c, pathInstrumentsStructuredNotes, req,
	)
}

func (s instrumentsServiceClient) TradingSchedules(
	ctx context.Context, req *V1TradingSchedulesRequest,
) (*V1TradingSchedulesResponse, error) {
	return do[*V1TradingSchedulesResponse](
		ctx, s.c, pathInstrumentsTradingSchedules, req,
	)
}
