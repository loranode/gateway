package registry

import (
	"fmt"

	"github.com/loranode/gateway/internal/models"
)

// AddMessage stores a received text message, returning whether it was inserted.
func (s *Service) AddMessage(m *models.Message) (bool, error) {
	inserted, err := s.repo.AddMessage(m)
	if err != nil {
		return false, fmt.Errorf("registry add message: %w", err)
	}

	return inserted, nil
}
