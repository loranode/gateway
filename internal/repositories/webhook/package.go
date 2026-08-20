// Package webhook is the outbound HTTP transport for delivering events to
// registered callback URLs.
package webhook

import (
	"errors"
	"net/http"
	"time"
)

// Repository POSTs event bodies to callback URLs.
type Repository struct {
	client *http.Client
}

const postTimeout = 5 * time.Second

// ErrNon2xx is returned when a callback responds with a non-2xx status.
var ErrNon2xx = errors.New("webhook returned non-2xx status")

// New returns a webhook repository with a per-request timeout.
func New() *Repository {
	return &Repository{client: &http.Client{Timeout: postTimeout}}
}
