package meshtastic

import (
	"context"

	"github.com/loranode/meshtastic/session"
)

// Connect opens the session and keeps it alive in the background. It returns once
// the session exists, so Read and Send (which block through reconnects) are safe
// to call afterwards.
func (r *Repository) Connect(ctx context.Context) {
	s := session.New(session.Config{Address: r.addr})
	r.session = s

	go s.Connect(ctx)
}
