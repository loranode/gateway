package worker

import (
	"time"

	pb "github.com/loranode/gateway/api/meshtastic"
	"github.com/loranode/gateway/internal/models"
)

// textMessage builds a Message from a text packet.
func (*Worker) textMessage(p *pb.MeshPacket, data *pb.Data) models.Message {
	rx := time.Unix(int64(p.GetRxTime()), 0).UTC()

	msg := models.Message{
		ID:      p.GetId(),
		From:    p.GetFrom(),
		To:      p.GetTo(),
		Channel: p.GetChannel(),
		Text:    string(data.GetPayload()),
		Snr:     p.GetRxSnr(),
		Rssi:    p.GetRxRssi(),
		RxTime:  rx,
	}

	if hs := p.GetHopStart(); hs > 0 && hs >= p.GetHopLimit() {
		hops := hs - p.GetHopLimit()
		msg.HopsAway = &hops
	}

	return msg
}
