package rest_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/acidsailor/tinvest"
	"github.com/acidsailor/tinvest/rest"
)

func TestClient_GetAccounts_HappyPath(t *testing.T) {
	var gotAuth, gotApp, gotPath string
	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			gotAuth = r.Header.Get("Authorization")
			gotApp = r.Header.Get("x-app-name")
			gotPath = r.URL.Path
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"accounts": []map[string]any{
					{"id": "acc-1", "name": "Brokerage"},
				},
			})
		}))
	defer srv.Close()

	c, err := rest.NewClient(srv.URL, "tkn-123")
	require.NoError(t, err)

	resp, err := c.Users.GetAccounts(
		context.Background(),
		&rest.V1GetAccountsRequest{},
	)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, "Bearer tkn-123", gotAuth)
	assert.NotEmpty(t, gotApp)
	assert.Equal(
		t,
		"/tinkoff.public.invest.api.contract.v1.UsersService/GetAccounts",
		gotPath,
	)
}

func TestClient_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"code":16,"message":"auth"}`))
		}))
	defer srv.Close()

	c, err := rest.NewClient(srv.URL, "bad")
	require.NoError(t, err)

	_, err = c.Users.GetAccounts(
		context.Background(),
		&rest.V1GetAccountsRequest{},
	)
	require.Error(t, err)
	var apiErr *rest.ResponseError
	require.ErrorAs(t, err, &apiErr)
	assert.Equal(t, http.StatusUnauthorized, apiErr.StatusCode)
}

func TestClient_RequestError_Unmarshal(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{ not valid json `))
		}))
	defer srv.Close()

	c, err := rest.NewClient(srv.URL, "tkn")
	require.NoError(t, err)

	_, err = c.Users.GetAccounts(
		context.Background(),
		&rest.V1GetAccountsRequest{},
	)
	require.Error(t, err)
	var reqErr *rest.RequestError
	require.ErrorAs(t, err, &reqErr)
	assert.Equal(t, rest.OpUnmarshal, reqErr.Op)
}

func TestClient_RequestError_Send(t *testing.T) {
	// A server that is shut down before the call yields a connection refused,
	// surfacing as a RequestError at the send stage.
	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {}))
	url := srv.URL
	srv.Close()

	c, err := rest.NewClient(url, "tkn")
	require.NoError(t, err)

	_, err = c.Users.GetAccounts(
		context.Background(),
		&rest.V1GetAccountsRequest{},
	)
	require.Error(t, err)
	var reqErr *rest.RequestError
	require.ErrorAs(t, err, &reqErr)
	assert.Equal(t, rest.OpSend, reqErr.Op)
}

func TestNewClient_Validation(t *testing.T) {
	var cfgErr *rest.ConfigError
	_, err := rest.NewClient("", "tkn")
	require.ErrorAs(t, err, &cfgErr)
	_, err = rest.NewClient(tinvest.EndpointProductionREST, "")
	require.ErrorAs(t, err, &cfgErr)
	_, err = rest.NewClient(tinvest.EndpointProductionREST, "tkn",
		rest.WithHTTPClient(nil))
	require.ErrorAs(t, err, &cfgErr)
}

func TestNewClient_WithOptions(t *testing.T) {
	var gotApp string
	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			gotApp = r.Header.Get("x-app-name")
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{}`))
		}))
	defer srv.Close()

	c, err := rest.NewClient(srv.URL, "tkn",
		rest.WithHTTPClient(&http.Client{}),
		rest.WithAppName("literal-app"),
	)
	require.NoError(t, err)

	_, err = c.Users.GetAccounts(
		context.Background(),
		&rest.V1GetAccountsRequest{},
	)
	require.NoError(t, err)
	assert.Equal(t, "literal-app", gotApp)
}

func TestNewClient_NilHTTPClientErrors(t *testing.T) {
	var cfgErr *rest.ConfigError
	_, err := rest.NewClient(tinvest.EndpointProductionREST, "tkn",
		rest.WithHTTPClient(nil))
	require.ErrorAs(t, err, &cfgErr)
}
