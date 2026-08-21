package meshtastic

import (
	"time"

	"github.com/loranode/meshtastic/pb/mesh"

	"github.com/loranode/gateway/internal/models"
)

// textMessage builds a Message from a text packet.
func (*Repository) textMessage(p *mesh.MeshPacket, data *mesh.Data) models.Message {
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
