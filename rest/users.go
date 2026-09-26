package rest

import "context"

type usersServiceClient struct{ c *Client }

const (
	pathUsersCurrencyTransfer endpoint[
		*V1CurrencyTransferRequest,
		*V1CurrencyTransferResponse,
	] = "/tinkoff.public.invest.api.contract.v1.UsersService/CurrencyTransfer"

	pathUsersGetAccountValues endpoint[
		*V1GetAccountValuesRequest,
		*V1GetAccountValuesResponse,
	] = "/tinkoff.public.invest.api.contract.v1.UsersService/GetAccountValues"

	pathUsersGetAccounts endpoint[
		*V1GetAccountsRequest,
		*V1GetAccountsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.UsersService/GetAccounts"

	pathUsersGetBankAccounts endpoint[
		*V1GetBankAccountsRequest,
		*V1GetBankAccountsResponse,
	] = "/tinkoff.public.invest.api.contract.v1.UsersService/GetBankAccounts"

	pathUsersGetInfo endpoint[
		*V1GetInfoRequest,
		*V1GetInfoResponse,
	] = "/tinkoff.public.invest.api.contract.v1.UsersService/GetInfo"

	pathUsersGetMarginAttributes endpoint[
		*V1GetMarginAttributesRequest,
		*V1GetMarginAttributesResponse,
	] = "/tinkoff.public.invest.api.contract.v1.UsersService/GetMarginAttributes"

	pathUsersGetUserTariff endpoint[
		*V1GetUserTariffRequest,
		*V1GetUserTariffResponse,
	] = "/tinkoff.public.invest.api.contract.v1.UsersService/GetUserTariff"

	pathUsersPayIn endpoint[
		*V1PayInRequest,
		*V1PayInResponse,
	] = "/tinkoff.public.invest.api.contract.v1.UsersService/PayIn"
)

func (s *usersServiceClient) CurrencyTransfer(
	ctx context.Context, req *V1CurrencyTransferRequest,
) (*V1CurrencyTransferResponse, error) {
	return call(ctx, s.c, pathUsersCurrencyTransfer, req)
}

func (s *usersServiceClient) GetAccountValues(
	ctx context.Context, req *V1GetAccountValuesRequest,
) (*V1GetAccountValuesResponse, error) {
	return call(ctx, s.c, pathUsersGetAccountValues, req)
}

func (s *usersServiceClient) GetAccounts(
	ctx context.Context, req *V1GetAccountsRequest,
) (*V1GetAccountsResponse, error) {
	return call(ctx, s.c, pathUsersGetAccounts, req)
}

func (s *usersServiceClient) GetBankAccounts(
	ctx context.Context, req *V1GetBankAccountsRequest,
) (*V1GetBankAccountsResponse, error) {
	return call(ctx, s.c, pathUsersGetBankAccounts, req)
}

func (s *usersServiceClient) GetInfo(
	ctx context.Context, req *V1GetInfoRequest,
) (*V1GetInfoResponse, error) {
	return call(ctx, s.c, pathUsersGetInfo, req)
}

func (s *usersServiceClient) GetMarginAttributes(
	ctx context.Context, req *V1GetMarginAttributesRequest,
) (*V1GetMarginAttributesResponse, error) {
	return call(ctx, s.c, pathUsersGetMarginAttributes, req)
}

func (s *usersServiceClient) GetUserTariff(
	ctx context.Context, req *V1GetUserTariffRequest,
) (*V1GetUserTariffResponse, error) {
	return call(ctx, s.c, pathUsersGetUserTariff, req)
}

func (s *usersServiceClient) PayIn(
	ctx context.Context, req *V1PayInRequest,
) (*V1PayInResponse, error) {
	return call(ctx, s.c, pathUsersPayIn, req)
}
