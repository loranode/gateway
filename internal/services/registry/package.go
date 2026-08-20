// Package registry is the service in front of the SQLite store: it is the
// controller's gateway for saving decoded node and message state and for
// answering the REST queries.
package registry

import "github.com/loranode/gateway/internal/repositories/sqlite"

// Service wraps the persistent store.
type Service struct {
	repo *sqlite.Repository
}

// New builds the registry service over repo.
func New(repo *sqlite.Repository) *Service {
	return &Service{repo: repo}
}
