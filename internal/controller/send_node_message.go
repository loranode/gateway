package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/merzzzl/proto-rest-api/runtime"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/loranode/gateway/api"
)

// SendNodeMessage transmits a text message to one node.
func (c *Controller) SendNodeMessage(ctx context.Context, req *api.SendMessageRequest) (*emptypb.Empty, error) {
	if req.GetText() == "" {
		return nil, runtime.Errorf(http.StatusBadRequest, "text is required")
	}

	if err := c.mesh.Send(ctx, req.GetText(), req.GetNum(), req.GetChannel(), req.GetReplyId()); err != nil {
		return nil, fmt.Errorf("send node message: %w", err)
	}

	return &emptypb.Empty{}, nil
}
