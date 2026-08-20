package registry

import (
	"fmt"

	"github.com/loranode/gateway/internal/models"
)

// NodeByNum returns the full record of one node by its number; the second result
// is false when it does not exist.
func (s *Service) NodeByNum(num uint32) (models.Node, bool, error) {
	node, ok, err := s.repo.NodeByNum(num)
	if err != nil {
		return models.Node{}, false, fmt.Errorf("registry node by num: %w", err)
	}

	return node, ok, nil
}
