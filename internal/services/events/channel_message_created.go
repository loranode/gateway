package events

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/loranode/gateway/internal/models"
)

// ChannelMessageCreated posts a channel_message_created event for a broadcast
// message, carrying the sender node number, the message id and the channel index.
func (s *Service) ChannelMessageCreated(ctx context.Context, from, messageID, channel uint32) {
	slog.Info("event registered", "event", typeChannelMessageCreated, "node", from)

	body, err := json.Marshal(models.Event{
		NodeID:    from,
		EventType: typeChannelMessageCreated,
		MessageID: &messageID,
		ChannelID: &channel,
	})
	if err != nil {
		slog.Error("marshal event", "err", err)

		return
	}

	urls, err := s.store.Callbacks(ctx)
	if err != nil {
		slog.Error("load callbacks", "err", err)

		return
	}

	for _, url := range urls {
		if err := s.webhook.Post(ctx, url, body); err != nil {
			slog.Warn("webhook delivery failed", "err", err, "url", url)
		}
	}
}
