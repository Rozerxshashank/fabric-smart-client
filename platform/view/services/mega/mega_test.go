package mega

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMega_Run(t *testing.T) {
	m := New()
	assert.Equal(t, "Component A executed then Component B executed", m.Run())
}
