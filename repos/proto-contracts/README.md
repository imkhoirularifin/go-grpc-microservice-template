# Proto Contracts

Centralized protobuf definitions for all microservices. This is a **standalone repository** in the polyrepo layout.

## Responsibilities

- Own all `.proto` files and API versioning
- Lint and enforce breaking-change rules with [Buf](https://buf.build/docs)
- Publish generated Go stubs as part of this module (`gen/go/`)

## Versioning

Tag releases with semantic versions:

```bash
git tag v1.0.0
git push origin v1.0.0
```

Services pin a specific tag in their `go.mod`:

```go
require github.com/imkhoirularifin/proto-contracts v1.0.0
```

## Generate Go code

```bash
make proto
make proto-lint
```

## Layout

```
common/v1/     Shared messages
user/v1/       User service API
gen/go/        Generated Go code (committed)
```

## Consuming from services

```go
import userv1 "github.com/imkhoirularifin/proto-contracts/gen/go/user/v1"
```
