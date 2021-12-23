package BinaryTree

import (
	"fmt"
	"testing"
)

func TestArr2Tree(t *testing.T) {
	arr := []int{5, 4, 8, 11, Null, 13, 4, 7, 2, Null, Null, Null, 1}
	node := Arr2Tree2(arr)

	fmt.Println(IsExistSum(node, 26, 0))
}
