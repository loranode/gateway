package mappers

import (
	"github.com/loranode/gateway/api"
	"github.com/loranode/gateway/internal/models"
)

// NodeSummaryToProto converts a compact node projection into its REST summary.
func NodeSummaryToProto(n *models.NodeInfo) *api.NodeSummary {
	return &api.NodeSummary{
		Num:       n.Num,
		ShortName: n.ShortName,
		LongName:  n.LongName,
	}
}
