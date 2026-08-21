package events

import (
	"log/slog"

	"github.com/loranode/gateway/internal/models"
)

// NodeCreated broadcasts a node_created event.
func (s *Service) NodeCreated(num uint32) {
	slog.Info("event", "type", typeNodeCreated, "node", num)
	s.publish(models.Event{NodeID: num, EventType: typeNodeCreated})
}
