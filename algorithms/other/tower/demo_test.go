package tower

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestTower(t *testing.T) {
	Tower(3, "a", "b", "c")
	assert.Equal(t, Total, 7)
}
