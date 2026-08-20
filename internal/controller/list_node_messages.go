package controller

import (
	"context"
	"fmt"

	"github.com/loranode/gateway/api/rest"
	"github.com/loranode/gateway/internal/mappers"
)

// ListNodeMessages returns direct messages received from one node.
func (c *Controller) ListNodeMessages(_ context.Context, req *rest.ListMessagesRequest) (*rest.ListMessagesResponse, error) {
	messages, err := c.registry.MessagesByNode(req.GetNum())
	if err != nil {
		return nil, fmt.Errorf("list node messages: %w", err)
	}

	out := make([]*rest.Message, 0, len(messages))
	for i := range messages {
		out = append(out, mappers.MessageToProto(&messages[i]))
	}

	return &rest.ListMessagesResponse{Messages: out}, nil
}
