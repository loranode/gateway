package sqlite

import (
	"context"
	"fmt"
)

// Callbacks returns every registered webhook URL.
func (r *Repository) Callbacks(ctx context.Context) ([]string, error) {
	const query = `SELECT url FROM callbacks ORDER BY created_at;`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query callbacks: %w", err)
	}

	defer func() { _ = rows.Close() }()

	var out []string

	for rows.Next() {
		var url string

		if err := rows.Scan(&url); err != nil {
			return nil, fmt.Errorf("scan callback: %w", err)
		}

		out = append(out, url)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate callbacks: %w", err)
	}

	return out, nil
}
