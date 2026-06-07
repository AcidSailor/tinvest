package grpc

import (
	"testing"

	"github.com/acidsailor/tinvest"
	"github.com/stretchr/testify/assert"
)

func TestNewConnConfig_StoresFields(t *testing.T) {
	cfg := NewConnConfig(tinvest.EndpointProduction, "test-token")

	assert.Equal(t, "test-token", cfg.Token)
	assert.Equal(t, tinvest.EndpointProduction, cfg.Endpoint)
	assert.Equal(t, tinvest.AppName, cfg.AppName)
}

func TestNewConnConfig_StructLiteral(t *testing.T) {
	cfg := &ConnConfig{
		Endpoint: tinvest.EndpointProduction,
		Token:    "tok",
		AppName:  "myapp",
	}

	assert.NoError(t, cfg.Validate())
	assert.Equal(t, "myapp", cfg.AppName)
}

func TestWithAppName(t *testing.T) {
	cfg := NewConnConfig(
		tinvest.EndpointProduction,
		"tok",
		WithAppName("myapp"),
	)

	assert.Equal(t, "myapp", cfg.AppName)
}

func TestWithAppName_EmptyClearsAppName(t *testing.T) {
	cfg := NewConnConfig(tinvest.EndpointProduction, "tok", WithAppName(""))

	assert.Equal(t, "", cfg.AppName)
}

func TestConnConfig_Validate_OK(t *testing.T) {
	cfg := NewConnConfig(tinvest.EndpointProduction, "tok")

	assert.NoError(t, cfg.Validate())
}

func TestConnConfig_Validate_NilConfig(t *testing.T) {
	var cfg *ConnConfig

	err := cfg.Validate()
	assert.ErrorIs(t, err, tinvest.ErrClient)
	assert.ErrorIs(t, err, tinvest.ErrNil)
}

func TestConnConfig_Validate_EmptyToken(t *testing.T) {
	cfg := NewConnConfig(tinvest.EndpointProduction, "")

	err := cfg.Validate()
	assert.ErrorIs(t, err, tinvest.ErrClient)
	assert.ErrorIs(t, err, tinvest.ErrInvalidConfig)
}

func TestConnConfig_Validate_EmptyEndpoint(t *testing.T) {
	cfg := NewConnConfig("", "tok")

	err := cfg.Validate()
	assert.ErrorIs(t, err, tinvest.ErrClient)
	assert.ErrorIs(t, err, tinvest.ErrInvalidConfig)
}

func TestClientConfig_Validate_OK(t *testing.T) {
	cfg := NewClientConfig()

	assert.NoError(t, cfg.Validate())
}

func TestClientConfig_Validate_NilConfig(t *testing.T) {
	var cfg *ClientConfig

	err := cfg.Validate()
	assert.ErrorIs(t, err, tinvest.ErrClient)
	assert.ErrorIs(t, err, tinvest.ErrNil)
}
