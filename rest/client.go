package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/acidsailor/tinvest"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

const defaultTimeout = 30 * time.Second

// Client is the T-Invest REST gateway client. Immutable after construction and
// safe for concurrent use. Service methods hang off the per-service fields.
type Client struct {
	baseURL    string
	token      string
	appName    string
	httpClient *http.Client

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

// Config holds construction-time options for a Client. Apply it with options via
// NewClient, or build a Config struct literal directly and pass it with
// WithConfig.
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
// client makes NewClient return an error.
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
// NewClient return an error. Options listed after WithConfig still take effect.
func WithConfig(cfg Config) ClientOption {
	return func(c *Config) { *c = cfg }
}

// NewClient builds a Client targeting endpoint (use EndpointProduction /
// EndpointSandbox) with the given API token. Returns an error wrapping
// ErrClient on empty endpoint/token or a nil *http.Client.
func NewClient(endpoint, token string, config *Config) (*Client, error) {
	c, err := func() (*Client, error) {
		if endpoint == "" {
			return nil, errors.New("empty endpoint")
		}
		if token == "" {
			return nil, errors.New("empty token")
		}
		if config.HTTPClient == nil {
			return nil, errors.New("nil *http.Client")
		}
		hc := *config.HTTPClient
		base := hc.Transport
		if base == nil {
			base = http.DefaultTransport
		}
		hc.Transport = otelhttp.NewTransport(base)

		cl := &Client{
			baseURL:    strings.TrimRight(endpoint, "/"),
			token:      token,
			appName:    config.AppName,
			httpClient: &hc,
		}
		cl.Instruments = instrumentsServiceClient{cl}
		cl.MarketData = marketDataServiceClient{cl}
		cl.Operations = operationsServiceClient{cl}
		cl.Orders = ordersServiceClient{cl}
		cl.Sandbox = sandboxServiceClient{cl}
		cl.Signals = signalsServiceClient{cl}
		cl.StopOrders = stopOrdersServiceClient{cl}
		cl.Users = usersServiceClient{cl}
		return cl, nil
	}()
	if err != nil {
		return nil, errors.Join(tinvest.ErrClient, err)
	}
	return c, nil
}

// do POSTs body as JSON to baseURL+path and decodes the response into *Resp. A
// non-2xx becomes an *APIError; all failures join ErrClient.
func do[Resp any](
	ctx context.Context, c *Client, path string, body any,
) (*Resp, error) {
	out, err := func() (result *Resp, err error) {
		raw, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("encode request body: %w", err)
		}
		req, err := http.NewRequestWithContext(
			ctx, http.MethodPost, c.baseURL+path, bytes.NewReader(raw),
		)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+c.token)
		req.Header.Set("x-app-name", c.appName)

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer func() { err = errors.Join(err, resp.Body.Close()) }()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("read response body: %w", err)
		}
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return nil, &APIError{
				StatusCode: resp.StatusCode,
				Body:       string(data),
			}
		}
		var decoded Resp
		if len(bytes.TrimSpace(data)) > 0 {
			if err := json.Unmarshal(data, &decoded); err != nil {
				return nil, fmt.Errorf("decode response: %w", err)
			}
		}
		return &decoded, nil
	}()
	if err != nil {
		return nil, errors.Join(tinvest.ErrClient, err)
	}
	return out, nil
}
