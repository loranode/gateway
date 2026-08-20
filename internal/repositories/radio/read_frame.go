package radio

import (
	"encoding/binary"
	"fmt"
	"io"
)

// ReadFrame reads the next framed payload, skipping any bytes before the next
// valid header. It returns (nil, nil) for an empty frame so the caller can
// simply continue.
func (r *Repository) ReadFrame() ([]byte, error) {
	r.mu.Lock()
	reader := r.reader
	r.mu.Unlock()

	if reader == nil {
		return nil, ErrNotConnected
	}

	// Advance to the next start1,start2 header pair.
	for {
		b, err := reader.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("radio sync: %w", err)
		}

		if b != start1 {
			continue
		}

		b2, err := reader.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("radio sync: %w", err)
		}

		if b2 == start2 {
			break
		}

		// b2 might itself begin the next header; put it back.
		if b2 == start1 {
			_ = reader.UnreadByte()
		}
	}

	var lb [2]byte

	if _, err := io.ReadFull(reader, lb[:]); err != nil {
		return nil, fmt.Errorf("radio read length: %w", err)
	}

	n := int(binary.BigEndian.Uint16(lb[:]))
	if n == 0 {
		return nil, nil
	}

	if n > maxFrame {
		return nil, ErrFrameTooLarge
	}

	buf := make([]byte, n)

	if _, err := io.ReadFull(reader, buf); err != nil {
		return nil, fmt.Errorf("radio read payload: %w", err)
	}

	return buf, nil
}
