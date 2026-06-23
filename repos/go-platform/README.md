# Go Platform

Shared Go libraries used across microservices in the polyrepo.

## Packages

| Package | Purpose |
|---------|---------|
| `pkg/config` | Environment-based configuration |
| `pkg/logger` | Zerolog setup |
| `pkg/observability` | OpenTelemetry tracing + Prometheus metrics |
| `pkg/grpcutil` | gRPC server/client with OTel instrumentation |

## Usage

```go
import (
    "github.com/imkhoirularifin/go-platform/pkg/config"
    "github.com/imkhoirularifin/go-platform/pkg/logger"
)
```

## Versioning

Tag releases and pin in service `go.mod`:

```go
require github.com/imkhoirularifin/go-platform v0.1.0
```

## Local development

When using the meta-repo orchestration layout, `go.work` at the repository root links this module locally. Services also include `replace` directives for sibling checkout.
