package meshtastic

import (
	"context"
	"fmt"

	"github.com/loranode/meshtastic/pb/mesh"
	"google.golang.org/protobuf/proto"

	"github.com/loranode/gateway/internal/models"
)

// Read returns the next decoded event from the node, waiting through a reconnect
// if the link is down. An undecodable frame yields an empty event, not an error.
func (r *Repository) Read(ctx context.Context) (models.MeshEvent, error) {
	raw, err := r.session.ReadRaw(ctx)
	if err != nil {
		return models.MeshEvent{}, fmt.Errorf("read: %w", err)
	}

	var fr mesh.FromRadio

	if proto.Unmarshal(raw, &fr) != nil {
		return models.MeshEvent{}, nil //nolint:nilerr // a corrupt frame is skipped, not fatal
	}

	return r.decode(&fr), nil
}
