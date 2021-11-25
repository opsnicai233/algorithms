package binaryTree

import (
	"log"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	node := &TreeNode{1, nil, nil}
	node.Left = &TreeNode{2, nil, &TreeNode{3, nil, nil}}
	node.Right = &TreeNode{2, nil, &TreeNode{3, nil, nil}}
	log.Println(IsSymmetric(node))
}

func TestIsSymmetric2(t *testing.T) {
	node := &TreeNode{1, nil, nil}
	node.Left = &TreeNode{2, nil, &TreeNode{3, nil, nil}}
	node.Right = &TreeNode{2, nil, &TreeNode{3, nil, nil}}
	log.Println(IsSymmetric2(node))
}

func TestIsSymmetricProMax(t *testing.T) {
	node := &TreeNode{1, nil, nil}
	node.Left = &TreeNode{2, nil, &TreeNode{3, nil, nil}}
	node.Right = &TreeNode{2, nil, &TreeNode{3, nil, nil}}
	log.Println(IsSymmetricProMax(node))
}
