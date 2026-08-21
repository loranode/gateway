package meshtastic

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.com/loranode/meshtastic/pb/base"
	"github.com/loranode/meshtastic/pb/mesh"
	"google.golang.org/protobuf/proto"
)

// Send transmits a text message into the mesh, waiting through a reconnect. A
// zero recipient broadcasts; a non-zero replyID marks a reply to that packet id.
func (r *Repository) Send(ctx context.Context, text string, to, channel, replyID uint32) error {
	if to == 0 {
		to = broadcast
	}

	frame, err := proto.Marshal(&mesh.ToRadio{
		PayloadVariant: &mesh.ToRadio_Packet{Packet: &mesh.MeshPacket{
			To:       to,
			Channel:  channel,
			Id:       rand.Uint32(),
			HopLimit: r.hopLimit.Load(),
			WantAck:  true,
			PayloadVariant: &mesh.MeshPacket_Decoded{Decoded: &mesh.Data{
				Portnum: base.PortNum_TEXT_MESSAGE_APP,
				Payload: []byte(text),
				ReplyId: replyID,
			}},
		}},
	})
	if err != nil {
		return fmt.Errorf("marshal send: %w", err)
	}

	if err := r.session.SendRaw(ctx, frame); err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	return nil
}
