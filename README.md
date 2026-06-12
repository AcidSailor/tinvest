# tinvest

Go client library for the [T-Invest (T-Bank Investments) gRPC and REST APIs](https://opensource.tbank.ru/invest/invest-contracts).

The transport clients live in sub-packages so the root `tinvest` package (endpoint constants, `AppName`, sentinel errors) stays dependency-light:

- `tinvest/grpc` — gRPC client (`NewConn` / `NewClient`) and proto-typed money helpers
- `tinvest/rest` — REST gateway client
- `tinvest/money` — protobuf-free `udecimal.Decimal` ↔ units/nano conversions

## Installation

```bash
go get github.com/acidsailor/tinvest
```

## Quick Start

```go
package main

import (
    "context"
    "log"

    "github.com/acidsailor/tinvest"
    "github.com/acidsailor/tinvest/grpc"
)

func main() {
    ctx := context.Background()

    conn, err := grpc.NewConn(ctx, tinvest.EndpointProduction, "your-api-token")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    client, err := grpc.NewClient(conn)
    if err != nil {
        log.Fatal(err)
    }

    // Use client.Instruments, client.MarketData, client.Orders, etc.
    _ = client
}
```

## Configuration

### Connection

`grpc.NewConn` takes a required endpoint and API token, plus optional functional options. Endpoint constants live in the root `tinvest` package:

```go
// Production environment
conn, err := grpc.NewConn(ctx, tinvest.EndpointProduction, token)

// Sandbox environment (for testing without real money)
conn, err := grpc.NewConn(ctx, tinvest.EndpointSandbox, token)

// Optionally set a custom app name (sent as x-app-name header)
conn, err := grpc.NewConn(ctx, tinvest.EndpointProduction, token, grpc.WithAppName("my-trading-bot"))
```

| Constant             | Value                                             |
|----------------------|---------------------------------------------------|
| `EndpointProduction` | `invest-public-api.tinkoff.ru:443`                |
| `EndpointSandbox`    | `sandbox-invest-public-api.tinkoff.ru:443`        |

`grpc.NewClient` then wraps the connection with the typed service sub-clients:

```go
client, err := grpc.NewClient(conn)
```

## API Services

`grpc.Client` exposes all T-Invest gRPC services as typed sub-clients:

| Field                 | Service                                             |
|-----------------------|-----------------------------------------------------|
| `Instruments`         | Instrument and security lookups                     |
| `MarketData`          | Candles, prices, order books, trading status        |
| `MarketDataStream`    | Streaming market data                               |
| `Operations`          | Portfolio, positions, broker reports                |
| `OperationsStream`    | Streaming portfolio, positions, and operations      |
| `Orders`              | Place, cancel, and manage orders                    |
| `OrdersStream`        | Streaming order state and trades                    |
| `StopOrders`          | Stop orders management                              |
| `Sandbox`             | Sandbox account management                          |
| `Users`               | User account information and settings               |
| `Signals`             | Trading signals and strategies                      |

The underlying proto-generated interfaces live in the `grpc/pb` sub-package.

## Financial Value Helpers

T-Invest encodes monetary values as protobuf `Quotation` and `MoneyValue` messages. The library provides helpers to convert these to and from [`udecimal.Decimal`](https://github.com/quagmt/udecimal):

```go
import "github.com/acidsailor/tinvest/grpc"

// Quotation ↔ Decimal
d, err := grpc.QuotationToDecimal(q)
q, err := grpc.DecimalToQuotation(d)

// MoneyValue ↔ Decimal (currency field is preserved separately)
d, err := grpc.MoneyValueToDecimal(m)
m, err := grpc.DecimalToMoneyValue(d, "RUB")
```

Financial values support up to 9 fractional digits. The protobuf-free
units/nano math underneath lives in the `tinvest/money` package.

## Error Handling

Invalid construction input surfaces as a typed `*ConfigError`, owned by each
transport and matched with `errors.As`:

```go
var ce *grpc.ConfigError // or *rest.ConfigError
if errors.As(err, &ce) {
    // ce.Reason names the missing/invalid value, e.g. "empty token"
}
```

The typed errors are:

- `grpc.ConfigError` / `rest.ConfigError` — invalid `NewConn` / `NewClient` input
- `rest.RequestError` — a REST call failed before a result (`Op` names the stage)
- `rest.ResponseError` — a non-2xx REST response (`StatusCode` and raw body)
- `money.ErrConversion` — invalid input converting between units/nano and decimal
- `money.ErrOverflow` — a value does not fit the target representation

`money` keeps `ErrConversion` / `ErrOverflow` as sentinels (match with
`errors.Is`); everywhere else the typed error is the category. Nil-argument
errors (e.g. a nil `*pb.Quotation`, or a nil `conn` to `grpc.NewClient`) are
returned as plain `errors.New` values — they carry a `tinvest:` message prefix
but are not meant to be matched on. Errors returned by gRPC RPC calls are passed
through unwrapped as standard gRPC status errors.

## OpenTelemetry

The gRPC connection is automatically instrumented with OpenTelemetry via `otelgrpc`. Traces and metrics are exported through any configured global OTEL provider.

## Connection Lifecycle

`grpc.NewConn` returns a lazily-dialed `*grpc.ClientConn` — no TCP connection is established until the first RPC call. The caller owns the connection and is responsible for closing it:

```go
conn, err := grpc.NewConn(ctx, connConfig)
if err != nil { ... }
defer conn.Close()
```

## Development

Regenerate protobuf bindings:

```bash
task proto
```

Run tests:

```bash
task test
```

Run linter:

```bash
task lint
```

Proto sources are fetched from `https://opensource.tbank.ru/invest/invest-contracts.git` into `tmp/invest-contracts/` at build time.

## Disclaimer

This library is provided "as is", without warranty of any kind. The author assumes **no financial, legal, or other liability** for any losses, damages, or consequences arising from the use of this library, including but not limited to losses incurred through trading, order placement, or interaction with the T-Invest API.

Nothing in this library, its documentation, or examples constitutes **investment advice, a recommendation, or solicitation** to buy or sell any financial instrument. All trading decisions are solely the responsibility of the user. Consult a licensed financial advisor before making investment decisions.

## Отказ от ответственности

Библиотека предоставляется «как есть», без каких-либо гарантий. Автор **не несёт финансовой, юридической или иной ответственности** за любые убытки, ущерб или последствия, возникшие в результате использования этой библиотеки, включая, но не ограничиваясь, убытки от торговли, выставления ордеров или взаимодействия с API T-Invest.

Ничто в этой библиотеке, её документации или примерах **не является индивидуальной инвестиционной рекомендацией**, предложением или побуждением к покупке или продаже каких-либо финансовых инструментов. Все торговые решения принимаются пользователем самостоятельно и под его ответственность. Перед принятием инвестиционных решений проконсультируйтесь с лицензированным финансовым советником.

## License

See [LICENSE](LICENSE).
