// Package meshtastic is the node transport: it wraps the meshtastic session
// keeper, keeps the link alive (want_config, heartbeat, reconnect) underneath,
// and decodes the FromRadio stream into models. The session is created in
// Connect so this file stays free of the library (the package.go layer rule
// allows stdlib only); it is held behind a local interface.
package meshtastic

import (
	"context"
	"sync"
	"sync/atomic"
)

// driver is the slice of the session keeper this repository needs; keeping it
// local lets package.go avoid importing the meshtastic library.
type driver interface {
	Connect(ctx context.Context)
	ReadRaw(ctx context.Context) ([]byte, error)
	SendRaw(ctx context.Context, frame []byte) error
}

// Repository is a keep-alive, decoding connection to a node's client API. The
// hop limit is captured from the node's config dump and read on every Send, so
// the two run on different goroutines and must go through the atomic. pkiNodes
// records which node numbers advertised a public key (from NodeInfo), so a
// direct message can request PKI encryption the same way the read loop and Send
// straddle goroutines.
type Repository struct {
	addr     string
	session  driver
	hopLimit atomic.Uint32
	pkiNodes sync.Map
}

// broadcast is the recipient value that marks a channel (broadcast) packet.
const broadcast = 0xffffffff

// defaultHopLimit seeds the hop limit until the node's config dump arrives. Zero
// (the proto default) would keep messages on direct neighbours only and starve
// broadcasts of the implicit ack, so match the Meshtastic client default of 5.
const defaultHopLimit = 5

// New builds a node transport for addr. Call Connect before Read or Send.
func New(addr string) *Repository {
	r := &Repository{addr: addr}
	r.hopLimit.Store(defaultHopLimit)

	return r
}
