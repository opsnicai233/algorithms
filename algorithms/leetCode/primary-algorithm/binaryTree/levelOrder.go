package binaryTree

//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//示例：
//二叉树：[3,9,20,null,null,15,7]
//[
//[3],
//[9,20],
//[15,7]
//]

// LevelOrder 迭代写法
func LevelOrder(root *TreeNode) [][]int {
	var result [][]int
	var currLevelNodes = []*TreeNode{root}
	for len(currLevelNodes) > 0 {
		var values []int
		n := len(currLevelNodes)
		for _, node := range currLevelNodes {
			if node != nil {
				values = append(values, node.Val)
				currLevelNodes = append(currLevelNodes, node.Left, node.Right)
			}
		}
		currLevelNodes = currLevelNodes[n:]
		if len(values) > 0 {
			result = append(result, values)
		}

	}
	return result
}

// LevelOrder2 递归写法
func LevelOrder2(root *TreeNode) [][]int {
	var result [][]int
	ReadLevel(root, 0, &result)
	return result
}

func ReadLevel(node *TreeNode, level int, result *[][]int) {
	if node == nil {
		return
	}
	for len(*result) < level+1 {
		*result = append(*result, []int{})
	}
	(*result)[level] = append((*result)[level], node.Val)
	level++
	ReadLevel(node.Left, level, result)
	ReadLevel(node.Right, level, result)
}
