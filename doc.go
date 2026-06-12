// Package tinvest holds shared primitives for the T-Invest (Tinkoff
// Investments) API clients: endpoint constants and the x-app-name value
// ([AppName]).
//
// The transport clients live in sub-packages so importing this package stays
// dependency-light: [github.com/acidsailor/tinvest/grpc] provides the gRPC
// client (NewConn / NewClient), and [github.com/acidsailor/tinvest/rest]
// provides the REST gateway client. Each transport owns its own typed
// *ConfigError for invalid construction input — there are no sentinel errors.
// Financial values are converted via the [github.com/acidsailor/tinvest/money]
// package.
package tinvest
