package BinaryTree

import "fmt"

//* Definition for a binary tree node.
// leetcode 上对于二叉树的定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (treeNode *TreeNode) toString() string {
	return fmt.Sprintf("TreeNode: {%d, %v, %v}",
		treeNode.Val, treeNode.Left, treeNode.Right)
}
