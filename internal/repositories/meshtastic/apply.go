package meshtastic

import (
	"github.com/loranode/meshtastic/pb/base"
	"github.com/loranode/meshtastic/pb/mesh"

	"github.com/loranode/gateway/internal/models"
)

// applyUser copies identity fields from a User onto the node.
func (*Repository) applyUser(n *models.Node, u *mesh.User) {
	n.LongName = u.GetLongName()
	n.ShortName = u.GetShortName()
	n.HwModel = u.GetHwModel().String()
	n.Role = u.GetRole().String()
}

// applyPosition copies a valid (non-zero) position onto the node.
func (*Repository) applyPosition(n *models.Node, pos *mesh.Position) {
	// posScale converts Meshtastic's fixed-point lat/lon (1e-7 degrees) to degrees.
	const posScale = 1e7

	if pos == nil {
		return
	}

	lat, lon := pos.GetLatitudeI(), pos.GetLongitudeI()
	if lat == 0 && lon == 0 {
		return
	}

	latf := float64(lat) / posScale
	lonf := float64(lon) / posScale
	n.Latitude = &latf
	n.Longitude = &lonf

	if alt := pos.GetAltitude(); alt != 0 {
		a := alt
		n.Altitude = &a
	}
}

// applyMetrics copies battery and voltage telemetry onto the node.
func (*Repository) applyMetrics(n *models.Node, m *base.DeviceMetrics) {
	if m == nil {
		return
	}

	if b := m.GetBatteryLevel(); b > 0 {
		bl := b
		n.Battery = &bl
	}

	if v := m.GetVoltage(); v != 0 {
		vl := v
		n.Voltage = &vl
	}
}
