package events

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/loranode/gateway/internal/models"
)

// NodeCreated posts a node_created event for the given node number.
func (s *Service) NodeCreated(ctx context.Context, num uint32) {
	slog.Info("event registered", "event", typeNodeCreated, "node", num)

	body, err := json.Marshal(models.Event{NodeID: num, EventType: typeNodeCreated})
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
