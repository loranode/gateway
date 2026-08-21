package controller

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/loranode/gateway/api"
)

// DeleteCallback unsubscribes a webhook URL.
func (c *Controller) DeleteCallback(ctx context.Context, req *api.DeleteCallbackRequest) (*emptypb.Empty, error) {
	if err := c.events.Delete(ctx, req.GetUrl()); err != nil {
		return nil, fmt.Errorf("delete callback: %w", err)
	}

	return &emptypb.Empty{}, nil
}
