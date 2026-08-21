package events

import "github.com/loranode/gateway/internal/models"

// publish broadcasts an event to all subscribers, dropping it for any whose
// buffer is full rather than blocking.
func (s *Service) publish(ev models.Event) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, ch := range s.subs {
		select {
		case ch <- ev:
		default:
		}
	}
}
