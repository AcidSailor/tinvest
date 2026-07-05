package rest

import (
	"context"
	"crypto/tls"
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
	httpClient    *http.Client
	httpClientSet bool
	appName       string
}

// WithHTTPClient sets the *http.Client (custom Timeout/Transport/proxy). It is
// used verbatim, so the caller owns TLS: to reach the T-Invest API the client's
// transport must trust the Russian Trusted Root CA (see tinvest.RootCAs).
// Passing a nil client makes NewClient return a *ConfigError.
func WithHTTPClient(h *http.Client) ClientOption {
	return func(c *config) {
		c.httpClient = h
		c.httpClientSet = true
	}
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

	cfg := config{appName: tinvest.AppName}
	for _, opt := range opts {
		opt(&cfg)
	}
	if cfg.httpClientSet {
		if cfg.httpClient == nil {
			return nil, &ConfigError{
				Name:   clientName,
				Reason: "nil *http.Client",
			}
		}
	} else {
		hc, err := defaultHTTPClient()
		if err != nil {
			return nil, &ConfigError{Name: clientName, Reason: err.Error()}
		}
		cfg.httpClient = hc
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

// defaultHTTPClient builds the *http.Client used when the caller does not supply
// one. It clones http.DefaultTransport (preserving proxy, pooling and HTTP/2
// settings) and pins a TLS config that trusts the Russian Trusted Root CA the
// T-Invest API chains to (see tinvest.RootCAs).
func defaultHTTPClient() (*http.Client, error) {
	pool, err := tinvest.RootCAs()
	if err != nil {
		return nil, err
	}
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.TLSClientConfig = &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    pool,
	}
	return &http.Client{Timeout: defaultTimeout, Transport: tr}, nil
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
