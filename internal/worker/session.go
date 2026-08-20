package worker

import (
	"context"
	"fmt"
)

// session handles one connection lifetime: it asks for the node database, starts
// the heartbeat, and applies every frame until the link drops.
func (w *Worker) session(ctx context.Context) error {
	if err := w.transport.Connect(ctx); err != nil {
		return fmt.Errorf("connect: %w", err)
	}

	// Close the connection on cancellation to unblock the blocking ReadFrame.
	stop := context.AfterFunc(ctx, func() { _ = w.transport.Close() })
	defer stop()
	defer func() { _ = w.transport.Close() }()

	frame, err := w.wantConfig()
	if err != nil {
		return fmt.Errorf("want_config: %w", err)
	}

	if err := w.transport.WriteFrame(frame); err != nil {
		return fmt.Errorf("send want_config: %w", err)
	}

	hbCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go w.heartbeat(hbCtx)

	for {
		payload, err := w.transport.ReadFrame()
		if err != nil {
			return fmt.Errorf("read frame: %w", err)
		}

		if payload == nil {
			continue
		}

		w.apply(ctx, payload)
	}
}
