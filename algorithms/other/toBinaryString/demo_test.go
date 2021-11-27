package toBinaryString

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestToBinaryString(t *testing.T) {
	assert.Equal(t, ToBinaryString(5), "101")
	assert.Equal(t, ToBinaryString(8), "1000")
}
