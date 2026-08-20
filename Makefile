# Meshtastic node proxy — developer tasks.
#
# Tooling expected on PATH (or under $(go env GOPATH)/bin): protoc,
# protoc-gen-go, protoc-gen-go-rest, golangci-lint, golangarch-lint.

GOBIN := $(shell go env GOPATH)/bin
BIN   := bin/app

.PHONY: generate lint build

# generate regenerates all generated code: the Meshtastic contracts from the
# pinned protobufs submodule (api/meshtastic) and our REST API from
# proto/proxy.proto (api/rest).
generate:
	./scripts/gen-proto.sh
	./scripts/gen-rest.sh

# lint runs the Go linters and the custom layered-architecture linter.
lint:
	go vet ./cmd/app
	golangci-lint run ./...
	$(GOBIN)/golangarch-lint lint .

# build compiles the service binary into bin/.
build:
	go build -o $(BIN) ./cmd/app
