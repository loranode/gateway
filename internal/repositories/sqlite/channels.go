package sqlite

import (
	"context"
	"fmt"

	"github.com/loranode/gateway/internal/models"
)

// Channels returns every configured channel, ordered by index.
func (r *Repository) Channels() ([]models.Channel, error) {
	const query = `
SELECT idx, name, role
FROM channels
ORDER BY idx;`

	rows, err := r.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("query channels: %w", err)
	}

	defer func() { _ = rows.Close() }()

	var out []models.Channel

	for rows.Next() {
		var c models.Channel

		if err := rows.Scan(&c.Index, &c.Name, &c.Role); err != nil {
			return nil, fmt.Errorf("scan channel: %w", err)
		}

		out = append(out, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate channels: %w", err)
	}

	return out, nil
}
