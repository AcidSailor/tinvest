// Package tinvest provides a Go client for the T-Invest (Tinkoff Investments) gRPC API.
//
// Use [NewConn] to create a configured gRPC connection and [NewClient] to wrap it
// with typed service sub-clients. Financial values are converted between proto
// representations ([pb.Quotation], [pb.MoneyValue]) and [udecimal.Decimal] via helper functions.
package tinvest
