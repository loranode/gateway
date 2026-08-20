package controller

import (
	"context"
	"fmt"

	"github.com/loranode/gateway/api/rest"
)

// ListCallbacks returns every registered webhook subscription.
func (c *Controller) ListCallbacks(ctx context.Context, _ *rest.ListCallbacksRequest) (*rest.ListCallbacksResponse, error) {
	urls, err := c.events.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list callbacks: %w", err)
	}

	return &rest.ListCallbacksResponse{Urls: urls}, nil
}
