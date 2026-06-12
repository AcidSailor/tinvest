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
	rk *restkit.Client

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
type ClientOption func(*Config)

// Config holds construction-time options for a Client. Build it with NewConfig
// and options, or as a struct literal, then pass the result to NewClient.
type Config struct {
	HTTPClient *http.Client
	AppName    string
}

func NewConfig(opts ...ClientOption) Config {
	cfg := Config{
		HTTPClient: &http.Client{Timeout: defaultTimeout},
		AppName:    tinvest.AppName,
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	return cfg
}

// WithHTTPClient sets the *http.Client (custom Timeout/Transport/proxy). A nil
// client makes NewClient return a *ConfigError.
func WithHTTPClient(h *http.Client) ClientOption {
	return func(c *Config) { c.HTTPClient = h }
}

// WithAppName sets the x-app-name header value identifying the application.
func WithAppName(name string) ClientOption {
	return func(c *Config) { c.AppName = name }
}

// WithConfig replaces the whole Config, letting callers build a struct literal
// instead of composing individual options. Because it overwrites every field
// (including defaults), set HTTPClient yourself — a nil HTTPClient makes
// NewClient return a *ConfigError. Options listed after WithConfig still take
// effect.
func WithConfig(cfg Config) ClientOption {
	return func(c *Config) { *c = cfg }
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
// / tinvest.EndpointSandboxREST) with the given API token. Returns a *ConfigError
// on an empty endpoint/token or a nil *http.Client.
func NewClient(endpoint, token string, config *Config) (*Client, error) {
	if token == "" {
		return nil, &ConfigError{Name: clientName, Reason: "empty token"}
	}
	if config.HTTPClient == nil {
		return nil, &ConfigError{Name: clientName, Reason: "nil *http.Client"}
	}
	rk, err := restkit.New(
		endpoint,
		restkit.WithName(clientName),
		restkit.WithHTTPClient(config.HTTPClient),
		restkit.WithHook(authHook(token, config.AppName)),
	)
	if err != nil {
		return nil, err
	}

	cl := &Client{rk: rk}
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
func do[T any](ctx context.Context, c *Client, path string, body any) (T, error) {
	return restkit.Do[T](ctx, c.rk, http.MethodPost, path, body)
}
