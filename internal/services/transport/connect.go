package transport

import (
	"context"
	"fmt"
)

// Connect dials the node, replacing any previous connection.
func (s *Service) Connect(ctx context.Context) error {
	if err := s.repo.Connect(ctx); err != nil {
		return fmt.Errorf("transport connect: %w", err)
	}

	return nil
}
