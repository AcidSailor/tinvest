// Package tinvest holds configuration and shared primitives for the T-Invest
// (Tinkoff Investments) API clients: endpoint constants, the x-app-name value
// ([AppName]), connection config ([ConnConfig], [NewConnConfig]), and the
// package sentinel errors.
//
// The transport clients live in sub-packages so importing this package stays
// dependency-light: [github.com/acidsailor/tinvest/grpc] provides the gRPC
// client (NewConn / NewClient), and [github.com/acidsailor/tinvest/rest]
// provides the REST gateway client. Financial values are converted via the
// [github.com/acidsailor/tinvest/money] package.
package tinvest
