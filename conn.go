package tinvest

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// NewConn creates a *grpc.ClientConn configured for the T-Invest API.
// The connection is lazy (grpc.NewClient default) — no actual TCP dial until the first RPC.
// The caller owns the connection lifecycle and is responsible for calling conn.Close().
// ctx is used for structured logging and trace propagation.
func NewConn(ctx context.Context, connConfig *ConnConfig) (*grpc.ClientConn, error) {
	if err := connConfig.Validate(); err != nil {
		return nil, err
	}

	tlsCreds := credentials.NewTLS(&tls.Config{MinVersion: tls.VersionTLS12})
	otelHandler := otelgrpc.NewClientHandler()
	unaryInt := unaryAuthInterceptor(connConfig.token, connConfig.appName)
	streamInt := streamAuthInterceptor(connConfig.token, connConfig.appName)

	conn, err := grpc.NewClient(
		connConfig.endpoint,
		grpc.WithTransportCredentials(tlsCreds),
		grpc.WithChainUnaryInterceptor(unaryInt),
		grpc.WithChainStreamInterceptor(streamInt),
		grpc.WithStatsHandler(otelHandler),
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %w", ErrClient, connConfig.endpoint, err)
	}

	slog.InfoContext(ctx, "tinvest connection created",
		slog.String("endpoint", connConfig.endpoint),
	)

	return conn, nil
}
