// Package controller implements the generated MeshService and CallbackService
// REST APIs: each method answers one endpoint via the registry / events
// services.
package controller

import (
	"github.com/loranode/gateway/api/rest"
	"github.com/loranode/gateway/internal/services/events"
	"github.com/loranode/gateway/internal/services/registry"
)

// Controller implements the MeshService and CallbackService REST interfaces.
type Controller struct {
	rest.UnimplementedMeshServiceWebServer
	rest.UnimplementedCallbackServiceWebServer

	registry *registry.Service
	events   *events.Service
	send     func(text string, to, channel uint32) error
}

// New builds a controller over the registry, events service and send callback.
func New(reg *registry.Service, ev *events.Service, send func(text string, to, channel uint32) error) *Controller {
	return &Controller{registry: reg, events: ev, send: send}
}
