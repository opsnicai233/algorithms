package string

import (
	"gopkg.in/go-playground/assert.v1"
	"math"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	assert.Equal(t, ReverseInteger(-123), -321)
	assert.Equal(t, ReverseInteger(-120), -21)

	assert.Equal(t, int(math.Pow(2, 31))-1, 1<<31-1)
}
