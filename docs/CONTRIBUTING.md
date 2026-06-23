# Contributing

Thank you for contributing to the Go gRPC Microservice Template. This guide explains how to set up your environment, follow conventions, and submit changes.

## Prerequisites

| Tool | Version | Purpose |
|------|---------|---------|
| Go | 1.22+ | Application runtime |
| [Buf CLI](https://buf.build/docs/installation) | latest | Protobuf linting and code generation |
| [Tilt](https://docs.tilt.dev/install.html) | latest | Local Kubernetes development |
| Docker | latest | Container builds |
| kubectl | latest | Kubernetes CLI |
| kind or minikube | optional | Local Kubernetes cluster |

## Getting started

```bash
git clone https://github.com/imkhoirularifin/go-grpc-microservice-template.git
cd go-grpc-microservice-template

# Copy environment variables
cp .env.example .env

# Initialize proto contracts (submodule or vendored example)
git submodule update --init --recursive

# Generate protobuf code
make proto

# Download Go dependencies
go mod tidy
```

## Local development with Tilt

Tilt orchestrates protobuf generation, Docker builds, and Kubernetes deployments for local development.

```bash
# Start a local cluster (example with kind)
kind create cluster --name go-grpc-template

# Run Tilt
tilt up
```

Open the Tilt UI (usually http://localhost:10350) to monitor resources. Port forwards:

- Gateway HTTP: http://localhost:8080
- User gRPC: localhost:50051
- Gateway metrics: http://localhost:9090/metrics
- User metrics: http://localhost:9091/metrics

## Project structure

```
.
├── cmd/                    # Service entrypoints
│   ├── gateway/            # Fiber HTTP gateway
│   └── user/               # gRPC user service
├── internal/               # Private service code
├── pkg/                    # Shared libraries (config, observability, grpc)
├── lib/                    # HTTP helpers (dto, middleware, common)
├── proto/contracts/        # Centralized protobuf (git submodule)
├── gen/go/                 # Generated Go protobuf code (do not edit)
├── deploy/                 # Docker, Kubernetes, monitoring manifests
├── docs/                   # Documentation
├── buf.yaml                # Buf module configuration
├── buf.gen.yaml            # Buf code generation plugins
└── Tiltfile                # Local development orchestration
```

## Development workflow

1. Create a feature branch from `main`:
   ```bash
   git checkout -b feature/my-change
   ```

2. If you change `.proto` files, regenerate code:
   ```bash
   make proto
   make proto-lint
   ```

3. Run tests and build:
   ```bash
   make test
   make build
   ```

4. Commit with a clear message:
   ```
   feat(gateway): add user list endpoint
   fix(user): handle missing pagination
   chore(proto): bump contracts to v1.1.0
   ```

5. Open a pull request against `main`.

## Code conventions

Follow patterns from [go-fiber-template](https://github.com/imkhoirularifin/go-fiber-template):

- **Configuration**: environment variables via `caarlos0/env`, loaded in `pkg/config`
- **Logging**: structured logging with `zerolog`
- **HTTP handlers**: thin handlers in `internal/<service>/handler`, business logic in `service`
- **Error handling**: centralized Fiber error handler in `lib/common`
- **Dependency wiring**: explicit constructor injection (no global state in handlers)

### Observability

- All services initialize OpenTelemetry tracing via `pkg/observability`
- Prometheus metrics are exposed on dedicated ports (`/metrics`)
- Propagate trace context across HTTP → gRPC calls using OTel middleware

### Protobuf contracts

- Proto definitions live in a **centralized repository** linked as a git submodule at `proto/contracts`
- Version contracts using git tags (`v1.0.0`, `v1.1.0`)
- Never edit generated code in `gen/go/`
- Run `make proto-breaking` before merging proto changes

## Pull request checklist

- [ ] `make proto-lint` passes
- [ ] `make test` passes
- [ ] `make build` succeeds
- [ ] Generated code is committed when `.proto` files change
- [ ] Documentation updated if behavior or setup changes
- [ ] No secrets committed (use `.env`, not version control)

## Reporting issues

Open a GitHub issue with:

- Steps to reproduce
- Expected vs actual behavior
- Environment (OS, Go version, Tilt version)
- Relevant logs or traces

## License

By contributing, you agree that your contributions will be licensed under the same license as the project.
