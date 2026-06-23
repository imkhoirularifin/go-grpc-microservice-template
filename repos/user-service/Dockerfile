# syntax=docker/dockerfile:1

FROM golang:1.22-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /src
COPY go.work ./
COPY repos/go-platform ./repos/go-platform
COPY repos/proto-contracts ./repos/proto-contracts
COPY repos/user-service ./repos/user-service

WORKDIR /src/repos/user-service
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/user ./cmd

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /out/user /user
USER nonroot:nonroot
EXPOSE 50051 9091
ENTRYPOINT ["/user"]
