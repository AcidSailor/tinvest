package grpc

import (
	"context"
	"testing"

	"github.com/acidsailor/tinvest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConn_MissingToken(t *testing.T) {
	_, err := NewConn(context.Background(), tinvest.EndpointProduction, "")
	require.Error(t, err)
	var cfgErr *ConfigError
	require.ErrorAs(t, err, &cfgErr)
	assert.Equal(t, "empty token", cfgErr.Reason)
}

func TestNewConn_MissingEndpoint(t *testing.T) {
	_, err := NewConn(context.Background(), "", "test-token")
	require.Error(t, err)
	var cfgErr *ConfigError
	require.ErrorAs(t, err, &cfgErr)
	assert.Equal(t, "empty endpoint", cfgErr.Reason)
}

func TestNewConn_WithToken(t *testing.T) {
	conn, err := NewConn(
		context.Background(),
		tinvest.EndpointProduction,
		"test-token",
	)
	require.NoError(t, err)
	require.NotNil(t, conn)
	assert.NoError(t, conn.Close())
}

func TestNewConn_WithAppName(t *testing.T) {
	conn, err := NewConn(
		context.Background(),
		tinvest.EndpointProduction,
		"test-token",
		WithAppName("myapp"),
	)
	require.NoError(t, err)
	require.NotNil(t, conn)
	assert.NoError(t, conn.Close())
}
