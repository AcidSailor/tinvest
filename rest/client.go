package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/acidsailor/restkit"
	"github.com/acidsailor/tinvest"
)

const (
	defaultTimeout = 30 * time.Second
	clientName     = "tinvest/rest"
)

// Client is the T-Invest REST gateway client. Immutable after construction and
// safe for concurrent use. Service methods hang off the per-service fields.
type Client struct {
	rkClient *restkit.Client

	Instruments instrumentsServiceClient
	MarketData  marketDataServiceClient
	Operations  operationsServiceClient
	Orders      ordersServiceClient
	Sandbox     sandboxServiceClient
	Signals     signalsServiceClient
	StopOrders  stopOrdersServiceClient
	Users       usersServiceClient
}

// ClientOption configures a Client at construction.
type ClientOption func(*config)

// config holds the resolved construction-time options for a Client.
type config struct {
	httpClient *http.Client
	appName    string
}

// WithHTTPClient sets the *http.Client (custom Timeout/Transport/proxy). Passing
// a nil client makes NewClient return a *ConfigError.
func WithHTTPClient(h *http.Client) ClientOption {
	return func(c *config) { c.httpClient = h }
}

// WithAppName sets the x-app-name header value identifying the application.
func WithAppName(name string) ClientOption {
	return func(c *config) { c.appName = name }
}

// authHook sets the per-client auth and app-name headers on every request.
func authHook(token, appName string) restkit.RequestHook {
	return func(r *http.Request) error {
		r.Header.Set("Authorization", "Bearer "+token)
		r.Header.Set("x-app-name", appName)
		return nil
	}
}

// NewClient builds a Client targeting endpoint (use tinvest.EndpointProductionREST
// / tinvest.EndpointSandboxREST) with the given API token. The HTTP client
// defaults to a 30s-timeout *http.Client and the x-app-name header to
// tinvest.AppName; override either with WithHTTPClient / WithAppName. Returns a
// *ConfigError on an empty token or a nil *http.Client.
func NewClient(endpoint, token string, opts ...ClientOption) (*Client, error) {
	if token == "" {
		return nil, &ConfigError{Name: clientName, Reason: "empty token"}
	}

	cfg := config{
		httpClient: &http.Client{Timeout: defaultTimeout},
		appName:    tinvest.AppName,
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	if cfg.httpClient == nil {
		return nil, &ConfigError{Name: clientName, Reason: "nil *http.Client"}
	}

	rkClient, err := restkit.New(
		endpoint,
		restkit.WithName(clientName),
		restkit.WithHTTPClient(cfg.httpClient),
		restkit.WithHook(authHook(token, cfg.appName)),
	)
	if err != nil {
		return nil, err
	}

	cl := &Client{rkClient: rkClient}
	cl.Instruments = instrumentsServiceClient{cl}
	cl.MarketData = marketDataServiceClient{cl}
	cl.Operations = operationsServiceClient{cl}
	cl.Orders = ordersServiceClient{cl}
	cl.Sandbox = sandboxServiceClient{cl}
	cl.Signals = signalsServiceClient{cl}
	cl.StopOrders = stopOrdersServiceClient{cl}
	cl.Users = usersServiceClient{cl}
	return cl, nil
}

// do POSTs body as JSON to path and decodes the 2xx response into T. A non-2xx
// becomes a *ResponseError; any other stage failure a *RequestError (both via
// errors.As).
func do[T any](
	ctx context.Context,
	c *Client,
	path string,
	body any,
) (T, error) {
	return restkit.Do[T](ctx, c.rkClient, http.MethodPost, path, body)
}
