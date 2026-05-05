package magic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMagic_Cast(t *testing.T) {
	m := New("Abra-Kadabra")
	assert.Equal(t, "Casting spell: Abra-Kadabra!", m.Cast())
}
