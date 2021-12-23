package BinaryTree

// 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。
// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
// 输出：true

type Tree struct {
	Data  int
	Left  *Tree
	Right *Tree
}

func IsExistSum(node *Tree, target, sum int) bool {
	if sum == target {
		return true
	}
	if node == nil {
		return sum == target
	}
	sum += node.Data
	return IsExistSum(node.Left, target, sum) || IsExistSum(node.Right, target, sum)
}

func Arr2Tree(arr []int) *Tree {
	n := len(arr)
	root := &Tree{arr[0], nil, nil}
	var queue = []*Tree{root}
	i, j := 0, 0
	for len(queue) > 0 {
		curr := queue[0]
		if arr[i] != Null {
			curr.Data = arr[i]
			queue = queue[1:]
		} else {
			i++
			j += 2
			continue
		}
		j++
		if curr.Left == nil && j < n && arr[j] != Null {
			curr.Left = new(Tree)
			queue = append(queue, curr.Left)
		}
		j++
		if curr.Right == nil && j < n && arr[j] != Null {
			curr.Right = new(Tree)
			queue = append(queue, curr.Right)
		}
		i++
	}
	return root
}

func Arr2Tree2(arr []int) *Tree {
	n := len(arr)
	if n == 0 || arr[0] == Null {
		return nil
	}
	var nodes []*Tree
	for i := 0; i < n; i++ {
		if arr[i] != Null {
			nodes = append(nodes, &Tree{arr[i], nil, nil})
		} else {
			nodes = append(nodes, nil)
		}
	}
	root := &Tree{arr[0], nil, nil}
	var queue = []*Tree{root}
	i := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr == nil {
			i += 2
			continue
		}
		i++
		if i < n {
			if curr.Left == nil {
				curr.Left = nodes[i]
				queue = append(queue, curr.Left)
			}
		} else {
			break
		}
		i++
		if i < n {
			if curr.Right == nil {
				curr.Right = nodes[i]
				queue = append(queue, curr.Right)
			}
		} else {
			break
		}
	}
	return root
}
