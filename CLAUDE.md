# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

A Go client library for the T-Invest (Tinkoff Investments) API. It exposes two
independent transports over the same upstream contract — a gRPC client and a
REST gateway client — plus a transport-agnostic money package. No `main`; this
is a library consumed by other modules.

## Commands

Tasks are driven by [Task](https://taskfile.dev) (`taskfile.yml`):

- `task test` — `go test -race ./...`
- `task lint` — golangci-lint **fmt + run --fix** (mutates files)
- `task ci` — read-only fmt `--diff` + lint (what CI enforces; fail-fast)
- `task check` — `lint` then `test`
- `task build` — GoReleaser snapshot (validates release config; no publish)

Run a single test directly:

```sh
go test -race ./grpc/ -run TestGRPCClientMatchesSpec
```

CI (`.github/workflows/ci.yml`) delegates to the shared
`acidsailor/go-scaffolds` `go-ci.yml` reusable workflow — it runs the equivalent
of `task ci` plus the race test suite.

## Code generation

Both transports are generated from one pinned upstream release. `CONTRACTS_VERSION`
in `taskfile.yml` is the **single source of truth** and pins both the proto
contracts (git tag) and the REST OpenAPI spec (raw tag URL) so the transports
never drift onto different upstream versions. Generated code is committed.

- `task proto` — clones `invest-contracts`, runs `buf generate` → `grpc/pb/**`
- `task spec` — downloads the upstream OpenAPI YAML → `spec/spec-upstream.yaml` (committed verbatim)
- `task rest` — `deref` + `overlay` + `gen`: produces `rest/models.gen.go` (oapi-codegen, plain structs) and `spec/spec-deref.json`

Regenerate after bumping `CONTRACTS_VERSION`. Do not hand-edit `grpc/pb/**`,
`rest/models.gen.go`, or the `spec/spec-*.{yaml,json}` artifacts.

## Architecture

Root package `tinvest` (`endpoints.go`, `doc.go`) holds only shared primitives —
endpoint constants and the `AppName` (`x-app-name`) value — so importing it stays
dependency-light. The transports live in sub-packages:

- **`grpc/`** — `NewConn(ctx, endpoint, token, ...)` builds a `*grpc.ClientConn`
  (lazy dial, TLS 1.2+, otel stats handler, auth interceptors). `NewClient(conn)`
  wraps it with one exported field per service (`Instruments`, `MarketData`,
  `Orders`, streaming variants, …). **The client does not own the connection —
  the caller calls `conn.Close()`.** Auth (`authorization: Bearer` + `x-app-name`)
  is injected via unary/stream client interceptors (`interceptors.go`).

- **`rest/`** — `NewClient(endpoint, token, ...)` over `acidsailor/restkit`.
  Every RPC is a `POST` of a JSON body to the gateway path
  `/<protoContractPkg><Service>/<Method>` (see the `path*` constants per service
  file). Service methods hang off exported per-service fields. Auth headers are
  set by a restkit request hook. `do[T]` is the shared generic POST+decode helper.

- **`money/`** — protobuf-free conversions between T-Invest `units`+`nano`
  (int64 + int32 nano-billionths) and `udecimal.Decimal`, plus sign handling and
  display formatting. It is the single home of this math; `grpc/helpers.go` and
  `grpc/format.go` adapt the proto `MoneyValue`/`Quotation` types onto it, and
  JSON-typed callers (e.g. an external MCP server) share the same implementation.

- **`spec/`** — embeds the dereferenced OpenAPI doc (`spec-deref.json`) via
  `//go:embed` for downstream JSON-schema assembly **and** for the parity tests.

### Transport parity (key invariant)

The gateway and gRPC method paths share one namespace
(`tinkoff.public.invest.api.contract.v1.<Service>/<Method>`), which is the same
namespace the REST spec paths use. This lets both transports be checked against
one source of truth:

- `grpc/parity_test.go::TestGRPCClientMatchesSpec` — reflects over `Client`
  fields and asserts it exposes **exactly** the unary operations the spec defines.
- `rest/paths_test.go::TestUnaryEndpointsMatchSpec` — the same check for REST.

Streaming RPCs are excluded by design (unary REST has no streaming counterpart;
in gRPC they are told apart by returning a stream-client interface, not a
`*Response` pointer). When adding/removing an endpoint, both parity tests must
stay green — that is what keeps the two transports in lockstep.

### Error model

No broad sentinel errors. The typed error *is* the category:

- `grpc.ConfigError` / `rest.ConfigError` — invalid construction input; match with `errors.As`.
- `rest.ResponseError` (non-2xx, has `StatusCode`) and `rest.RequestError` (per-stage failure, has `Op`) are aliases re-exported from restkit so callers needn't import restkit.
- `money` is the exception: it exposes `ErrConversion` / `ErrOverflow` sentinels matched with `errors.Is`.

## Conventions

- Go 1.26. Linting: `gofumpt` (extra-rules) + `golines` at **80 cols** + `modernize`. Keep lines ≤ 80.
- Prefer global `slog`/otel over dependency injection. A `ctx` parameter in a constructor is for trace/log context propagation (`slog.InfoContext`), not connection lifecycle.
- Keep `Client` service fields **exported** — no accessor methods.
- This repo is scaffolded from `acidsailor/go-scaffolds` (Copier). Do not hand-edit `.copier-answers.yml`; pull tooling updates with `task update`.

## Releases

Library release via GoReleaser (`builds: skip` — no binaries; GitHub serves the
source tarball). **Only push a git tag** — `goreleaser`/CI cuts the GitHub
Release and changelog. Do not run `gh release create` manually. Changelog
excludes `docs:`/`test:`/`chore:` commits, so use Conventional Commit prefixes.
