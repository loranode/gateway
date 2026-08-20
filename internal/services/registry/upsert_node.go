package registry

import (
	"fmt"

	"github.com/loranode/gateway/internal/models"
)

// UpsertNode merges a partial node record into the store, returning whether the
// node was newly created.
func (s *Service) UpsertNode(patch *models.Node) (bool, error) {
	created, err := s.repo.UpsertNode(patch)
	if err != nil {
		return false, fmt.Errorf("registry upsert node: %w", err)
	}

	return created, nil
}
