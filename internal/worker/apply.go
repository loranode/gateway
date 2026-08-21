package worker

import (
	"context"

	"github.com/loranode/gateway/internal/models"
)

// apply folds one decoded event into the registry and emits the matching events.
func (w *Worker) apply(ctx context.Context, ev models.MeshEvent) {
	if ev.Node != nil {
		if created, err := w.registry.UpsertNode(ev.Node); err == nil {
			if created {
				w.events.NodeCreated(ctx, ev.Node.Num)
			} else {
				w.events.NodeUpdated(ctx, ev.Node.Num)
			}
		}
	}

	if ev.Message != nil {
		if inserted, err := w.registry.AddMessage(ev.Message); err == nil && inserted {
			if ev.Message.To == broadcast {
				w.events.ChannelMessageCreated(ctx, ev.Message.From, ev.Message.ID, ev.Message.Channel)
			} else {
				w.events.NodeMessageCreated(ctx, ev.Message.From, ev.Message.ID)
			}
		}
	}

	if ev.Channel != nil {
		_ = w.registry.UpsertChannel(ev.Channel)
	}
}
