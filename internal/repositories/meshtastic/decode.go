package meshtastic

import (
	"github.com/loranode/meshtastic/pb/base"
	"github.com/loranode/meshtastic/pb/mesh"
	"google.golang.org/protobuf/proto"

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
	case fr.GetConfig().GetLora() != nil:
		// The node's config dump carries its configured hop limit; adopt it so
		// sent packets travel as far as the node itself would send them.
		if hl := fr.GetConfig().GetLora().GetHopLimit(); hl > 0 {
			r.hopLimit.Store(hl)
		}

		return models.MeshEvent{}
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
	if data.GetPortnum() == base.PortNum_TEXT_MESSAGE_APP {
		msg := r.textMessage(p, data)
		ev.Message = &msg
	}

	// A routing packet carrying a request_id is the ack for one of our sends;
	// hand its result to the Send blocked on that packet id, if any.
	if data.GetPortnum() == base.PortNum_ROUTING_APP {
		if ch, ok := r.pending.Load(data.GetRequestId()); ok {
			if c, isChan := ch.(chan mesh.Routing_Error); isChan {
				var routing mesh.Routing

				_ = proto.Unmarshal(data.GetPayload(), &routing)

				select {
				case c <- routing.GetErrorReason():
				default:
				}
			}
		}
	}

	return ev
}
