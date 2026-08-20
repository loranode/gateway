package mappers

import (
	"github.com/loranode/gateway/api/rest"
	"github.com/loranode/gateway/internal/models"
)

// MessageToProto converts a stored message into its REST representation. The API
// id is the on-air Meshtastic packet id. RxTime is Unix seconds.
func MessageToProto(m *models.Message) *rest.Message {
	return &rest.Message{
		Id:       m.ID,
		From:     m.From,
		To:       m.To,
		Channel:  m.Channel,
		Text:     m.Text,
		Snr:      m.Snr,
		Rssi:     m.Rssi,
		HopsAway: m.HopsAway,
		RxTime:   m.RxTime.Unix(),
	}
}
