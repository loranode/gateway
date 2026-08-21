package controller

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/loranode/gateway/api"
)

// RegisterCallback subscribes a webhook URL to mesh events.
func (c *Controller) RegisterCallback(ctx context.Context, req *api.RegisterCallbackRequest) (*emptypb.Empty, error) {
	if err := c.events.Register(ctx, req.GetUrl()); err != nil {
		return nil, fmt.Errorf("register callback: %w", err)
	}

	return &emptypb.Empty{}, nil
}
