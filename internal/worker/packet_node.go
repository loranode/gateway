package worker

import (
	"time"

	"google.golang.org/protobuf/proto"

	pb "github.com/loranode/gateway/api/meshtastic"
	"github.com/loranode/gateway/internal/models"
)

// packetNode builds a node record from a mesh packet's link stats and, for the
// payloads we care about, its decoded position, user or telemetry.
func (w *Worker) packetNode(p *pb.MeshPacket) models.Node {
	node := models.Node{Num: p.GetFrom()}

	if p.GetRxSnr() != 0 {
		node.Snr = p.GetRxSnr()
	}

	if p.GetRxRssi() != 0 {
		node.Rssi = p.GetRxRssi()
	}

	if rt := p.GetRxTime(); rt > 0 {
		t := time.Unix(int64(rt), 0).UTC()
		node.LastHeard = &t
	}

	if hs := p.GetHopStart(); hs > 0 && hs >= p.GetHopLimit() {
		hops := hs - p.GetHopLimit()
		node.HopsAway = &hops
	}

	data := p.GetDecoded()
	if data == nil {
		return node
	}

	//nolint:exhaustive // only a few port numbers carry node state; the rest are link stats only
	switch data.GetPortnum() {
	case pb.PortNum_POSITION_APP:
		var pos pb.Position
		if proto.Unmarshal(data.GetPayload(), &pos) == nil {
			w.applyPosition(&node, &pos)
		}
	case pb.PortNum_NODEINFO_APP:
		var u pb.User
		if proto.Unmarshal(data.GetPayload(), &u) == nil {
			w.applyUser(&node, &u)
		}
	case pb.PortNum_TELEMETRY_APP:
		var tel pb.Telemetry
		if proto.Unmarshal(data.GetPayload(), &tel) == nil {
			w.applyMetrics(&node, tel.GetDeviceMetrics())
		}
	default:
		// Nothing beyond the link stats already captured above.
	}

	return node
}
