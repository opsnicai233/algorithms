package BinaryTree

import (
	"fmt"
	"testing"
)

func TestArrayToTree(t *testing.T) {
	arr := []int{5, 4, 8, 11, Null, 13, 4, 7, 2, Null, Null, Null, 1}
	node := ArrayToTree(arr)

	fmt.Println(node)
}
