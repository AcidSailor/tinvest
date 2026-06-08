package rest

import "context"

type usersServiceClient struct{ c *Client }

const (
	pathUsersCurrencyTransfer    = "/tinkoff.public.invest.api.contract.v1.UsersService/CurrencyTransfer"
	pathUsersGetAccountValues    = "/tinkoff.public.invest.api.contract.v1.UsersService/GetAccountValues"
	pathUsersGetAccounts         = "/tinkoff.public.invest.api.contract.v1.UsersService/GetAccounts"
	pathUsersGetBankAccounts     = "/tinkoff.public.invest.api.contract.v1.UsersService/GetBankAccounts"
	pathUsersGetInfo             = "/tinkoff.public.invest.api.contract.v1.UsersService/GetInfo"
	pathUsersGetMarginAttributes = "/tinkoff.public.invest.api.contract.v1.UsersService/GetMarginAttributes"
	pathUsersGetUserTariff       = "/tinkoff.public.invest.api.contract.v1.UsersService/GetUserTariff"
	pathUsersPayIn               = "/tinkoff.public.invest.api.contract.v1.UsersService/PayIn"
)

func (s usersServiceClient) CurrencyTransfer(
	ctx context.Context, req *V1CurrencyTransferRequest,
) (*V1CurrencyTransferResponse, error) {
	return do[*V1CurrencyTransferResponse](
		ctx, s.c, pathUsersCurrencyTransfer, req,
	)
}

func (s usersServiceClient) GetAccountValues(
	ctx context.Context, req *V1GetAccountValuesRequest,
) (*V1GetAccountValuesResponse, error) {
	return do[*V1GetAccountValuesResponse](
		ctx, s.c, pathUsersGetAccountValues, req,
	)
}

func (s usersServiceClient) GetAccounts(
	ctx context.Context, req *V1GetAccountsRequest,
) (*V1GetAccountsResponse, error) {
	return do[*V1GetAccountsResponse](ctx, s.c, pathUsersGetAccounts, req)
}

func (s usersServiceClient) GetBankAccounts(
	ctx context.Context, req *V1GetBankAccountsRequest,
) (*V1GetBankAccountsResponse, error) {
	return do[*V1GetBankAccountsResponse](
		ctx, s.c, pathUsersGetBankAccounts, req,
	)
}

func (s usersServiceClient) GetInfo(
	ctx context.Context, req *V1GetInfoRequest,
) (*V1GetInfoResponse, error) {
	return do[*V1GetInfoResponse](ctx, s.c, pathUsersGetInfo, req)
}

func (s usersServiceClient) GetMarginAttributes(
	ctx context.Context, req *V1GetMarginAttributesRequest,
) (*V1GetMarginAttributesResponse, error) {
	return do[*V1GetMarginAttributesResponse](
		ctx, s.c, pathUsersGetMarginAttributes, req,
	)
}

func (s usersServiceClient) GetUserTariff(
	ctx context.Context, req *V1GetUserTariffRequest,
) (*V1GetUserTariffResponse, error) {
	return do[*V1GetUserTariffResponse](ctx, s.c, pathUsersGetUserTariff, req)
}

func (s usersServiceClient) PayIn(
	ctx context.Context, req *V1PayInRequest,
) (*V1PayInResponse, error) {
	return do[*V1PayInResponse](ctx, s.c, pathUsersPayIn, req)
}
