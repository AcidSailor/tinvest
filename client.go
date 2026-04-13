package tinvest

import (
	"fmt"

	pb "github.com/acidsailor/tinvest/pb"
	"google.golang.org/grpc"
)

// Client provides access to all T-Invest API services.
// It wraps an existing *grpc.ClientConn — use NewConn to create one
// with standard T-Invest configuration, or pass your own.
// Does not own the connection — caller is responsible for conn.Close().
type Client struct {
	// Instruments provides reference data for bonds, shares, ETFs, futures, etc.
	Instruments pb.InstrumentsServiceClient
	// MarketData provides current prices, order books, and candles.
	MarketData pb.MarketDataServiceClient
	// MarketDataStream provides real-time market data via server-side streaming.
	MarketDataStream pb.MarketDataStreamServiceClient
	// Operations provides trade and portfolio operation history.
	Operations pb.OperationsServiceClient
	// OperationsStream provides real-time portfolio updates via server-side streaming.
	OperationsStream pb.OperationsStreamServiceClient
	// Orders provides order placement and management.
	Orders pb.OrdersServiceClient
	// OrdersStream provides real-time order state updates via server-side streaming.
	OrdersStream pb.OrdersStreamServiceClient
	// StopOrders provides stop-order placement and management.
	StopOrders pb.StopOrdersServiceClient
	// Sandbox provides a paper-trading environment that mirrors the production API.
	Sandbox pb.SandboxServiceClient
	// Users provides account and tariff information.
	Users pb.UsersServiceClient
	// Signals provides trading signals and strategies.
	Signals pb.SignalServiceClient
}

// NewClient wraps an existing gRPC connection with T-Invest service sub-clients.
// Does not own the connection — caller is responsible for conn.Close().
// Returns an error if conn is nil or if config fails validation.
func NewClient(conn *grpc.ClientConn, config *ClientConfig) (*Client, error) {
	if conn == nil {
		return nil, fmt.Errorf("%w: %w: conn", ErrTInvestClient, ErrNil)
	}
	if err := config.Validate(); err != nil {
		return nil, err
	}
	return &Client{
		Instruments:      pb.NewInstrumentsServiceClient(conn),
		MarketData:       pb.NewMarketDataServiceClient(conn),
		MarketDataStream: pb.NewMarketDataStreamServiceClient(conn),
		Operations:       pb.NewOperationsServiceClient(conn),
		OperationsStream: pb.NewOperationsStreamServiceClient(conn),
		Orders:           pb.NewOrdersServiceClient(conn),
		OrdersStream:     pb.NewOrdersStreamServiceClient(conn),
		StopOrders:       pb.NewStopOrdersServiceClient(conn),
		Sandbox:          pb.NewSandboxServiceClient(conn),
		Users:            pb.NewUsersServiceClient(conn),
		Signals:          pb.NewSignalServiceClient(conn),
	}, nil
}
