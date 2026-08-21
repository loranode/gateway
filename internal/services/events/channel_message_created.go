package events

import (
	"log/slog"

	"github.com/loranode/gateway/internal/models"
)

// ChannelMessageCreated broadcasts a channel_message_created event for a broadcast.
func (s *Service) ChannelMessageCreated(from, messageID, channel uint32) {
	slog.Info("event", "type", typeChannelMessageCreated, "from", from, "message", messageID, "channel", channel)
	s.publish(models.Event{NodeID: from, EventType: typeChannelMessageCreated, MessageID: &messageID, ChannelID: &channel})
}
