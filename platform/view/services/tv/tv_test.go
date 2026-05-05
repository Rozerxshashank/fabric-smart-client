package tv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTV_Workflow(t *testing.T) {
	tv := New()
	
	// Step 1: TV on
	assert.Equal(t, "TV is now ON", tv.On())
	
	// Step 2: TV off
	assert.Equal(t, "TV is now OFF", tv.Off())
	
	// Step 3: Netflix on (requires TV on)
	assert.Equal(t, "TV is now ON", tv.On())
	assert.Equal(t, "Netflix is now ON", tv.NetflixOn())
}
