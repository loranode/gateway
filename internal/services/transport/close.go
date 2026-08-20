package transport

import "fmt"

// Close drops the current connection.
func (s *Service) Close() error {
	if err := s.repo.Close(); err != nil {
		return fmt.Errorf("transport close: %w", err)
	}

	return nil
}
