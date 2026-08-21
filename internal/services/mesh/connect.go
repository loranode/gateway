package mesh

import "context"

// Connect starts the node session and keeps it alive in the background.
func (s *Service) Connect(ctx context.Context) {
	s.repo.Connect(ctx)
}
