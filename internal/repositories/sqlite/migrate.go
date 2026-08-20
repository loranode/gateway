package sqlite

import (
	"fmt"

	"github.com/pressly/goose/v3"

	"github.com/loranode/gateway/migrations"
)

// Migrate applies all embedded goose migrations to the database.
func (r *Repository) Migrate() error {
	goose.SetBaseFS(migrations.FS)

	if err := goose.SetDialect("sqlite3"); err != nil {
		return fmt.Errorf("set dialect: %w", err)
	}

	if err := goose.Up(r.db, "."); err != nil {
		return fmt.Errorf("apply migrations: %w", err)
	}

	return nil
}
