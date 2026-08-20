package sqlite

import (
	"context"
	"fmt"

	"github.com/loranode/gateway/internal/models"
)

// Nodes returns the compact info of every known node, most recently heard first.
func (r *Repository) Nodes() ([]models.NodeInfo, error) {
	const query = `
SELECT num, short_name, long_name
FROM nodes
ORDER BY COALESCE(last_heard, updated_at) DESC;`

	rows, err := r.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("query nodes: %w", err)
	}

	defer func() { _ = rows.Close() }()

	var out []models.NodeInfo

	for rows.Next() {
		var n models.NodeInfo

		if err := rows.Scan(&n.Num, &n.ShortName, &n.LongName); err != nil {
			return nil, fmt.Errorf("scan node: %w", err)
		}

		out = append(out, n)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate nodes: %w", err)
	}

	return out, nil
}
