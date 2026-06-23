.PHONY: tidy test build run

tidy:
	go mod tidy

test:
	go test ./...

build:
	go build -o bin/user ./cmd

run:
	go run ./cmd
