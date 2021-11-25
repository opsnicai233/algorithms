package binaryTree

import (
	"math"
)

func maxDepth(root *TreeNode) int {
	return GoDeep(root)
}

func GoDeep(node *TreeNode) int {
	if node == nil {
		return 0
	}
	lDepth := GoDeep(node.Left)
	rDepth := GoDeep(node.Right)
	return int(math.Max(float64(lDepth), float64(rDepth))) + 1
}
