# tinvest

A Go client library for the [T-Invest](https://russianinvestments.github.io/investAPI/)
(Tinkoff Investments) API.

It provides two independent transports over the same upstream contract — a
**gRPC** client and a **REST** gateway client — generated from one pinned
contract release, plus a transport-agnostic **money** package for the API's
`units`+`nano` decimal representation.

```sh
go get github.com/acidsailor/tinvest
```

Requires Go 1.26+. Authenticate with a T-Invest API token; target either the
production or sandbox endpoint via the constants in the root package.

## Packages

| Import path                                  | Purpose                                                                 |
| -------------------------------------------- | ----------------------------------------------------------------------- |
| `github.com/acidsailor/tinvest`              | Shared primitives: endpoint constants and the `x-app-name` value.       |
| `github.com/acidsailor/tinvest/grpc`         | gRPC client (`NewConn` / `NewClient`).                                   |
| `github.com/acidsailor/tinvest/rest`         | REST gateway client (`NewClient`).                                       |
| `github.com/acidsailor/tinvest/money`        | `units`+`nano` ⇄ `udecimal.Decimal` conversions and money formatting.   |

Endpoint constants (root package): `EndpointProduction`, `EndpointSandbox`
(gRPC, `host:443`) and `EndpointProductionREST`, `EndpointSandboxREST` (HTTPS
gateway URLs).

## gRPC usage

`NewConn` builds a configured `*grpc.ClientConn` (lazy dial, TLS 1.2+, OpenTelemetry
stats handler, bearer-token + `x-app-name` auth interceptors). `NewClient` wraps
the connection with one field per service. **The client does not own the
connection — you are responsible for closing it.**

```go
package main

import (
	"context"
	"log"

	"github.com/acidsailor/tinvest"
	tgrpc "github.com/acidsailor/tinvest/grpc"
	pb "github.com/acidsailor/tinvest/grpc/pb"
)

func main() {
	ctx := context.Background()

	conn, err := tgrpc.NewConn(ctx, tinvest.EndpointSandbox, "<token>")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client, err := tgrpc.NewClient(conn)
	if err != nil {
		log.Fatal(err)
	}

	accounts, err := client.Users.GetAccounts(ctx, &pb.GetAccountsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d accounts", len(accounts.GetAccounts()))
}
```

Service fields cover `Instruments`, `MarketData`, `Operations`, `Orders`,
`StopOrders`, `Sandbox`, `Users`, `Signals`, and their streaming variants
(`MarketDataStream`, `OperationsStream`, `OrdersStream`). Override the
`x-app-name` header with `tgrpc.WithAppName(...)`.

## REST usage

The REST client targets the same operations over the JSON gateway. Each call
is a `POST`; request/response models live in the `rest` package (generated from
the OpenAPI spec).

```go
client, err := rest.NewClient(tinvest.EndpointSandboxREST, "<token>")
if err != nil {
	log.Fatal(err)
}

accounts, err := client.Users.GetAccounts(ctx, &rest.V1GetAccountsRequest{})
```

Defaults to a 30s-timeout `*http.Client`; override with `rest.WithHTTPClient(...)`
and the app name with `rest.WithAppName(...)`. The REST client is immutable after
construction and safe for concurrent use.

## Money

The API encodes monetary values as `units` (int64) + `nano` (int32
nano-billionths). The `money` package converts to/from `udecimal.Decimal` with
no protobuf dependency; the `grpc` package adds adapters for the proto
`MoneyValue` / `Quotation` types.

```go
import "github.com/acidsailor/tinvest/money"

d, err := money.UnitsNanoToDecimal(250, 500_000_000) // 250.5
s, err := money.FormatMoney(250, 500_000_000, "RUB")  // "250.50 RUB"

// From proto types:
d, err := tgrpc.MoneyValueToDecimal(mv)
s, err := tgrpc.FormatQuotation(q)
```

Conversion failures are matched with `errors.Is` against `money.ErrConversion`
and `money.ErrOverflow`.

## Error handling

Errors are typed, not sentinel-based — match the concrete type with `errors.As`:

```go
var ce *rest.ConfigError    // or *grpc.ConfigError — invalid construction input
var re *rest.ResponseError  // non-2xx response; re.StatusCode
var rq *rest.RequestError   // per-stage failure; rq.Op (e.g. rest.OpSend)
```

## Development

Tasks run via [Task](https://taskfile.dev):

- `task test` — `go test -race ./...`
- `task check` — lint (autofix) + tests
- `task ci` — read-only fmt/lint, as enforced in CI

Both transports are generated from one pinned upstream release
(`CONTRACTS_VERSION` in `taskfile.yml`, the single source of truth for the proto
contracts and the REST spec): `task proto` regenerates `grpc/pb`, `task rest`
regenerates `rest/models.gen.go`. Generated code is committed. Parity tests
assert both transports expose exactly the unary operations the contract defines.

## Disclaimer

This library is provided "as is", without warranty of any kind. The author
assumes **no financial, legal, or other liability** for any losses, damages, or
consequences arising from the use of this library, including but not limited to
losses incurred through trading, order placement, or interaction with the
T-Invest API.

Nothing in this library, its documentation, or examples constitutes **investment
advice, a recommendation, or solicitation** to buy or sell any financial
instrument. All trading decisions are solely the responsibility of the user.
Consult a licensed financial advisor before making investment decisions.

## Отказ от ответственности

Библиотека предоставляется «как есть», без каких-либо гарантий. Автор **не несёт
финансовой, юридической или иной ответственности** за любые убытки, ущерб или
последствия, возникшие в результате использования этой библиотеки, включая, но
не ограничиваясь, убытки от торговли, выставления ордеров или взаимодействия с
API T-Invest.

Ничто в этой библиотеке, её документации или примерах **не является
индивидуальной инвестиционной рекомендацией**, предложением или побуждением к
покупке или продаже каких-либо финансовых инструментов. Все торговые решения
принимаются пользователем самостоятельно и под его ответственность. Перед
принятием инвестиционных решений проконсультируйтесь с лицензированным
финансовым советником.

## License

[AGPL-3.0](LICENSE).
