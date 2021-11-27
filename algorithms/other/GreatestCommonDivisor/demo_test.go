package GreatestCommonDivisor

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestGcd(t *testing.T) {
	assert.Equal(t, Gcd(24, 3), uint32(3))
	assert.Equal(t, Gcd(5, 24), uint32(1))
}
