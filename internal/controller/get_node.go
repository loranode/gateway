package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/merzzzl/proto-rest-api/runtime"

	"github.com/loranode/gateway/api"
	"github.com/loranode/gateway/internal/mappers"
)

// GetNode returns all stored information for one node by its number, or 404.
func (c *Controller) GetNode(_ context.Context, req *api.GetNodeRequest) (*api.Node, error) {
	node, ok, err := c.registry.NodeByNum(req.GetNum())
	if err != nil {
		return nil, fmt.Errorf("get node: %w", err)
	}

	if !ok {
		return nil, runtime.Errorf(http.StatusNotFound, "node %d not found", req.GetNum())
	}

	return mappers.NodeToProto(&node), nil
}
