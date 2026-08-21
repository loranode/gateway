// Package mesh wraps the meshtastic repository, exposing the node's decoded event
// stream and text send to the rest of the app.
package mesh

import "github.com/loranode/gateway/internal/repositories/meshtastic"

// Service wraps the node transport repository.
type Service struct {
	repo *meshtastic.Repository
}

// New wraps a meshtastic repository.
func New(r *meshtastic.Repository) *Service {
	return &Service{repo: r}
}
