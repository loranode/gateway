// Package controller implements the generated MeshService (REST + gRPC) and the
// EventService gRPC event stream, answering each via the registry, events and
// mesh services.
package controller

import (
	"github.com/loranode/gateway/api"
	"github.com/loranode/gateway/internal/services/events"
	"github.com/loranode/gateway/internal/services/mesh"
	"github.com/loranode/gateway/internal/services/registry"
)

// Controller implements MeshService over both REST and gRPC and streams events
// over gRPC.
type Controller struct {
	api.UnimplementedMeshServiceWebServer
	api.UnimplementedMeshServiceServer
	api.UnimplementedEventServiceServer

	registry *registry.Service
	events   *events.Service
	mesh     *mesh.Service
}

// New builds a controller over the registry, events and mesh services.
func New(reg *registry.Service, ev *events.Service, m *mesh.Service) *Controller {
	return &Controller{registry: reg, events: ev, mesh: m}
}
