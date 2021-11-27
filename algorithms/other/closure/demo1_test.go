package closure

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetVal(t *testing.T) {
	assert.Equal(t, GetFunc2Val(), 10)
	assert.Equal(t, GetFunc2proVal(), 11)
}
