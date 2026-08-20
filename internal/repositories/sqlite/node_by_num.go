package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/loranode/gateway/internal/models"
)

// NodeByNum returns the full record of the node with the given number. The
// second result is false when no such node exists.
func (r *Repository) NodeByNum(num uint32) (models.Node, bool, error) {
	const query = `
SELECT num, long_name, short_name, hw_model, role, snr, rssi,
       hops_away, latitude, longitude, altitude, battery, voltage,
       via_mqtt, is_favorite, last_heard, updated_at
FROM nodes
WHERE num = ?
LIMIT 1;`

	rows, err := r.db.QueryContext(context.Background(), query, num)
	if err != nil {
		return models.Node{}, false, fmt.Errorf("query node: %w", err)
	}

	defer func() { _ = rows.Close() }()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return models.Node{}, false, fmt.Errorf("iterate node: %w", err)
		}

		return models.Node{}, false, nil
	}

	var (
		n          models.Node
		hopsAway   sql.NullInt64
		latitude   sql.NullFloat64
		longitude  sql.NullFloat64
		altitude   sql.NullInt64
		battery    sql.NullInt64
		voltage    sql.NullFloat64
		lastHeard  sql.NullInt64
		viaMqtt    int
		isFavorite int
		updatedAt  int64
	)

	if err := rows.Scan(
		&n.Num, &n.LongName, &n.ShortName, &n.HwModel, &n.Role, &n.Snr, &n.Rssi,
		&hopsAway, &latitude, &longitude, &altitude, &battery, &voltage,
		&viaMqtt, &isFavorite, &lastHeard, &updatedAt,
	); err != nil {
		return models.Node{}, false, fmt.Errorf("scan node: %w", err)
	}

	if hopsAway.Valid {
		h := uint32(hopsAway.Int64)
		n.HopsAway = &h
	}

	if latitude.Valid {
		v := latitude.Float64
		n.Latitude = &v
	}

	if longitude.Valid {
		v := longitude.Float64
		n.Longitude = &v
	}

	if altitude.Valid {
		v := int32(altitude.Int64)
		n.Altitude = &v
	}

	if battery.Valid {
		v := uint32(battery.Int64)
		n.Battery = &v
	}

	if voltage.Valid {
		v := float32(voltage.Float64)
		n.Voltage = &v
	}

	if lastHeard.Valid {
		t := time.Unix(lastHeard.Int64, 0).UTC()
		n.LastHeard = &t
	}

	n.ViaMqtt = viaMqtt != 0
	n.IsFavorite = isFavorite != 0
	n.UpdatedAt = time.Unix(updatedAt, 0).UTC()

	return n, true, nil
}
