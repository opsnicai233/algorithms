package isPrime

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestIsPrime(t *testing.T) {
	assert.Equal(t, IsPrime(2), true)
	assert.Equal(t, IsPrime(4), false)
	assert.Equal(t, IsPrime(8), false)
}
