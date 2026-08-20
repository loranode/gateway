package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/merzzzl/proto-rest-api/runtime"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/loranode/gateway/api/rest"
)

// SendNodeMessage transmits a text message to one node.
func (c *Controller) SendNodeMessage(_ context.Context, req *rest.SendMessageRequest) (*emptypb.Empty, error) {
	if req.GetText() == "" {
		return nil, runtime.Errorf(http.StatusBadRequest, "text is required")
	}

	if err := c.send(req.GetText(), req.GetNum(), req.GetChannel()); err != nil {
		return nil, fmt.Errorf("send node message: %w", err)
	}

	return &emptypb.Empty{}, nil
}
