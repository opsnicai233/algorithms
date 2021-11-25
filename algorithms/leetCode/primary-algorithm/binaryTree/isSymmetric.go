package binaryTree

// IsSymmetric 迭代法，效率并不高
//执行用时：4 ms, 在所有 Go 提交中击败了28.92%的用户
//内存消耗：3.1 MB, 在所有 Go 提交中击败了9.35%的用户
func IsSymmetric(root *TreeNode) bool {
	var nodeArr []*TreeNode
	nodeArr = append(nodeArr, root, root)
	for len(nodeArr) > 0 {
		l, r := nodeArr[0], nodeArr[1]
		nodeArr = nodeArr[2:]
		if l == nil && r == nil {
			continue
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		nodeArr = append(nodeArr, l.Left)
		nodeArr = append(nodeArr, r.Right)
		nodeArr = append(nodeArr, l.Right)
		nodeArr = append(nodeArr, r.Left)
	}
	return true
}

func IsSymmetric2(root *TreeNode) bool {
	var lNodes, rNodes []int
	getLeft(root.Left, &lNodes)
	getRight(root.Right, &rNodes)
	//log.Printf("lNodes: %v\n", lNodes)
	if len(lNodes) != len(rNodes) {
		return false
	}
	for i := 0; i < len(lNodes); i++ {
		if lNodes[i] != rNodes[i] {
			return false
		}
	}
	return true
}

func getRight(root *TreeNode, nodes *[]int) {
	if root != nil {
		*nodes = append(*nodes, root.Val)
		getRight(root.Right, nodes)
		getRight(root.Left, nodes)
	} else {
		*nodes = append(*nodes, null)
	}
}

func getLeft(root *TreeNode, nodes *[]int) {
	if root != nil {
		*nodes = append(*nodes, root.Val)
		getLeft(root.Left, nodes)
		getLeft(root.Right, nodes)
	} else {
		*nodes = append(*nodes, null)
	}
	//log.Println(nodes)
}

func IsSymmetricProMax(root *TreeNode) bool {
	return helper(root.Left, root.Right)
}

func helper(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}
