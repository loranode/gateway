package registry

import (
	"fmt"

	"github.com/loranode/gateway/internal/models"
)

// UpsertChannel stores a channel configuration, overwriting the previous entry
// for the same index.
func (s *Service) UpsertChannel(ch *models.Channel) error {
	if err := s.repo.UpsertChannel(ch); err != nil {
		return fmt.Errorf("registry upsert channel: %w", err)
	}

	return nil
}
