package sqlite

import (
	"context"
	"fmt"

	"github.com/loranode/gateway/internal/models"
)

// AddMessage stores a received text message, returning whether it was inserted
// (false for an ignored retransmission). History is unbounded.
func (r *Repository) AddMessage(m *models.Message) (bool, error) {
	const query = `
INSERT OR IGNORE INTO messages (id, from_num, to_num, channel, text, snr, rssi, hops_away, rx_time)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`

	var hopsAway any
	if m.HopsAway != nil {
		hopsAway = int64(*m.HopsAway)
	}

	res, err := r.db.ExecContext(context.Background(), query,
		m.ID, m.From, m.To, m.Channel, m.Text, m.Snr, m.Rssi, hopsAway, m.RxTime.UTC().Unix(),
	)
	if err != nil {
		return false, fmt.Errorf("add message: %w", err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("add message rows: %w", err)
	}

	return n == 1, nil
}
