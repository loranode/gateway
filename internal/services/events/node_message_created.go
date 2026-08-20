package events

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/loranode/gateway/internal/models"
)

// NodeMessageCreated posts a node_message_created event for a direct message,
// carrying the sender node number and the message id.
func (s *Service) NodeMessageCreated(ctx context.Context, from, messageID uint32) {
	slog.Info("event registered", "event", typeNodeMessageCreated, "node", from)

	body, err := json.Marshal(models.Event{
		NodeID:    from,
		EventType: typeNodeMessageCreated,
		MessageID: &messageID,
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
