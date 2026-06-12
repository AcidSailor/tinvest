package grpc

// connConfig holds the resolved gRPC connection settings applied by NewConn.
type connConfig struct {
	appName string
}

// ConnOption customizes the gRPC connection created by NewConn.
type ConnOption func(*connConfig)

// WithAppName sets the x-app-name header sent with every request.
// T-Invest uses this to identify the client application in their logs.
func WithAppName(name string) ConnOption {
	return func(c *connConfig) { c.appName = name }
}
