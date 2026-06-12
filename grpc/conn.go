package grpc

import (
	"context"
	"crypto/tls"
	"log/slog"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/acidsailor/tinvest"
)

// NewConn creates a *grpc.ClientConn configured for the T-Invest API targeting
// endpoint (use tinvest.EndpointProduction / tinvest.EndpointSandbox) with the
// given API token. The x-app-name header defaults to tinvest.AppName; override
// it with WithAppName.
//
// The connection is lazy (grpc.NewClient default) — no actual TCP dial until the
// first RPC. The caller owns the connection lifecycle and is responsible for
// calling conn.Close(). ctx is used for structured logging and trace propagation.
//
// Returns a *ConfigError on an empty endpoint or token.
func NewConn(
	ctx context.Context,
	endpoint, token string,
	opts ...ConnOption,
) (*grpc.ClientConn, error) {
	if endpoint == "" {
		return nil, &ConfigError{Name: connName, Reason: "empty endpoint"}
	}
	if token == "" {
		return nil, &ConfigError{Name: connName, Reason: "empty token"}
	}

	cfg := connConfig{appName: tinvest.AppName}
	for _, opt := range opts {
		opt(&cfg)
	}

	tlsCreds := credentials.NewTLS(&tls.Config{MinVersion: tls.VersionTLS12})
	otelHandler := otelgrpc.NewClientHandler()
	unaryInt := unaryAuthInterceptor(token, cfg.appName)
	streamInt := streamAuthInterceptor(token, cfg.appName)

	conn, err := grpc.NewClient(
		endpoint,
		grpc.WithTransportCredentials(tlsCreds),
		grpc.WithChainUnaryInterceptor(unaryInt),
		grpc.WithChainStreamInterceptor(streamInt),
		grpc.WithStatsHandler(otelHandler),
	)
	if err != nil {
		return nil, err
	}

	slog.InfoContext(ctx, "tinvest connection created",
		slog.String("endpoint", endpoint),
	)

	return conn, nil
}
