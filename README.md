# tinvest

Go client library for the [T-Invest (Tinkoff Investments) gRPC API](https://opensource.tbank.ru/invest/invest-contracts).

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
)

func main() {
    ctx := context.Background()

    connConfig := tinvest.NewConnConfig(tinvest.EndpointProduction, "your-api-token")
    conn, err := tinvest.NewConn(ctx, connConfig)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    client, err := tinvest.NewClient(conn, tinvest.NewClientConfig())
    if err != nil {
        log.Fatal(err)
    }

    // Use client.Instruments, client.MarketData, client.Orders, etc.
    _ = client
}
```

## Configuration

### ConnConfig

`ConnConfig` holds gRPC connection settings and is created with a required endpoint and API token:

```go
// Production environment
config := tinvest.NewConnConfig(tinvest.EndpointProduction, token)

// Sandbox environment (for testing without real money)
config := tinvest.NewConnConfig(tinvest.EndpointSandbox, token)

// Optionally set a custom app name (sent as x-app-name header)
config = config.WithAppName("my-trading-bot")
```

| Constant             | Value                                             |
|----------------------|---------------------------------------------------|
| `EndpointProduction` | `invest-public-api.tinkoff.ru:443`                |
| `EndpointSandbox`    | `sandbox-invest-public-api.tinkoff.ru:443`        |

### ClientConfig

`ClientConfig` holds client-level settings:

```go
config := tinvest.NewClientConfig()
```

## API Services

`Client` exposes all T-Invest gRPC services as typed sub-clients:

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

The underlying proto-generated interfaces live in the `pb` sub-package.

## Financial Value Helpers

T-Invest encodes monetary values as protobuf `Quotation` and `MoneyValue` messages. The library provides helpers to convert these to and from [`udecimal.Decimal`](https://github.com/quagmt/udecimal):

```go
import "github.com/acidsailor/tinvest"

// Quotation ↔ Decimal
d, err := tinvest.QuotationToDecimal(q)
q, err := tinvest.DecimalToQuotation(d)

// MoneyValue ↔ Decimal (currency field is preserved separately)
d, err := tinvest.MoneyValueToDecimal(m)
m, err := tinvest.DecimalToMoneyValue(d, "RUB")
```

Financial values support up to 9 fractional digits.

## Error Handling

Errors originating inside the `tinvest` package (validation, conversion) wrap `ErrClient` and can be detected with `errors.Is`:

```go
if errors.Is(err, tinvest.ErrClient) {
    // configuration or conversion error
}
```

Errors returned by gRPC RPC calls are passed through unwrapped as standard gRPC status errors.

## OpenTelemetry

The gRPC connection is automatically instrumented with OpenTelemetry via `otelgrpc`. Traces and metrics are exported through any configured global OTEL provider.

## Connection Lifecycle

`NewConn` returns a lazily-dialed `*grpc.ClientConn` — no TCP connection is established until the first RPC call. The caller owns the connection and is responsible for closing it:

```go
conn, err := tinvest.NewConn(ctx, connConfig)
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

## License

See [LICENSE](LICENSE).
