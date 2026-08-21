package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/merzzzl/proto-rest-api/runtime"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/loranode/gateway/api"
)

// SendChannelMessage broadcasts a text message on one channel.
func (c *Controller) SendChannelMessage(ctx context.Context, req *api.SendChannelMessageRequest) (*emptypb.Empty, error) {
	if req.GetText() == "" {
		return nil, runtime.Errorf(http.StatusBadRequest, "text is required")
	}

	if err := c.mesh.Send(ctx, req.GetText(), 0, req.GetIndex(), req.GetReplyId()); err != nil {
		return nil, fmt.Errorf("send channel message: %w", err)
	}

	return &emptypb.Empty{}, nil
}
