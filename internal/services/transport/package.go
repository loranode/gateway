// Package transport is the service in front of the radio repository: it exposes
// the node connection's raw framing primitives to the controller, which owns the
// connect/read loop.
package transport

import "github.com/loranode/gateway/internal/repositories/radio"

// Service wraps a single radio transport.
type Service struct {
	repo *radio.Repository
}

// New builds the transport service over repo.
func New(repo *radio.Repository) *Service {
	return &Service{repo: repo}
}
