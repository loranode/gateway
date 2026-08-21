package controller

import (
	"context"
	"fmt"

	"github.com/loranode/gateway/api"
	"github.com/loranode/gateway/internal/mappers"
)

// ListChannels returns every configured channel on the connected radio.
func (c *Controller) ListChannels(_ context.Context, _ *api.ListChannelsRequest) (*api.ListChannelsResponse, error) {
	channels, err := c.registry.Channels()
	if err != nil {
		return nil, fmt.Errorf("list channels: %w", err)
	}

	out := make([]*api.Channel, 0, len(channels))
	for i := range channels {
		out = append(out, mappers.ChannelToProto(channels[i]))
	}

	return &api.ListChannelsResponse{Channels: out}, nil
}
