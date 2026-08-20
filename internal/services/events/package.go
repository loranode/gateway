// Package events posts mesh events to every registered webhook callback,
// synchronously and best-effort (one attempt per callback).
package events

import (
	"github.com/loranode/gateway/internal/repositories/sqlite"
	"github.com/loranode/gateway/internal/repositories/webhook"
)

// Service stores webhook callbacks and posts events to them.
type Service struct {
	store   *sqlite.Repository
	webhook *webhook.Repository
}

const (
	typeNodeCreated           = "node_created"
	typeNodeUpdated           = "node_updated"
	typeNodeMessageCreated    = "node_message_created"
	typeChannelMessageCreated = "channel_message_created"
)

// New builds the events service over the store and webhook repository.
func New(store *sqlite.Repository, wh *webhook.Repository) *Service {
	return &Service{store: store, webhook: wh}
}
