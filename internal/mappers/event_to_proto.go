package mappers

import (
	"github.com/loranode/gateway/api"
	"github.com/loranode/gateway/internal/models"
)

// EventToProto maps a domain event to its proto form.
func EventToProto(e models.Event) *api.Event {
	return &api.Event{
		NodeId:    e.NodeID,
		EventType: e.EventType,
		MessageId: e.MessageID,
		ChannelId: e.ChannelID,
	}
}
