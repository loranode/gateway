package worker

import (
	"time"

	pb "github.com/loranode/gateway/api/meshtastic"
	"github.com/loranode/gateway/internal/models"
)

// nodeInfo builds a node record from a node-database entry.
func (w *Worker) nodeInfo(ni *pb.NodeInfo) models.Node {
	node := models.Node{Num: ni.GetNum()}

	if u := ni.GetUser(); u != nil {
		w.applyUser(&node, u)
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

	w.applyPosition(&node, ni.GetPosition())
	w.applyMetrics(&node, ni.GetDeviceMetrics())

	return node
}
