package registry

import (
	"fmt"

	"github.com/loranode/gateway/internal/models"
)

// Nodes returns the compact info of every known node, most recently heard first.
func (s *Service) Nodes() ([]models.NodeInfo, error) {
	nodes, err := s.repo.Nodes()
	if err != nil {
		return nil, fmt.Errorf("registry nodes: %w", err)
	}

	return nodes, nil
}
