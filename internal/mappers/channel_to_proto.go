package mappers

import (
	"github.com/loranode/gateway/api"
	"github.com/loranode/gateway/internal/models"
)

// ChannelToProto converts a stored channel into its REST representation.
func ChannelToProto(c models.Channel) *api.Channel {
	return &api.Channel{
		Index: c.Index,
		Name:  c.Name,
		Role:  c.Role,
	}
}
