package events

import (
	"context"
	"fmt"
)

// Register subscribes a webhook URL to mesh events.
func (s *Service) Register(ctx context.Context, url string) error {
	if err := s.store.AddCallback(ctx, url); err != nil {
		return fmt.Errorf("register callback: %w", err)
	}

	return nil
}
