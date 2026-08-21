// Package controller implements the generated MeshService and CallbackService
// REST APIs: each method answers one endpoint via the registry, events and mesh
// services.
package controller

import (
	"github.com/loranode/gateway/api"
	"github.com/loranode/gateway/internal/services/events"
	"github.com/loranode/gateway/internal/services/mesh"
	"github.com/loranode/gateway/internal/services/registry"
)

// Controller implements the MeshService and CallbackService REST interfaces.
type Controller struct {
	api.UnimplementedMeshServiceWebServer
	api.UnimplementedCallbackServiceWebServer

	registry *registry.Service
	events   *events.Service
	mesh     *mesh.Service
}

// New builds a controller over the registry, events and mesh services.
func New(reg *registry.Service, ev *events.Service, m *mesh.Service) *Controller {
	return &Controller{registry: reg, events: ev, mesh: m}
}
