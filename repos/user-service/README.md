# User Service

Example gRPC microservice implementing the User API from proto-contracts.

## Dependencies

- [go-platform](../go-platform) — shared config, logging, observability, gRPC helpers
- [proto-contracts](../proto-contracts) — gRPC server stubs

## Run locally

```bash
cp .env.example .env
go run ./cmd
```

## Docker

Built from the meta-repo root:

```bash
docker build -f repos/user-service/Dockerfile -t user-service:latest .
```
