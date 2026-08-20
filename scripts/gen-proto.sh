#!/usr/bin/env bash
#
# Regenerate the Go Meshtastic contracts from the pinned protobufs submodule.
#
# We generate only the transitive closure of mesh.proto — the messages this
# service actually needs (mesh, portnums, telemetry and their imports). The full
# tree is deliberately NOT generated: on the upstream master admin.proto's
# `AS3935_config` and telemetry.proto's `AS3935Config` both map to the same Go
# identifier, which will not compile in a single package. admin.proto is device
# administration a read-only monitor never uses, so excluding it is also leaner.
#
# The upstream .proto files declare go_package = github.com/meshtastic/go/generated,
# so every file is remapped onto this module's package via -M flags and written,
# with module= stripping the module prefix, into pkg/meshtastic/generated.
#
# Requires: protoc, protoc-gen-go (on PATH). Run after updating the submodule.
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
PROTO_DIR="$ROOT/third_party/protobufs"
OUT_PATH="github.com/loranode/gateway/api/meshtastic"
OUT_PKG="${OUT_PATH};meshtastic" # ;name overrides the upstream package clause
OUT_DIR="$ROOT/api/meshtastic"

# Transitive closure of meshtastic/mesh.proto. Keep in sync with the imports if
# the service starts using additional message types.
PROTOS=(
	meshtastic/mesh.proto
	meshtastic/channel.proto
	meshtastic/config.proto
	meshtastic/device_ui.proto
	meshtastic/module_config.proto
	meshtastic/portnums.proto
	meshtastic/telemetry.proto
	meshtastic/xmodem.proto
	meshtastic/atak.proto
)

if [ ! -d "$PROTO_DIR/meshtastic" ]; then
	echo "protobufs submodule missing; run: git submodule update --init" >&2
	exit 1
fi

cd "$PROTO_DIR"

mflags=()
for p in "${PROTOS[@]}"; do
	mflags+=("--go_opt=M${p}=${OUT_PKG}")
done

rm -f "$OUT_DIR"/*.pb.go
mkdir -p "$OUT_DIR"

protoc \
	--proto_path=. \
	--go_out="$ROOT" \
	--go_opt=module=github.com/loranode/gateway \
	"${mflags[@]}" \
	"${PROTOS[@]}"

echo "generated $(find "$OUT_DIR" -name '*.pb.go' | wc -l | tr -d ' ') files into api/meshtastic"
