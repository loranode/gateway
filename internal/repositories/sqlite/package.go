// Package sqlite is the persistent store of the state the service builds from
// the node's stream: the known nodes and every text message received. There are
// no size or history limits — messages accumulate indefinitely.
package sqlite

import (
	"database/sql"
	"fmt"
)

// Repository owns the database handle.
type Repository struct {
	db *sql.DB
}

// New opens (creating the file if needed) the database at path. Call Migrate to
// apply the schema before use.
func New(path string) (*Repository, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	return &Repository{db: db}, nil
}
