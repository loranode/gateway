package events

import "github.com/loranode/gateway/internal/models"

// Subscribe registers a subscriber and returns its event channel and an
// unsubscribe func. The channel is buffered; a slow subscriber drops events.
func (s *Service) Subscribe() (<-chan models.Event, func()) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.next++
	id := s.next
	ch := make(chan models.Event, subBuffer)
	s.subs[id] = ch

	return ch, func() {
		s.mu.Lock()
		defer s.mu.Unlock()

		if c, ok := s.subs[id]; ok {
			delete(s.subs, id)
			close(c)
		}
	}
}
