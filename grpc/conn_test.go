package grpc

import (
	"context"
	"testing"

	"github.com/acidsailor/tinvest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConn_MissingToken(t *testing.T) {
	_, err := NewConn(
		context.Background(),
		NewConnConfig(tinvest.EndpointProduction, ""),
	)
	require.Error(t, err)
	assert.ErrorIs(t, err, tinvest.ErrInvalidConfig)
}

func TestNewConn_WithToken(t *testing.T) {
	conn, err := NewConn(
		context.Background(),
		NewConnConfig(tinvest.EndpointProduction, "test-token"),
	)
	require.NoError(t, err)
	require.NotNil(t, conn)
	assert.NoError(t, conn.Close())
}
