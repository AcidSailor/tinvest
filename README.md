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

    connConfig := grpc.NewConnConfig(tinvest.EndpointProduction, "your-api-token")
    conn, err := grpc.NewConn(ctx, connConfig)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    client, err := grpc.NewClient(conn, grpc.NewClientConfig())
    if err != nil {
        log.Fatal(err)
    }

    // Use client.Instruments, client.MarketData, client.Orders, etc.
    _ = client
}
```

## Configuration

### ConnConfig

`grpc.ConnConfig` holds gRPC connection settings and is created with a required endpoint and API token. Endpoint constants live in the root `tinvest` package:

```go
// Production environment
config := grpc.NewConnConfig(tinvest.EndpointProduction, token)

// Sandbox environment (for testing without real money)
config := grpc.NewConnConfig(tinvest.EndpointSandbox, token)

// Optionally set a custom app name (sent as x-app-name header) via a functional option
config = grpc.NewConnConfig(tinvest.EndpointProduction, token, grpc.WithAppName("my-trading-bot"))
```

| Constant             | Value                                             |
|----------------------|---------------------------------------------------|
| `EndpointProduction` | `invest-public-api.tinkoff.ru:443`                |
| `EndpointSandbox`    | `sandbox-invest-public-api.tinkoff.ru:443`        |

### ClientConfig

`grpc.ClientConfig` holds client-level settings:

```go
config := grpc.NewClientConfig()
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

Errors raised by the library's own validation and conversion logic wrap `tinvest.ErrClient` and can be detected with `errors.Is`:

```go
if errors.Is(err, tinvest.ErrClient) {
    // configuration or conversion error
}
```

Finer-grained sentinels (`tinvest.ErrNil`, `tinvest.ErrInvalidConfig`) are joined alongside `ErrClient`, so both match. Errors returned by gRPC RPC calls are passed through unwrapped as standard gRPC status errors.

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
