package binaryTree

import (
	"log"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	node := &TreeNode{3, nil, nil}
	node.Left = &TreeNode{9, nil, nil}
	node.Right = &TreeNode{
		20,
		&TreeNode{15, nil, nil},
		&TreeNode{7, nil, nil},
	}
	log.Println(maxDepth(node))
}
