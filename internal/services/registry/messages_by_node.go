package registry

import (
	"fmt"

	"github.com/loranode/gateway/internal/models"
)

// MessagesByNode returns the direct messages received from one node (those it
// sent to us, excluding broadcasts), newest first.
func (s *Service) MessagesByNode(from uint32) ([]models.Message, error) {
	messages, err := s.repo.MessagesByNode(from)
	if err != nil {
		return nil, fmt.Errorf("registry messages by node: %w", err)
	}

	return messages, nil
}
