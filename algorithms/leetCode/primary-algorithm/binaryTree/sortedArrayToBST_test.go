package binaryTree

import (
	"log"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	log.Println(SortedArrayToBST(arr))
}
