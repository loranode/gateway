package events

import (
	"context"
	"fmt"
)

// List returns every registered webhook URL.
func (s *Service) List(ctx context.Context) ([]string, error) {
	urls, err := s.store.Callbacks(ctx)
	if err != nil {
		return nil, fmt.Errorf("list callbacks: %w", err)
	}

	return urls, nil
}
