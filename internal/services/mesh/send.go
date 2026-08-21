package mesh

import (
	"context"
	"fmt"
)

// Send transmits a text message into the mesh, waiting through a reconnect.
func (s *Service) Send(ctx context.Context, text string, to, channel, replyID uint32) error {
	if err := s.repo.Send(ctx, text, to, channel, replyID); err != nil {
		return fmt.Errorf("send: %w", err)
	}

	return nil
}
