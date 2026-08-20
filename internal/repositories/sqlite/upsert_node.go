package sqlite

import (
	"context"
	"fmt"
	"time"

	"github.com/loranode/gateway/internal/models"
)

// UpsertNode merges a partial node record into the store, returning whether the
// node was newly created (vs updated).
func (r *Repository) UpsertNode(patch *models.Node) (bool, error) {
	// On conflict each column keeps its existing value unless the patch carries a
	// meaningful one: NULLIF treats empty strings and zero numbers as "unset";
	// nullable columns are unset when NULL.
	const query = `
INSERT INTO nodes (
	num, long_name, short_name, hw_model, role, snr, rssi,
	hops_away, latitude, longitude, altitude, battery, voltage,
	via_mqtt, is_favorite, last_heard, updated_at
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
ON CONFLICT(num) DO UPDATE SET
	long_name   = COALESCE(NULLIF(excluded.long_name, ''), nodes.long_name),
	short_name  = COALESCE(NULLIF(excluded.short_name, ''), nodes.short_name),
	hw_model    = COALESCE(NULLIF(excluded.hw_model, ''), nodes.hw_model),
	role        = COALESCE(NULLIF(excluded.role, ''), nodes.role),
	snr         = COALESCE(NULLIF(excluded.snr, 0), nodes.snr),
	rssi        = COALESCE(NULLIF(excluded.rssi, 0), nodes.rssi),
	hops_away   = COALESCE(excluded.hops_away, nodes.hops_away),
	latitude    = COALESCE(excluded.latitude, nodes.latitude),
	longitude   = COALESCE(excluded.longitude, nodes.longitude),
	altitude    = COALESCE(excluded.altitude, nodes.altitude),
	battery     = COALESCE(excluded.battery, nodes.battery),
	voltage     = COALESCE(excluded.voltage, nodes.voltage),
	via_mqtt    = MAX(nodes.via_mqtt, excluded.via_mqtt),
	is_favorite = MAX(nodes.is_favorite, excluded.is_favorite),
	last_heard  = COALESCE(excluded.last_heard, nodes.last_heard),
	updated_at  = excluded.updated_at;`

	var hopsAway, latitude, longitude, altitude, battery, voltage, lastHeard any

	if patch.HopsAway != nil {
		hopsAway = int64(*patch.HopsAway)
	}

	if patch.Latitude != nil {
		latitude = *patch.Latitude
	}

	if patch.Longitude != nil {
		longitude = *patch.Longitude
	}

	if patch.Altitude != nil {
		altitude = int64(*patch.Altitude)
	}

	if patch.Battery != nil {
		battery = int64(*patch.Battery)
	}

	if patch.Voltage != nil {
		voltage = float64(*patch.Voltage)
	}

	if patch.LastHeard != nil {
		lastHeard = patch.LastHeard.Unix()
	}

	viaMqtt := 0
	if patch.ViaMqtt {
		viaMqtt = 1
	}

	isFavorite := 0
	if patch.IsFavorite {
		isFavorite = 1
	}

	var exists bool

	if err := r.db.QueryRowContext(context.Background(),
		`SELECT EXISTS(SELECT 1 FROM nodes WHERE num = ?);`, patch.Num).Scan(&exists); err != nil {
		return false, fmt.Errorf("check node: %w", err)
	}

	if _, err := r.db.ExecContext(context.Background(), query,
		patch.Num, patch.LongName, patch.ShortName, patch.HwModel, patch.Role,
		patch.Snr, patch.Rssi, hopsAway, latitude, longitude, altitude, battery, voltage,
		viaMqtt, isFavorite, lastHeard, time.Now().UTC().Unix(),
	); err != nil {
		return false, fmt.Errorf("upsert node: %w", err)
	}

	return !exists, nil
}
