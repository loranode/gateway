package worker

import (
	"context"

	"google.golang.org/protobuf/proto"

	pb "github.com/loranode/gateway/api/meshtastic"
	"github.com/loranode/gateway/internal/models"
)

// apply decodes one raw FromRadio frame and folds it into the registry.
func (w *Worker) apply(ctx context.Context, frame []byte) {
	var fr pb.FromRadio

	if proto.Unmarshal(frame, &fr) != nil {
		return
	}

	switch {
	case fr.GetNodeInfo() != nil:
		node := w.nodeInfo(fr.GetNodeInfo())

		created, err := w.registry.UpsertNode(&node)
		if err != nil {
			return
		}

		if created {
			w.events.NodeCreated(ctx, node.Num)
		} else {
			w.events.NodeUpdated(ctx, node.Num)
		}
	case fr.GetPacket() != nil:
		w.applyPacket(ctx, fr.GetPacket())
	case fr.GetChannel() != nil:
		ch := fr.GetChannel()

		channel := models.Channel{
			Index: uint32(ch.GetIndex()),
			Name:  ch.GetSettings().GetName(),
			Role:  ch.GetRole().String(),
		}

		if err := w.registry.UpsertChannel(&channel); err != nil {
			return
		}
	default:
		// Other FromRadio variants carry nothing we track.
	}
}

// applyPacket stores the link stats and any decoded payload of a mesh packet.
func (w *Worker) applyPacket(ctx context.Context, p *pb.MeshPacket) {
	node := w.packetNode(p)

	created, err := w.registry.UpsertNode(&node)
	if err != nil {
		return
	}

	if created {
		w.events.NodeCreated(ctx, node.Num)
	} else {
		w.events.NodeUpdated(ctx, node.Num)
	}

	data := p.GetDecoded()
	if data == nil || data.GetPortnum() != pb.PortNum_TEXT_MESSAGE_APP {
		return
	}

	msg := w.textMessage(p, data)

	inserted, err := w.registry.AddMessage(&msg)
	if err != nil {
		return
	}

	if inserted {
		if msg.To == broadcast {
			w.events.ChannelMessageCreated(ctx, msg.From, msg.ID, msg.Channel)
		} else {
			w.events.NodeMessageCreated(ctx, msg.From, msg.ID)
		}
	}
}

// applyUser copies identity fields from a User onto the node.
func (*Worker) applyUser(n *models.Node, u *pb.User) {
	n.LongName = u.GetLongName()
	n.ShortName = u.GetShortName()
	n.HwModel = u.GetHwModel().String()
	n.Role = u.GetRole().String()
}

// applyPosition copies a valid (non-zero) position onto the node.
func (*Worker) applyPosition(n *models.Node, pos *pb.Position) {
	// posScale converts Meshtastic's fixed-point lat/lon (1e-7 degrees) to degrees.
	const posScale = 1e7

	if pos == nil {
		return
	}

	lat, lon := pos.GetLatitudeI(), pos.GetLongitudeI()
	if lat == 0 && lon == 0 {
		return
	}

	latf := float64(lat) / posScale
	lonf := float64(lon) / posScale
	n.Latitude = &latf
	n.Longitude = &lonf

	if alt := pos.GetAltitude(); alt != 0 {
		a := alt
		n.Altitude = &a
	}
}

// applyMetrics copies battery and voltage telemetry onto the node.
func (*Worker) applyMetrics(n *models.Node, m *pb.DeviceMetrics) {
	if m == nil {
		return
	}

	if b := m.GetBatteryLevel(); b > 0 {
		bl := b
		n.Battery = &bl
	}

	if v := m.GetVoltage(); v != 0 {
		vl := v
		n.Voltage = &vl
	}
}
