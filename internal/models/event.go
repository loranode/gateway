package models

// Event is the payload POSTed to registered webhooks. NodeID is the node the
// event concerns (for messages, the sender); EventType is one of "node_created",
// "node_updated", "node_message_created", "channel_message_created". MessageID is
// set for message events; ChannelID is set only for channel (broadcast) messages.
type Event struct {
	NodeID    uint32  `json:"node_id"`
	EventType string  `json:"event_type"`
	MessageID *uint32 `json:"message_id,omitempty"`
	ChannelID *uint32 `json:"channel_id,omitempty"`
}
