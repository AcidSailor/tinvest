package tinvest

import (
	"fmt"
)

// ConnConfig holds configuration for the gRPC connection created by NewConn.
// Construct it either as a struct literal or via NewConnConfig with options:
//
//	cfg := &ConnConfig{Endpoint: EndpointProduction, Token: tok, AppName: "myapp"}
//	cfg := NewConnConfig(EndpointProduction, tok, WithAppName("myapp"))
type ConnConfig struct {
	Endpoint string
	Token    string
	AppName  string
}

// ConnOption customizes a ConnConfig in NewConnConfig.
type ConnOption func(*ConnConfig)

// WithAppName sets the x-app-name header sent with every request.
// T-Invest uses this to identify the client application in their logs.
func WithAppName(name string) ConnOption {
	return func(c *ConnConfig) { c.AppName = name }
}

// NewConnConfig creates a ConnConfig with the required endpoint and API token,
// defaulting AppName to the package AppName. Use EndpointProduction or
// EndpointSandbox as the endpoint value, and options such as WithAppName to
// override defaults.
func NewConnConfig(endpoint, token string, opts ...ConnOption) *ConnConfig {
	c := &ConnConfig{
		Endpoint: endpoint,
		Token:    token,
		AppName:  AppName,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Validate checks that the ConnConfig is valid before use.
func (c *ConnConfig) Validate() error {
	f := func() error {
		if c == nil {
			return fmt.Errorf("ConnConfig: %w", ErrNil)
		}
		if c.Endpoint == "" {
			return fmt.Errorf("empty endpoint: %w", ErrInvalidConfig)
		}
		if c.Token == "" {
			return fmt.Errorf("empty token: %w", ErrInvalidConfig)
		}
		return nil
	}
	err := f()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrClient, err)
	}
	return nil
}

// ClientConfig holds configuration for the Client created by NewClient.
// Construct it as a struct literal or via NewClientConfig.
type ClientConfig struct{}

// NewClientConfig creates a ClientConfig with default values.
func NewClientConfig() *ClientConfig {
	return &ClientConfig{}
}

// Validate checks that the ClientConfig is valid before use.
func (c *ClientConfig) Validate() error {
	f := func() error {
		if c == nil {
			return fmt.Errorf("ClientConfig: %w", ErrNil)
		}
		return nil
	}
	err := f()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrClient, err)
	}
	return nil
}
