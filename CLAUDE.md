# tinvest

Go client library for the T-Invest (Tinkoff Investments) gRPC API.

## Structure

- Root package `tinvest` — connection, client, config, helpers, interceptors
- `pb/` — proto-generated gRPC bindings (do not edit by hand; regenerate via `task proto`)
- `proto/` — buf generation config

## Tasks

- `task proto` — regenerate `pb/` from upstream T-Invest proto contracts
- `task test` — run tests
- `task lint` — run gofmt + golangci-lint

## Proto generation

Proto sources are cloned at build time from `https://opensource.tbank.ru/invest/invest-contracts.git` into `tmp/invest-contracts/`.

## Notes
- MCP server lives in the separate `tinvestmcp` repo
