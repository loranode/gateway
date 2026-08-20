package transport

import "fmt"

// WriteFrame sends a framed payload to the node.
func (s *Service) WriteFrame(payload []byte) error {
	if err := s.repo.WriteFrame(payload); err != nil {
		return fmt.Errorf("transport write: %w", err)
	}

	return nil
}
