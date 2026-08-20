package worker

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/protobuf/proto"

	pb "github.com/loranode/gateway/api/meshtastic"
)

// heartbeat periodically sends the heartbeat frame so the node does not drop us.
func (w *Worker) heartbeat(ctx context.Context) {
	frame, err := w.heartbeatFrame()
	if err != nil {
		return
	}

	t := time.NewTicker(heartbeatEvery)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			if err := w.transport.WriteFrame(frame); err != nil {
				return
			}
		}
	}
}

// heartbeatFrame builds the ToRadio heartbeat frame.
func (*Worker) heartbeatFrame() ([]byte, error) {
	b, err := proto.Marshal(&pb.ToRadio{
		PayloadVariant: &pb.ToRadio_Heartbeat{Heartbeat: &pb.Heartbeat{}},
	})
	if err != nil {
		return nil, fmt.Errorf("marshal heartbeat: %w", err)
	}

	return b, nil
}
