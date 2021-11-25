package binaryTree

import "math"

func isValidBST(root *TreeNode) bool {
	return checkCurrTree(root, math.MinInt64, math.MaxInt64)
}

func checkCurrTree(node *TreeNode, lower, higher int) bool {
	if node == nil {
		return true
	}

	if node.Val <= lower || node.Val >= higher {
		return false
	}
	return checkCurrTree(node.Left, lower, node.Val) && checkCurrTree(node.Right, node.Val, higher)
}
