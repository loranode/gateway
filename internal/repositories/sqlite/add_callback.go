package sqlite

import (
	"context"
	"fmt"
	"time"
)

// AddCallback registers a webhook URL; a repeated URL is ignored.
func (r *Repository) AddCallback(ctx context.Context, url string) error {
	const query = `INSERT OR IGNORE INTO callbacks (url, created_at) VALUES (?, ?);`

	if _, err := r.db.ExecContext(ctx, query, url, time.Now().UTC().Unix()); err != nil {
		return fmt.Errorf("add callback: %w", err)
	}

	return nil
}
