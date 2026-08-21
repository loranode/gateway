// Package meshtastic is the node transport: it wraps the meshtastic session
// keeper, keeps the link alive (want_config, heartbeat, reconnect) underneath,
// and decodes the FromRadio stream into models. The session is created in
// Connect so this file stays free of the library (the package.go layer rule
// allows stdlib only); it is held behind a local interface.
package meshtastic

import "context"

// driver is the slice of the session keeper this repository needs; keeping it
// local lets package.go avoid importing the meshtastic library.
type driver interface {
	Connect(ctx context.Context)
	ReadRaw(ctx context.Context) ([]byte, error)
	SendRaw(ctx context.Context, frame []byte) error
}

// Repository is a keep-alive, decoding connection to a node's client API.
type Repository struct {
	addr    string
	session driver
}

// broadcast is the recipient value that marks a channel (broadcast) packet.
const broadcast = 0xffffffff

// New builds a node transport for addr. Call Connect before Read or Send.
func New(addr string) *Repository {
	return &Repository{addr: addr}
}
