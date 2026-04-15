package tinvest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestNewClient_AllServicesInitialized(t *testing.T) {
	conn, err := grpc.NewClient("passthrough:///dummy", grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer func() {
		require.NoError(t, conn.Close())
	}()

	client, err := NewClient(conn, NewClientConfig())
	require.NoError(t, err)
	require.NotNil(t, client)

	assert.NotNil(t, client.Instruments)
	assert.NotNil(t, client.MarketData)
	assert.NotNil(t, client.MarketDataStream)
	assert.NotNil(t, client.Operations)
	assert.NotNil(t, client.OperationsStream)
	assert.NotNil(t, client.Orders)
	assert.NotNil(t, client.OrdersStream)
	assert.NotNil(t, client.StopOrders)
	assert.NotNil(t, client.Sandbox)
	assert.NotNil(t, client.Users)
	assert.NotNil(t, client.Signals)
}

func TestNewClient_NilConn(t *testing.T) {
	client, err := NewClient(nil, NewClientConfig())
	assert.Nil(t, client)
	require.Error(t, err)
	assert.ErrorIs(t, err, ErrClient)
	assert.ErrorIs(t, err, ErrNil)
}

func TestNewClient_NilConfig(t *testing.T) {
	conn, err := grpc.NewClient("passthrough:///dummy", grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer func() {
		require.NoError(t, conn.Close())
	}()

	client, err := NewClient(conn, nil)
	assert.Nil(t, client)
	require.Error(t, err)
	assert.ErrorIs(t, err, ErrClient)
	assert.ErrorIs(t, err, ErrNil)
}
