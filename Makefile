GOBIN := $(shell go env GOPATH)/bin
BIN   := bin/app

.PHONY: generate lint build

# generate regenerates our REST API from proto/gateway.proto (api). The
# Meshtastic contracts now come from the github.com/loranode/meshtastic library.
generate:
	./scripts/gen-rest.sh

# lint runs the Go linters and the custom layered-architecture linter.
lint:
	go vet ./cmd/app
	golangci-lint run ./...
	$(GOBIN)/golangarch-lint docs .
	$(GOBIN)/golangarch-lint lint .

# build compiles the service binary into bin/.
build:
	go build -o $(BIN) ./cmd/app
