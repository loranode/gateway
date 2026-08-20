package controller

import (
	"context"
	"fmt"

	"github.com/loranode/gateway/api/rest"
	"github.com/loranode/gateway/internal/mappers"
)

// ListChannels returns every configured channel on the connected radio.
func (c *Controller) ListChannels(_ context.Context, _ *rest.ListChannelsRequest) (*rest.ListChannelsResponse, error) {
	channels, err := c.registry.Channels()
	if err != nil {
		return nil, fmt.Errorf("list channels: %w", err)
	}

	out := make([]*rest.Channel, 0, len(channels))
	for i := range channels {
		out = append(out, mappers.ChannelToProto(channels[i]))
	}

	return &rest.ListChannelsResponse{Channels: out}, nil
}
