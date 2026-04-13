package tinvest

import "fmt"

const (
	// EndpointProduction is the T-Invest live trading API endpoint.
	EndpointProduction = "invest-public-api.tinkoff.ru:443"
	// EndpointSandbox is the T-Invest sandbox API endpoint for testing without real money.
	EndpointSandbox = "sandbox-invest-public-api.tinkoff.ru:443"
	// AppName is the default x-app-name header value identifying this client library.
	AppName = "github.com/acidsailor/tinvest"
)

// ConnConfig holds configuration for the gRPC connection created by NewConn.
// Build with NewConnConfig and use WithAppName to override the default app name.
type ConnConfig struct {
	endpoint string
	token    string
	appName  string
}

// NewConnConfig creates a ConnConfig with the required endpoint and API token.
// Use EndpointProduction or EndpointSandbox as the endpoint value.
func NewConnConfig(endpoint string, token string) *ConnConfig {
	return &ConnConfig{
		token:    token,
		endpoint: endpoint,
		appName:  AppName,
	}
}

// WithAppName sets the x-app-name header sent with every request.
// T-Invest uses this to identify the client application in their logs.
func (c *ConnConfig) WithAppName(name string) *ConnConfig {
	c.appName = name
	return c
}

// Validate checks that the ConnConfig is valid before use.
func (c *ConnConfig) Validate() error {
	if c == nil {
		return fmt.Errorf("%w: %w: conn config", ErrTInvestClient, ErrNil)
	}
	if c.endpoint == "" {
		return fmt.Errorf("%w: %w: empty endpoint", ErrTInvestClient, ErrInvalidConfig)
	}
	if c.token == "" {
		return fmt.Errorf("%w: %w: empty token", ErrTInvestClient, ErrInvalidConfig)
	}
	return nil
}

// ClientConfig holds configuration for the Client created by NewClient.
// Build with NewClientConfig.
type ClientConfig struct {
}

// NewClientConfig creates a ClientConfig with default values.
func NewClientConfig() *ClientConfig {
	return &ClientConfig{}
}

// Validate checks that the ClientConfig is valid before use.
func (c *ClientConfig) Validate() error {
	if c == nil {
		return fmt.Errorf("%w: %w: client config", ErrTInvestClient, ErrNil)
	}
	return nil
}
