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

	// The firmware only PKI-encrypts a client-API direct message when the packet
	// asks for it; without the flag it falls back to channel encryption, which a
	// recipient expecting a PKI DM silently drops. Request it for direct messages
	// to nodes we have seen advertise a public key.
	_, hasKey := r.pkiNodes.Load(to)

	frame, err := proto.Marshal(&mesh.ToRadio{
		PayloadVariant: &mesh.ToRadio_Packet{Packet: &mesh.MeshPacket{
			To:           to,
			Channel:      channel,
			Id:           rand.Uint32(),
			HopLimit:     r.hopLimit.Load(),
			WantAck:      true,
			PkiEncrypted: to != broadcast && hasKey,
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
