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

	Instruments instrumentsService
	MarketData  marketDataService
	Operations  operationsService
	Orders      ordersService
	Sandbox     sandboxService
	Signals     signalService
	StopOrders  stopOrdersService
	Users       usersService
}

// ClientOption configures a Client at construction.
type ClientOption func(*config)

type config struct {
	httpClient *http.Client
	appName    string
}

// WithHTTPClient sets the *http.Client (custom Timeout/Transport/proxy). A nil
// client makes NewClient return an error.
func WithHTTPClient(h *http.Client) ClientOption {
	return func(c *config) { c.httpClient = h }
}

// WithAppName sets the x-app-name header value identifying the application.
func WithAppName(name string) ClientOption {
	return func(c *config) { c.appName = name }
}

// NewClient builds a Client targeting endpoint (use EndpointProduction /
// EndpointSandbox) with the given API token. Returns an error wrapping
// ErrClient on empty endpoint/token or a nil *http.Client.
func NewClient(endpoint, token string, opts ...ClientOption) (*Client, error) {
	c, err := func() (*Client, error) {
		if endpoint == "" {
			return nil, errors.New("empty endpoint")
		}
		if token == "" {
			return nil, errors.New("empty token")
		}
		cfg := config{
			httpClient: &http.Client{Timeout: defaultTimeout},
			appName:    defaultAppName,
		}
		for _, opt := range opts {
			opt(&cfg)
		}
		if cfg.httpClient == nil {
			return nil, errors.New("nil *http.Client")
		}
		hc := *cfg.httpClient
		base := hc.Transport
		if base == nil {
			base = http.DefaultTransport
		}
		hc.Transport = otelhttp.NewTransport(base)

		cl := &Client{
			baseURL:    strings.TrimRight(endpoint, "/"),
			token:      token,
			appName:    cfg.appName,
			httpClient: &hc,
		}
		cl.Instruments = instrumentsService{cl}
		cl.MarketData = marketDataService{cl}
		cl.Operations = operationsService{cl}
		cl.Orders = ordersService{cl}
		cl.Sandbox = sandboxService{cl}
		cl.Signals = signalService{cl}
		cl.StopOrders = stopOrdersService{cl}
		cl.Users = usersService{cl}
		return cl, nil
	}()
	if err != nil {
		return nil, errors.Join(ErrClient, err)
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
		return nil, errors.Join(ErrClient, err)
	}
	return out, nil
}
