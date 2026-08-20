package controller

import (
	"context"
	"fmt"

	"github.com/loranode/gateway/api/rest"
	"github.com/loranode/gateway/internal/mappers"
)

// ListNodes returns a compact summary of every node the connected radio sees.
func (c *Controller) ListNodes(_ context.Context, _ *rest.ListNodesRequest) (*rest.ListNodesResponse, error) {
	nodes, err := c.registry.Nodes()
	if err != nil {
		return nil, fmt.Errorf("list nodes: %w", err)
	}

	out := make([]*rest.NodeSummary, 0, len(nodes))
	for i := range nodes {
		out = append(out, mappers.NodeSummaryToProto(&nodes[i]))
	}

	return &rest.ListNodesResponse{Nodes: out}, nil
}
