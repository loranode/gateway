package radio

import "fmt"

// Close drops the current connection, unblocking any in-flight ReadFrame.
func (r *Repository) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.conn == nil {
		return nil
	}

	err := r.conn.Close()
	r.conn = nil
	r.reader = nil

	if err != nil {
		return fmt.Errorf("radio close: %w", err)
	}

	return nil
}
