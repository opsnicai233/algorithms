package BinaryTree

import (
	"log"
	"testing"
)

func TestGetNodeDataDesc(t *testing.T) {
	var node = &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}

	node.Left = &TreeNode{2, nil, nil}
	node.Right = &TreeNode{3, nil, nil}
	node.Left.Left = &TreeNode{4, nil, nil}
	node.Left.Right = &TreeNode{7, nil, nil}
	var res []int
	GetNodeDataDesc(node, &res)
	log.Println(res)
}

func TestModifyArr(t *testing.T) {
	var arr []int
	ModifyArr(&arr)
	log.Println(arr)
}
