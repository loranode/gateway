package mappers

import (
	"github.com/loranode/gateway/api/rest"
	"github.com/loranode/gateway/internal/models"
)

// NodeSummaryToProto converts a compact node projection into its REST summary.
func NodeSummaryToProto(n *models.NodeInfo) *rest.NodeSummary {
	return &rest.NodeSummary{
		Num:       n.Num,
		ShortName: n.ShortName,
		LongName:  n.LongName,
	}
}
