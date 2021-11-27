package binaryserrch

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{
		1, 3, 5, 7, 9, 18, 33, 46, 74, 89,
	}

	assert.Equal(t, BinarySearch(5, arr), 2)
	assert.Equal(t, BinarySearch(45, arr), -1)
}
