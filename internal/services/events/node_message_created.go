package events

import (
	"log/slog"

	"github.com/loranode/gateway/internal/models"
)

// NodeMessageCreated broadcasts a node_message_created event for a direct message.
func (s *Service) NodeMessageCreated(from, messageID uint32) {
	slog.Info("event", "type", typeNodeMessageCreated, "from", from, "message", messageID)
	s.publish(models.Event{NodeID: from, EventType: typeNodeMessageCreated, MessageID: &messageID})
}
