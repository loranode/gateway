package webhook

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

// Post delivers one JSON body to one URL with a single attempt, returning an
// error on transport failure or a non-2xx response.
func (r *Repository) Post(ctx context.Context, url string, body []byte) error {
	reqCtx, cancel := context.WithTimeout(ctx, postTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(reqCtx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("build webhook request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		return fmt.Errorf("post webhook: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("%w: %d", ErrNon2xx, resp.StatusCode)
	}

	return nil
}
