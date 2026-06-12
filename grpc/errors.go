package grpc

import "fmt"

// connName identifies this package in ConfigError messages.
const connName = "tinvest/grpc"

// ConfigError reports invalid construction input to NewConn (an empty endpoint
// or token). Match it with errors.As — there is no sentinel:
//
//	var ce *grpc.ConfigError
//	if errors.As(err, &ce) { _ = ce.Reason } // e.g. "empty token"
type ConfigError struct {
	Name   string // "tinvest/grpc"
	Reason string // e.g. "empty token"
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("%s: invalid config: %s", e.Name, e.Reason)
}
