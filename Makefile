.PHONY: help proto proto-lint proto-breaking tidy test build docker-build tilt

BUF ?= buf
PROTO_DIR := repos/proto-contracts

help:
	@echo "Polyrepo orchestration targets:"
	@echo "  proto          Generate Go code in repos/proto-contracts"
	@echo "  proto-lint     Lint protobuf definitions"
	@echo "  proto-breaking Check breaking proto changes"
	@echo "  tidy           Run go mod tidy in all repos"
	@echo "  test           Run tests in all service repos"
	@echo "  build          Build gateway and user service binaries"
	@echo "  docker-build   Build Docker images"
	@echo "  tilt           Start local development with Tilt"

proto:
	cd $(PROTO_DIR) && $(BUF) dep update && $(BUF) generate

proto-lint:
	cd $(PROTO_DIR) && $(BUF) lint

proto-breaking:
	cd $(PROTO_DIR) && $(BUF) breaking --against '.git#branch=main'

tidy:
	cd repos/go-platform && go mod tidy
	cd repos/proto-contracts && go mod tidy
	cd repos/gateway-service && go mod tidy
	cd repos/user-service && go mod tidy

test:
	cd repos/go-platform && go test ./...
	cd repos/user-service && go test ./...
	cd repos/gateway-service && go test ./...

build:
	cd repos/gateway-service && go build -o ../../bin/gateway ./cmd
	cd repos/user-service && go build -o ../../bin/user ./cmd

docker-build:
	docker build -f repos/gateway-service/Dockerfile -t go-grpc-template/gateway:latest .
	docker build -f repos/user-service/Dockerfile -t go-grpc-template/user:latest .

tilt:
	tilt up
