package meshtastic

import (
	"time"

	"github.com/loranode/meshtastic/pb/base"
	"github.com/loranode/meshtastic/pb/mesh"
	"google.golang.org/protobuf/proto"

	"github.com/loranode/gateway/internal/models"
)

// packetNode builds a node record from a mesh packet's link stats and, for the
// payloads we care about, its decoded position, user or telemetry.
func (r *Repository) packetNode(p *mesh.MeshPacket) models.Node {
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
	case base.PortNum_POSITION_APP:
		var pos mesh.Position
		if proto.Unmarshal(data.GetPayload(), &pos) == nil {
			r.applyPosition(&node, &pos)
		}
	case base.PortNum_NODEINFO_APP:
		var u mesh.User
		if proto.Unmarshal(data.GetPayload(), &u) == nil {
			r.applyUser(&node, &u)
		}
	case base.PortNum_TELEMETRY_APP:
		var tel base.Telemetry
		if proto.Unmarshal(data.GetPayload(), &tel) == nil {
			r.applyMetrics(&node, tel.GetDeviceMetrics())
		}
	default:
		// Nothing beyond the link stats already captured above.
	}

	return node
}
