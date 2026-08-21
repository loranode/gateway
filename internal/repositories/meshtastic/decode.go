package meshtastic

import (
	"github.com/loranode/meshtastic/pb/base"
	"github.com/loranode/meshtastic/pb/mesh"

	"github.com/loranode/gateway/internal/models"
)

// decode turns one FromRadio into a MeshEvent, or an empty event for variants we
// do not track.
func (r *Repository) decode(fr *mesh.FromRadio) models.MeshEvent {
	switch {
	case fr.GetNodeInfo() != nil:
		node := r.nodeInfo(fr.GetNodeInfo())

		return models.MeshEvent{Node: &node}
	case fr.GetPacket() != nil:
		return r.decodePacket(fr.GetPacket())
	case fr.GetChannel() != nil:
		ch := fr.GetChannel()

		return models.MeshEvent{Channel: &models.Channel{
			Index: uint32(ch.GetIndex()),
			Name:  ch.GetSettings().GetName(),
			Role:  ch.GetRole().String(),
		}}
	default:
		return models.MeshEvent{}
	}
}

// decodePacket builds a Node from a packet's link stats and, for a text payload,
// the Message it carries.
func (r *Repository) decodePacket(p *mesh.MeshPacket) models.MeshEvent {
	node := r.packetNode(p)
	ev := models.MeshEvent{Node: &node}

	data := p.GetDecoded()
	if data != nil && data.GetPortnum() == base.PortNum_TEXT_MESSAGE_APP {
		msg := r.textMessage(p, data)
		ev.Message = &msg
	}

	return ev
}
