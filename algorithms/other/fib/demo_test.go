package fib

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestFib(t *testing.T) {
	assert.Equal(t, Fib(1, 1, 1), 1)
	assert.Equal(t, Fib(5, 1, 1), 5)
}
