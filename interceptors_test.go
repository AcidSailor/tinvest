package tinvest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func extractOutgoingMetadata(ctx context.Context) metadata.MD {
	md, _ := metadata.FromOutgoingContext(ctx)
	return md
}

func captureUnaryCtx(
	t *testing.T,
	interceptor grpc.UnaryClientInterceptor,
) context.Context {
	t.Helper()
	var captured context.Context
	invoker := func(ctx context.Context, _ string, _, _ any, _ *grpc.ClientConn, _ ...grpc.CallOption) error {
		captured = ctx
		return nil
	}
	err := interceptor(
		context.Background(),
		"/test/Method",
		nil,
		nil,
		nil,
		invoker,
	)
	require.NoError(t, err)
	return captured
}

func captureStreamCtx(
	t *testing.T,
	interceptor grpc.StreamClientInterceptor,
) context.Context {
	t.Helper()
	var captured context.Context
	streamer := func(
		ctx context.Context,
		_ *grpc.StreamDesc,
		_ *grpc.ClientConn,
		_ string,
		_ ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		captured = ctx
		return nil, nil
	}
	_, err := interceptor(
		context.Background(),
		nil,
		nil,
		"/test/Stream",
		streamer,
	)
	require.NoError(t, err)
	return captured
}

func TestUnaryInterceptor_InjectsToken(t *testing.T) {
	ctx := captureUnaryCtx(t, unaryAuthInterceptor("my-token", AppName))
	md := extractOutgoingMetadata(ctx)

	assert.Equal(t, []string{"Bearer my-token"}, md.Get("authorization"))
}

func TestUnaryInterceptor_InjectsAppName(t *testing.T) {
	ctx := captureUnaryCtx(t, unaryAuthInterceptor("tok", "testapp"))
	md := extractOutgoingMetadata(ctx)

	assert.Equal(t, []string{"testapp"}, md.Get("x-app-name"))
}

func TestUnaryInterceptor_OmitsAppNameWhenEmpty(t *testing.T) {
	ctx := captureUnaryCtx(t, unaryAuthInterceptor("tok", ""))
	md := extractOutgoingMetadata(ctx)

	assert.Empty(t, md.Get("x-app-name"))
}

func TestStreamInterceptor_OmitsAppNameWhenEmpty(t *testing.T) {
	ctx := captureStreamCtx(t, streamAuthInterceptor("tok", ""))
	md := extractOutgoingMetadata(ctx)

	assert.Empty(t, md.Get("x-app-name"))
}

func TestStreamInterceptor_InjectsToken(t *testing.T) {
	ctx := captureStreamCtx(t, streamAuthInterceptor("stream-token", AppName))
	md := extractOutgoingMetadata(ctx)

	assert.Equal(t, []string{"Bearer stream-token"}, md.Get("authorization"))
}

func TestStreamInterceptor_InjectsAppName(t *testing.T) {
	ctx := captureStreamCtx(t, streamAuthInterceptor("tok", "streamapp"))
	md := extractOutgoingMetadata(ctx)

	assert.Equal(t, []string{"streamapp"}, md.Get("x-app-name"))
}

func TestUnaryInterceptor_PropagatesInvokerError(t *testing.T) {
	want := assert.AnError
	invoker := func(ctx context.Context, _ string, _, _ any, _ *grpc.ClientConn, _ ...grpc.CallOption) error {
		return want
	}
	err := unaryAuthInterceptor(
		"tok",
		AppName,
	)(
		context.Background(),
		"/test/Method",
		nil,
		nil,
		nil,
		invoker,
	)

	assert.ErrorIs(t, err, want)
}

func TestStreamInterceptor_PropagatesStreamerError(t *testing.T) {
	want := assert.AnError
	streamer := func(ctx context.Context, _ *grpc.StreamDesc, _ *grpc.ClientConn, _ string, _ ...grpc.CallOption) (grpc.ClientStream, error) {
		return nil, want
	}
	_, err := streamAuthInterceptor(
		"tok",
		AppName,
	)(
		context.Background(),
		nil,
		nil,
		"/test/Stream",
		streamer,
	)

	assert.ErrorIs(t, err, want)
}

func TestInjectMetadata_OverwritesExistingAuth(t *testing.T) {
	existing := metadata.Pairs(
		"authorization",
		"Bearer old-token",
		"x-app-name",
		"old-app",
	)
	ctx := metadata.NewOutgoingContext(context.Background(), existing)

	ctx = injectMetadata(ctx, "new-token", "new-app")
	md := extractOutgoingMetadata(ctx)

	assert.Equal(t, []string{"Bearer new-token"}, md.Get("authorization"))
	assert.Equal(t, []string{"new-app"}, md.Get("x-app-name"))
}

func TestInjectMetadata_PreservesExistingMetadata(t *testing.T) {
	existing := metadata.Pairs("x-custom-header", "custom-value")
	ctx := metadata.NewOutgoingContext(context.Background(), existing)

	ctx = injectMetadata(ctx, "my-token", "myapp")
	md := extractOutgoingMetadata(ctx)

	assert.Equal(t, []string{"custom-value"}, md.Get("x-custom-header"))
	assert.Equal(t, []string{"Bearer my-token"}, md.Get("authorization"))
	assert.Equal(t, []string{"myapp"}, md.Get("x-app-name"))
}
