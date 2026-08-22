package meshtastic

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"

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
	id := rand.Uint32()

	frame, err := proto.Marshal(&mesh.ToRadio{
		PayloadVariant: &mesh.ToRadio_Packet{Packet: &mesh.MeshPacket{
			To:           to,
			Channel:      channel,
			Id:           id,
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

	// A broadcast has only an unreliable implicit ack, so send it and return. A
	// direct message registers a waiter before sending and blocks until the read
	// loop reports the node's routing ack, so the caller learns whether the mesh
	// accepted it.
	if to == broadcast {
		if err := r.session.SendRaw(ctx, frame); err != nil {
			return fmt.Errorf("send message: %w", err)
		}

		return nil
	}

	ack := make(chan mesh.Routing_Error, 1)
	r.pending.Store(id, ack)

	defer r.pending.Delete(id)

	if err := r.session.SendRaw(ctx, frame); err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	select {
	case reason := <-ack:
		if reason != mesh.Routing_NONE {
			return fmt.Errorf("%w: %s", ErrSendRejected, reason)
		}

		return nil
	case <-time.After(ackTimeout):
		return ErrSendNotAcked
	case <-ctx.Done():
		return fmt.Errorf("await ack: %w", ctx.Err())
	}
}
