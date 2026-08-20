package mappers

import (
	"github.com/loranode/gateway/api/rest"
	"github.com/loranode/gateway/internal/models"
)

// NodeToProto converts a stored node into its REST representation. Timestamps are
// Unix seconds.
func NodeToProto(n *models.Node) *rest.Node {
	out := &rest.Node{
		Num:        n.Num,
		LongName:   n.LongName,
		ShortName:  n.ShortName,
		HwModel:    n.HwModel,
		Role:       n.Role,
		Snr:        n.Snr,
		Rssi:       n.Rssi,
		HopsAway:   n.HopsAway,
		Latitude:   n.Latitude,
		Longitude:  n.Longitude,
		Altitude:   n.Altitude,
		Battery:    n.Battery,
		Voltage:    n.Voltage,
		ViaMqtt:    n.ViaMqtt,
		IsFavorite: n.IsFavorite,
		UpdatedAt:  n.UpdatedAt.Unix(),
	}

	if n.LastHeard != nil {
		out.LastHeard = n.LastHeard.Unix()
	}

	return out
}
