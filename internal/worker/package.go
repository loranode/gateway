// Package worker runs the ingest loop: it reads decoded events from the node
// (mesh service), stores them via the registry and emits webhook events. It
// coordinates several services, which a single service may not do.
package worker

import (
	"github.com/loranode/gateway/internal/services/events"
	"github.com/loranode/gateway/internal/services/mesh"
	"github.com/loranode/gateway/internal/services/registry"
)

// Worker folds the node's decoded event stream into the registry and events.
type Worker struct {
	mesh     *mesh.Service
	registry *registry.Service
	events   *events.Service
}

// broadcast marks a channel (broadcast) message; any other recipient is direct.
const broadcast = 0xffffffff

// New builds the ingest worker over the mesh, registry and events services.
func New(m *mesh.Service, reg *registry.Service, ev *events.Service) *Worker {
	return &Worker{mesh: m, registry: reg, events: ev}
}
