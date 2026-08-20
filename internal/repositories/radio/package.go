// Package radio is the transport to a physical Meshtastic node over its TCP
// stream API. It deals only in raw framed payloads; decoding is the mappers'
// job. Meshtastic frames every protobuf with a 4-byte header,
// 0x94 0xc3 <len_hi> <len_lo>, whose magic bytes let a reader resynchronise
// after log text or noise on the same stream.
package radio

import (
	"bufio"
	"errors"
	"net"
	"sync"
)

// Repository owns a single connection to the node.
type Repository struct {
	addr   string
	mu     sync.Mutex
	conn   net.Conn
	reader *bufio.Reader
}

const (
	start1   = 0x94
	start2   = 0xc3
	maxFrame = 512 // firmware's MAX_TO_FROM_RADIO_SIZE
)

// ErrNotConnected is returned when a read or write happens before Connect.
var ErrNotConnected = errors.New("radio: not connected")

// ErrFrameTooLarge guards against a corrupt length header requesting an absurd
// allocation.
var ErrFrameTooLarge = errors.New("radio: frame exceeds maximum size")

// New returns a repository that talks to the node at addr.
func New(addr string) *Repository {
	return &Repository{addr: addr}
}
