package sqlite

import (
	"context"
	"fmt"
	"time"

	"github.com/loranode/gateway/internal/models"
)

// UpsertChannel stores a channel configuration, overwriting the previous entry
// for the same index.
func (r *Repository) UpsertChannel(ch *models.Channel) error {
	const query = `
INSERT INTO channels (idx, name, role, updated_at)
VALUES (?, ?, ?, ?)
ON CONFLICT(idx) DO UPDATE SET
	name       = excluded.name,
	role       = excluded.role,
	updated_at = excluded.updated_at;`

	if _, err := r.db.ExecContext(context.Background(), query,
		ch.Index, ch.Name, ch.Role, time.Now().UTC().Unix(),
	); err != nil {
		return fmt.Errorf("upsert channel: %w", err)
	}

	return nil
}
