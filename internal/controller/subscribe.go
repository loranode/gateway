package controller

import (
	"fmt"

	"google.golang.org/grpc"

	"github.com/loranode/gateway/api"
	"github.com/loranode/gateway/internal/mappers"
)

// Subscribe streams mesh events to a gRPC client until it disconnects.
func (c *Controller) Subscribe(_ *api.SubscribeRequest, stream grpc.ServerStreamingServer[api.Event]) error {
	ch, unsub := c.events.Subscribe()
	defer unsub()

	ctx := stream.Context()

	for {
		select {
		case <-ctx.Done():
			return nil
		case ev := <-ch:
			if err := stream.Send(mappers.EventToProto(ev)); err != nil {
				return fmt.Errorf("send event: %w", err)
			}
		}
	}
}
