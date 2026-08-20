package events

import (
	"context"
	"fmt"
)

// Delete unsubscribes a webhook URL.
func (s *Service) Delete(ctx context.Context, url string) error {
	if err := s.store.DeleteCallback(ctx, url); err != nil {
		return fmt.Errorf("delete callback: %w", err)
	}

	return nil
}
