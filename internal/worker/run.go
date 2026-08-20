package worker

import (
	"context"
	"time"
)

// Run connects to the node and processes its stream until ctx is cancelled,
// reconnecting after any error.
func (w *Worker) Run(ctx context.Context) {
	for {
		_ = w.session(ctx)

		select {
		case <-ctx.Done():
			return
		case <-time.After(reconnectDelay):
		}
	}
}
