package worker

import "github.com/loranode/gateway/internal/models"

// apply folds one decoded event into the registry and emits the matching events.
func (w *Worker) apply(ev models.MeshEvent) {
	if ev.Node != nil {
		if created, err := w.registry.UpsertNode(ev.Node); err == nil {
			if created {
				w.events.NodeCreated(ev.Node.Num)
			} else {
				w.events.NodeUpdated(ev.Node.Num)
			}
		}
	}

	if ev.Message != nil {
		if inserted, err := w.registry.AddMessage(ev.Message); err == nil && inserted {
			if ev.Message.To == broadcast {
				w.events.ChannelMessageCreated(ev.Message.From, ev.Message.ID, ev.Message.Channel)
			} else {
				w.events.NodeMessageCreated(ev.Message.From, ev.Message.ID)
			}
		}
	}

	if ev.Channel != nil {
		_ = w.registry.UpsertChannel(ev.Channel)
	}
}
