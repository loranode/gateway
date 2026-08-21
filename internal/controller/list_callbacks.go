package controller

import (
	"context"
	"fmt"

	"github.com/loranode/gateway/api"
)

// ListCallbacks returns every registered webhook subscription.
func (c *Controller) ListCallbacks(ctx context.Context, _ *api.ListCallbacksRequest) (*api.ListCallbacksResponse, error) {
	urls, err := c.events.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list callbacks: %w", err)
	}

	return &api.ListCallbacksResponse{Urls: urls}, nil
}
