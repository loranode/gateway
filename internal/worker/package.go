// Package worker owns the whole conversation with the node: it dials the radio,
// asks for the node database, decodes the protobuf stream into models, stores
// them in the registry, and keeps the link alive — reconnecting for as long as
// it runs.
package worker

import (
	"time"

	"github.com/loranode/gateway/internal/services/events"
	"github.com/loranode/gateway/internal/services/registry"
	"github.com/loranode/gateway/internal/services/transport"
)

// Worker owns the ingest loop over a single radio transport.
type Worker struct {
	transport *transport.Service
	registry  *registry.Service
	events    *events.Service
}

const (
	// wantConfigID is echoed back by the node in config_complete_id once the
	// initial node-database dump finishes.
	wantConfigID = 0x4d4d // "MM"
	// heartbeatEvery keeps the API connection awake; the firmware drops idle
	// clients after ~15 minutes.
	heartbeatEvery = 5 * time.Minute
	// reconnectDelay is how long we wait before redialing after a dropped link.
	reconnectDelay = 3 * time.Second
	// broadcast is the recipient value that marks a channel (broadcast) packet;
	// any other recipient is a direct message addressed to a specific node.
	broadcast = 0xffffffff
)

// New builds a worker over the transport, registry and events services.
func New(tr *transport.Service, reg *registry.Service, ev *events.Service) *Worker {
	return &Worker{transport: tr, registry: reg, events: ev}
}
