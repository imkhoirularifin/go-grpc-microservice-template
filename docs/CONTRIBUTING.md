# Contributing

Thank you for contributing. This template uses a **polyrepo** layout: each directory under `repos/` is an independent repository with its own `go.mod`, CI, and release cycle. The root repository orchestrates local development and deployment.

## Repositories

| Repo | Path | When to change |
|------|------|----------------|
| proto-contracts | `repos/proto-contracts` | API schema changes |
| go-platform | `repos/go-platform` | Shared libraries used by 2+ services |
| gateway-service | `repos/gateway-service` | HTTP routes, gRPC client wiring |
| user-service | `repos/user-service` | User domain logic |
| meta (this repo) | root | Tilt, K8s manifests, docs, orchestration |

## Prerequisites

| Tool | Version | Purpose |
|------|---------|---------|
| Go | 1.22+ | Application runtime |
| [Buf CLI](https://buf.build/docs/installation) | latest | Protobuf linting and codegen |
| [Tilt](https://docs.tilt.dev/install.html) | latest | Local Kubernetes development |
| Docker | latest | Container builds |
| kubectl | latest | Kubernetes CLI |

## Getting started

```bash
git clone https://github.com/imkhoirularifin/go-grpc-microservice-template.git
cd go-grpc-microservice-template

# Optional: use git submodules instead of vendored repos
# git submodule update --init --recursive

make proto
make tidy
```

The root `go.work` file links all local repos for development without publishing modules.

## Local development with Tilt

```bash
kind create cluster --name go-grpc-template
tilt up
```

Tilt builds from the meta-repo root using each service's Dockerfile and `go.work`.

## Development workflow

1. Identify which repo your change belongs to (see table above).
2. Create a feature branch in that repository.
3. If changing `.proto` files, work in `repos/proto-contracts` first:
   ```bash
   cd repos/proto-contracts
   make proto-lint
   make proto
   ```
4. Update dependent services and run tests from the meta repo:
   ```bash
   make test
   make build
   ```
5. Commit with a clear message:
   ```
   feat(gateway-service): add order routes
   fix(proto-contracts): correct pagination field
   chore(go-platform): export retry helper
   ```

## Code conventions

Follow patterns from [go-fiber-template](https://github.com/imkhoirularifin/go-fiber-template):

- **Configuration**: `go-platform/pkg/config` with environment variables
- **Logging**: `go-platform/pkg/logger` (Zerolog)
- **HTTP handlers**: thin handlers, logic in `internal/service`
- **Cross-service APIs**: defined only in `proto-contracts`

## Dependency management

### Local development

`go.work` at the repository root links all modules. Service `go.mod` files include `replace` directives for sibling repos:

```go
replace (
    github.com/imkhoirularifin/go-platform => ../go-platform
    github.com/imkhoirularifin/proto-contracts => ../proto-contracts
)
```

### Production

Pin tagged versions in each service `go.mod`:

```go
require (
    github.com/imkhoirularifin/go-platform v0.1.0
    github.com/imkhoirularifin/proto-contracts v1.0.0
)
```

Remove `replace` directives before publishing service releases.

## Pull request checklist

- [ ] Changes are in the correct repository
- [ ] `make proto-lint` passes (if proto changed)
- [ ] `make test` passes from meta repo root
- [ ] `make build` succeeds
- [ ] Generated code committed when `.proto` files change
- [ ] README updated in affected repo(s)

## Splitting into real separate repos

See the root [README](../README.md#splitting-into-separate-github-repos) for instructions on pushing each `repos/*` directory to its own GitHub repository and wiring submodules.

## License

By contributing, you agree that your contributions will be licensed under the same license as the project.
