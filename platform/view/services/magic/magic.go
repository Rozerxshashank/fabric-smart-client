package magic

import (
	"fmt"
)

// Magic is a dummy service for the contribution process testbed (Task 3).
type Magic struct {
	spell string
}

// New returns a new Magic service.
func New(spell string) *Magic {
	return &Magic{spell: spell}
}

// Cast returns a magic message.
func (m *Magic) Cast() string {
	return fmt.Sprintf("Casting spell: %s!", m.spell)
}
