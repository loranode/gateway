package registry

import (
	"fmt"

	"github.com/loranode/gateway/internal/models"
)

// Channels returns every configured channel, ordered by index.
func (s *Service) Channels() ([]models.Channel, error) {
	channels, err := s.repo.Channels()
	if err != nil {
		return nil, fmt.Errorf("registry channels: %w", err)
	}

	return channels, nil
}
