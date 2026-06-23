.PHONY: proto proto-lint proto-breaking tidy

BUF ?= buf

proto:
	$(BUF) dep update
	$(BUF) generate

proto-lint:
	$(BUF) lint

proto-breaking:
	$(BUF) breaking --against '.git#branch=main'

tidy:
	go mod tidy
