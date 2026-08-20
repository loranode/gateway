package mappers

import (
	"github.com/loranode/gateway/api/rest"
	"github.com/loranode/gateway/internal/models"
)

// ChannelToProto converts a stored channel into its REST representation.
func ChannelToProto(c models.Channel) *rest.Channel {
	return &rest.Channel{
		Index: c.Index,
		Name:  c.Name,
		Role:  c.Role,
	}
}
