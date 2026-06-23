# syntax=docker/dockerfile:1

FROM golang:1.22-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/gateway ./cmd/gateway

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /out/gateway /gateway
USER nonroot:nonroot
EXPOSE 8080 9090
ENTRYPOINT ["/gateway"]
