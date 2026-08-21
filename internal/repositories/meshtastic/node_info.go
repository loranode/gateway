package meshtastic

import (
	"time"

	"github.com/loranode/meshtastic/pb/mesh"

	"github.com/loranode/gateway/internal/models"
)

// nodeInfo builds a node record from a node-database entry.
func (r *Repository) nodeInfo(ni *mesh.NodeInfo) models.Node {
	node := models.Node{Num: ni.GetNum()}

	if u := ni.GetUser(); u != nil {
		r.applyUser(&node, u)
	}

	if ni.GetSnr() != 0 {
		node.Snr = ni.GetSnr()
	}

	if lh := ni.GetLastHeard(); lh > 0 {
		t := time.Unix(int64(lh), 0).UTC()
		node.LastHeard = &t
	}

	if ni.HopsAway != nil {
		hops := ni.GetHopsAway()
		node.HopsAway = &hops
	}

	node.ViaMqtt = ni.GetViaMqtt()
	node.IsFavorite = ni.GetIsFavorite()

	r.applyPosition(&node, ni.GetPosition())
	r.applyMetrics(&node, ni.GetDeviceMetrics())

	return node
}
