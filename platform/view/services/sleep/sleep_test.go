package sleep

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSleep_Zzz(t *testing.T) {
	s := New(8)
	assert.Equal(t, "Zzz... Sleeping for 8 hours.", s.Zzz())
}
