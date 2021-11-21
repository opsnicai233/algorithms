package array

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	RemoveDuplicates(arr)
	arr = []int{1, 1, 2}
	RemoveDuplicates(arr)

}

func TestRemoveDuplicates2(t *testing.T) {
	arr := []int{0, 0, 0, 3}
	removeDuplicatesPro(arr)
}
