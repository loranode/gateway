package transport

import "fmt"

// ReadFrame reads the next framed payload from the node.
func (s *Service) ReadFrame() ([]byte, error) {
	b, err := s.repo.ReadFrame()
	if err != nil {
		return nil, fmt.Errorf("transport read: %w", err)
	}

	return b, nil
}
