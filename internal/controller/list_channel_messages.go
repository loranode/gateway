package controller

import (
	"context"
	"fmt"

	"github.com/loranode/gateway/api/rest"
	"github.com/loranode/gateway/internal/mappers"
)

// ListChannelMessages returns messages received on one channel.
func (c *Controller) ListChannelMessages(_ context.Context, req *rest.ListChannelMessagesRequest) (*rest.ListMessagesResponse, error) {
	messages, err := c.registry.MessagesByChannel(req.GetIndex())
	if err != nil {
		return nil, fmt.Errorf("list channel messages: %w", err)
	}

	out := make([]*rest.Message, 0, len(messages))
	for i := range messages {
		out = append(out, mappers.MessageToProto(&messages[i]))
	}

	return &rest.ListMessagesResponse{Messages: out}, nil
}
