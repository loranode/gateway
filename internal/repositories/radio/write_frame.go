package radio

import (
	"encoding/binary"
	"fmt"
)

// WriteFrame frames and sends a payload. It is safe to call concurrently with
// ReadFrame.
func (r *Repository) WriteFrame(payload []byte) error {
	if len(payload) > maxFrame {
		return ErrFrameTooLarge
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.conn == nil {
		return ErrNotConnected
	}

	var hdr [4]byte

	hdr[0] = start1
	hdr[1] = start2
	binary.BigEndian.PutUint16(hdr[2:], uint16(len(payload)))

	if _, err := r.conn.Write(hdr[:]); err != nil {
		return fmt.Errorf("radio write header: %w", err)
	}

	if _, err := r.conn.Write(payload); err != nil {
		return fmt.Errorf("radio write payload: %w", err)
	}

	return nil
}
