package closure_test

import (
	"learn-gin/interview/closure"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGetSums(t *testing.T) {
	sum1, sum2 := closure.GetSums()
	assert.Equal(t, sum1, 1)
	assert.Equal(t, sum2, 2)
}
