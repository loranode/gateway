// Package migrations holds the embedded goose SQL migrations for the database.
package migrations

import "embed"

// FS holds every migration file, embedded into the binary.
//
//go:embed *.sql
var FS embed.FS
