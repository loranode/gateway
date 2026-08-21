// Package events broadcasts mesh events to every live subscriber (the gRPC event
// stream). Subscribers are in-memory; nothing is persisted.
package events

import (
	"sync"

	"github.com/loranode/gateway/internal/models"
)

// Service fans mesh events out to subscribers.
type Service struct {
	mu   sync.Mutex
	subs map[uint64]chan models.Event
	next uint64
}

const (
	// subBuffer bounds a subscriber's pending events.
	subBuffer = 64

	typeNodeCreated           = "node_created"
	typeNodeUpdated           = "node_updated"
	typeNodeMessageCreated    = "node_message_created"
	typeChannelMessageCreated = "channel_message_created"
)

// New builds an events service with no subscribers.
func New() *Service {
	return &Service{subs: make(map[uint64]chan models.Event)}
}
