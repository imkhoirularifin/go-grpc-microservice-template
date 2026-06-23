.PHONY: help proto proto-lint proto-breaking tidy test build run-gateway run-user docker-build tilt

MODULE := github.com/imkhoirularifin/go-grpc-microservice-template
BUF ?= buf

help:
	@echo "Targets:"
	@echo "  proto          Generate Go code from protobuf definitions"
	@echo "  proto-lint     Lint protobuf definitions"
	@echo "  proto-breaking Check breaking changes against main"
	@echo "  tidy           Run go mod tidy"
	@echo "  test           Run unit tests"
	@echo "  build          Build gateway and user service binaries"
	@echo "  run-gateway    Run HTTP gateway locally"
	@echo "  run-user       Run user gRPC service locally"
	@echo "  docker-build   Build Docker images"
	@echo "  tilt           Start local development with Tilt"

proto:
	$(BUF) dep update
	$(BUF) generate

proto-lint:
	$(BUF) lint

proto-breaking:
	$(BUF) breaking --against '.git#branch=main'

tidy:
	go mod tidy

test:
	go test ./...

build: proto
	go build -o bin/gateway ./cmd/gateway
	go build -o bin/user ./cmd/user

run-gateway:
	go run ./cmd/gateway

run-user:
	go run ./cmd/user

docker-build:
	docker build -f deploy/docker/gateway.Dockerfile -t go-grpc-template/gateway:latest .
	docker build -f deploy/docker/user.Dockerfile -t go-grpc-template/user:latest .

tilt:
	tilt up
