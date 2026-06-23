# Gateway Service

Fiber HTTP gateway that exposes REST endpoints and proxies to backend gRPC services.

## Dependencies

- [go-platform](../go-platform) — shared config, logging, observability
- [proto-contracts](../proto-contracts) — gRPC client stubs

## Run locally

```bash
cp .env.example .env
go run ./cmd
```

## Endpoints

- `GET /api/v1/healthz`
- `GET /api/v1/readyz`
- `GET /api/v1/users/:id`
- `GET /api/v1/users`

## Docker

Built from the meta-repo root:

```bash
docker build -f repos/gateway-service/Dockerfile -t gateway-service:latest .
```
