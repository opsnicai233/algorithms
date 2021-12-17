package BinaryTree

// 目标： 从最底层横向把节点值放到数组中

func GetNodeDataDesc(root *TreeNode, res *[]int) {
	if root.Left != nil {
		GetNodeDataDesc(root.Left, res)
	}
	if root.Right != nil {
		GetNodeDataDesc(root.Right, res)
	}
	*res = append(*res, root.Val)
}

func ModifyArr(nums *[]int) {
	*nums = append(*nums, 22)
}
