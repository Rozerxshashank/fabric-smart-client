package sleep

import (
	"fmt"
)

// Sleep represents a sleep service.
type Sleep struct {
	hours int
}

// New returns a new Sleep service.
func New(hours int) *Sleep {
	return &Sleep{hours: hours}
}

// Zzz returns a sleep message.
func (s *Sleep) Zzz() string {
	return fmt.Sprintf("Zzz... Sleeping for %d hours.", s.hours)
}
