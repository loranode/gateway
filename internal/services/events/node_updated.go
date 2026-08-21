package events

import "github.com/loranode/gateway/internal/models"

// NodeUpdated broadcasts a node_updated event.
func (s *Service) NodeUpdated(num uint32) {
	s.publish(models.Event{NodeID: num, EventType: typeNodeUpdated})
}
