package sqlite

import (
	"context"
	"fmt"
)

// DeleteCallback unregisters a webhook URL.
func (r *Repository) DeleteCallback(ctx context.Context, url string) error {
	const query = `DELETE FROM callbacks WHERE url = ?;`

	if _, err := r.db.ExecContext(ctx, query, url); err != nil {
		return fmt.Errorf("delete callback: %w", err)
	}

	return nil
}
