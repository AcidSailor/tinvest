# tinvest

Go client library for the T-Invest (Tinkoff Investments) gRPC API.

## Structure

- Root package `tinvest` — connection, client, config, helpers, interceptors
- `pb/` — proto-generated gRPC bindings (do not edit by hand; regenerate via `task proto`)
- `proto/` — buf generation config

## Tasks

- `task proto` — regenerate `pb/` from upstream T-Invest proto contracts
- `task lint` — run formatters (gofumpt + golines) and linters with autofix
- `task ci` — read-only fmt + lint verification (fail-fast; used by GitHub Actions)
- `task test` — run tests
- `task check` — local composite: `task lint` + `task test`

## Proto generation

Proto sources are cloned at build time from `https://opensource.tbank.ru/invest/invest-contracts.git` into `tmp/invest-contracts/`.
