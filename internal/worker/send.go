package worker

import (
	"fmt"
	"math/rand/v2"

	"google.golang.org/protobuf/proto"

	pb "github.com/loranode/gateway/api/meshtastic"
)

// Send transmits a text message into the mesh. A zero recipient broadcasts.
func (w *Worker) Send(text string, to, channel uint32) error {
	if to == 0 {
		to = broadcast
	}

	frame, err := proto.Marshal(&pb.ToRadio{
		PayloadVariant: &pb.ToRadio_Packet{Packet: &pb.MeshPacket{
			To:      to,
			Channel: channel,
			Id:      rand.Uint32(),
			WantAck: true,
			PayloadVariant: &pb.MeshPacket_Decoded{Decoded: &pb.Data{
				Portnum: pb.PortNum_TEXT_MESSAGE_APP,
				Payload: []byte(text),
			}},
		}},
	})
	if err != nil {
		return fmt.Errorf("marshal send: %w", err)
	}

	if err := w.transport.WriteFrame(frame); err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	return nil
}
