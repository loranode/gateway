package mesh

import (
	"context"
	"fmt"

	"github.com/loranode/gateway/internal/models"
)

// Read returns the next decoded event from the node, waiting through a reconnect.
func (s *Service) Read(ctx context.Context) (models.MeshEvent, error) {
	ev, err := s.repo.Read(ctx)
	if err != nil {
		return models.MeshEvent{}, fmt.Errorf("read: %w", err)
	}

	return ev, nil
}
