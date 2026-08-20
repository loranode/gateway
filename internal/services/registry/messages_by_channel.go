package registry

import (
	"fmt"

	"github.com/loranode/gateway/internal/models"
)

// MessagesByChannel returns the messages received on one channel, newest first.
func (s *Service) MessagesByChannel(channel uint32) ([]models.Message, error) {
	messages, err := s.repo.MessagesByChannel(channel)
	if err != nil {
		return nil, fmt.Errorf("registry messages by channel: %w", err)
	}

	return messages, nil
}
