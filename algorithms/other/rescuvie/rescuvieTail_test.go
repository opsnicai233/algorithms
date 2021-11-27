package rescuvie

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

// 尾递归
func TestRescuvieTail(t *testing.T) {
	assert.Equal(t, RescuvieTail(1, 1), 1)
	assert.Equal(t, RescuvieTail(5, 1), 120)
	assert.Equal(t, RescuvieTail(3, 1), 6)
}
