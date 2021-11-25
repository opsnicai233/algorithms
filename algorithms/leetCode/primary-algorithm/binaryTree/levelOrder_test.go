package binaryTree

import (
	"log"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	node := &TreeNode{3, nil, nil}
	node.Left = &TreeNode{9, nil, nil}
	node.Right = &TreeNode{
		20,
		&TreeNode{15, nil, nil},
		&TreeNode{7, nil, nil},
	}
	log.Println(LevelOrder(node))
}

func TestLevelOrder2(t *testing.T) {
	node := &TreeNode{3, nil, nil}
	node.Left = &TreeNode{9, nil, nil}
	node.Right = &TreeNode{
		20,
		&TreeNode{15, nil, nil},
		&TreeNode{7, nil, nil},
	}
	log.Println(LevelOrder2(node))
}
