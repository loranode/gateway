#!/usr/bin/env bash
#
# Generate the Go API code from our own proto contracts in proto/ into api:
# protobuf models, the gRPC service, and the REST/websocket layer
# (merzzzl/proto-rest-api). api/ is a nested module the gateway and botkit share.
#
# Requires: protoc, protoc-gen-go, protoc-gen-go-grpc, protoc-gen-go-rest (PATH).
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MODULE="github.com/loranode/gateway"

# The restapi annotations ship with the proto-rest-api module; resolve their
# include path from the module cache.
RESTAPI_INC="$(go list -m -f '{{.Dir}}' github.com/merzzzl/proto-rest-api)"

# Well-known types (google/protobuf/*.proto) bundled with protoc.
WKT_INC="$(dirname "$(command -v protoc)")/../include"

OUT_DIR="$ROOT/api"
rm -f "$OUT_DIR"/*.pb.go "$OUT_DIR"/*_swagger.go "$OUT_DIR"/*_swagger.json
mkdir -p "$OUT_DIR"

protoc \
	--proto_path="$ROOT/proto" \
	--proto_path="$RESTAPI_INC" \
	--proto_path="$WKT_INC" \
	--go_out="$ROOT" --go_opt=module="$MODULE" \
	--go-grpc_out="$ROOT" --go-grpc_opt=module="$MODULE" \
	--go-rest_out="$ROOT" --go-rest_opt=module="$MODULE" \
	gateway.proto

echo "generated proto + gRPC + REST into api"

# Render Markdown API docs from the generated swagger JSON into SWAGGER.md at the
# repo root. The first 4 lines are the generator's title/preamble — trim them.
npx --yes swagger-markdown -i "$OUT_DIR/gateway_swagger.json" -o "$ROOT/SWAGGER.md.tmp"
sed '1,4d' "$ROOT/SWAGGER.md.tmp" >"$ROOT/SWAGGER.md" && rm "$ROOT/SWAGGER.md.tmp"

echo "generated SWAGGER.md at repo root"
