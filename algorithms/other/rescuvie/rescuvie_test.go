package rescuvie

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestRescuvie(t *testing.T) {
	assert.Equal(t, Rescuvie(1), 1)
	assert.Equal(t, Rescuvie(5), 120)
	assert.Equal(t, Rescuvie(3), 6)
}
