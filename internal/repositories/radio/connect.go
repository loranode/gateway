package radio

import (
	"bufio"
	"context"
	"fmt"
	"net"
)

// Connect dials the node and prepares the framed reader, replacing any previous
// connection.
func (r *Repository) Connect(ctx context.Context) error {
	var d net.Dialer

	conn, err := d.DialContext(ctx, "tcp", r.addr)
	if err != nil {
		return fmt.Errorf("radio dial: %w", err)
	}

	r.mu.Lock()
	r.conn = conn
	r.reader = bufio.NewReaderSize(conn, maxFrame)
	r.mu.Unlock()

	return nil
}
