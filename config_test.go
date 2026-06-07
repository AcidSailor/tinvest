package tinvest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConnConfig_StoresFields(t *testing.T) {
	cfg := NewConnConfig(EndpointProduction, "test-token")

	assert.Equal(t, "test-token", cfg.Token)
	assert.Equal(t, EndpointProduction, cfg.Endpoint)
	assert.Equal(t, AppName, cfg.AppName)
}

func TestNewConnConfig_StructLiteral(t *testing.T) {
	cfg := &ConnConfig{
		Endpoint: EndpointProduction,
		Token:    "tok",
		AppName:  "myapp",
	}

	assert.NoError(t, cfg.Validate())
	assert.Equal(t, "myapp", cfg.AppName)
}

func TestWithAppName(t *testing.T) {
	cfg := NewConnConfig(EndpointProduction, "tok", WithAppName("myapp"))

	assert.Equal(t, "myapp", cfg.AppName)
}

func TestWithAppName_EmptyClearsAppName(t *testing.T) {
	cfg := NewConnConfig(EndpointProduction, "tok", WithAppName(""))

	assert.Equal(t, "", cfg.AppName)
}

func TestConnConfig_Validate_OK(t *testing.T) {
	cfg := NewConnConfig(EndpointProduction, "tok")

	assert.NoError(t, cfg.Validate())
}

func TestConnConfig_Validate_NilConfig(t *testing.T) {
	var cfg *ConnConfig

	err := cfg.Validate()
	assert.ErrorIs(t, err, ErrClient)
	assert.ErrorIs(t, err, ErrNil)
}

func TestConnConfig_Validate_EmptyToken(t *testing.T) {
	cfg := NewConnConfig(EndpointProduction, "")

	err := cfg.Validate()
	assert.ErrorIs(t, err, ErrClient)
	assert.ErrorIs(t, err, ErrInvalidConfig)
}

func TestConnConfig_Validate_EmptyEndpoint(t *testing.T) {
	cfg := NewConnConfig("", "tok")

	err := cfg.Validate()
	assert.ErrorIs(t, err, ErrClient)
	assert.ErrorIs(t, err, ErrInvalidConfig)
}

func TestClientConfig_Validate_OK(t *testing.T) {
	cfg := NewClientConfig()

	assert.NoError(t, cfg.Validate())
}

func TestClientConfig_Validate_NilConfig(t *testing.T) {
	var cfg *ClientConfig

	err := cfg.Validate()
	assert.ErrorIs(t, err, ErrClient)
	assert.ErrorIs(t, err, ErrNil)
}
