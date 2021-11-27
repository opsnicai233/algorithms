package binaryTree

// SortedArrayToBST 将一个有序数组变为高度平衡的二叉搜索树
func SortedArrayToBST(nums []int) *TreeNode {
	start, end := 0, len(nums)
	return PutNode(nums, start, end)
}

func PutNode(arr []int, start, end int) *TreeNode {
	if start >= end {
		return nil
	}

	mid := (start + end) / 2
	tree := &TreeNode{arr[mid], nil, nil}
	tree.Left = PutNode(arr, start, mid)
	tree.Right = PutNode(arr, mid+1, end)
	return tree
}
