package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/loranode/gateway/internal/models"
)

// MessagesByChannel returns the messages received on one channel, newest first.
func (r *Repository) MessagesByChannel(channel uint32) ([]models.Message, error) {
	const limit = 1000

	const query = `
SELECT id, from_num, to_num, channel, text, snr, rssi, hops_away, rx_time
FROM messages
WHERE channel = ?
ORDER BY rx_time DESC, id DESC
LIMIT ?;`

	rows, err := r.db.QueryContext(context.Background(), query, channel, limit)
	if err != nil {
		return nil, fmt.Errorf("query channel messages: %w", err)
	}

	defer func() { _ = rows.Close() }()

	var out []models.Message

	for rows.Next() {
		var (
			m        models.Message
			hopsAway sql.NullInt64
			rxTime   int64
		)

		if err := rows.Scan(
			&m.ID, &m.From, &m.To, &m.Channel, &m.Text, &m.Snr, &m.Rssi, &hopsAway, &rxTime,
		); err != nil {
			return nil, fmt.Errorf("scan message: %w", err)
		}

		if hopsAway.Valid {
			h := uint32(hopsAway.Int64)
			m.HopsAway = &h
		}

		m.RxTime = time.Unix(rxTime, 0).UTC()

		out = append(out, m)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate channel messages: %w", err)
	}

	return out, nil
}
