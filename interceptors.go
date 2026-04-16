package tinvest

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// injectMetadata adds the bearer token and x-app-name headers to the outgoing
// gRPC context. The x-app-name header is omitted when appName is empty.
// Existing authorization and x-app-name values are overwritten.
func injectMetadata(
	ctx context.Context,
	token, appName string,
) context.Context {
	pairs := []string{"authorization", "Bearer " + token}
	if appName != "" {
		pairs = append(pairs, "x-app-name", appName)
	}

	md, ok := metadata.FromOutgoingContext(ctx)
	if ok {
		md = md.Copy()
	} else {
		md = metadata.MD{}
	}
	for i := 0; i < len(pairs); i += 2 {
		md.Set(pairs[i], pairs[i+1])
	}

	return metadata.NewOutgoingContext(ctx, md)
}

// unaryAuthInterceptor returns a gRPC unary client interceptor that injects
// the bearer token and x-app-name into request metadata.
func unaryAuthInterceptor(token, appName string) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		return invoker(
			injectMetadata(ctx, token, appName),
			method,
			req,
			reply,
			cc,
			opts...)
	}
}

// streamAuthInterceptor returns a gRPC stream client interceptor that injects
// the bearer token and x-app-name into request metadata.
func streamAuthInterceptor(token, appName string) grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		return streamer(
			injectMetadata(ctx, token, appName),
			desc,
			cc,
			method,
			opts...)
	}
}
