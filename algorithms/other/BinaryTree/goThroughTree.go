package BinaryTree

var arrDeep []interface{}

// 深度优先遍历二叉树
func GoDeepThroughTree(node *Node) {
	data := node.GetData()
	arrDeep = append(arrDeep, data)
	if data == null {
		return
	}
	GoDeepThroughTree(node.GetNodeLeft())
	GoDeepThroughTree(node.GetNodeRight())
}

// 广度优先遍历二叉树
func GoWidthThroughTree(node *Node) []interface{} {
	arrWidth := []interface{}{}
	queue := []*Node{}
	queue = append(queue, node)
	arrWidth = append(arrWidth, node.GetData())
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		left := curr.GetNodeLeft()
		arrWidth = append(arrWidth, left.GetData())
		if !left.IsEmpty() {
			queue = append(queue, left)
		}

		right := curr.GetNodeRight()
		arrWidth = append(arrWidth, right.GetData())
		if !right.IsEmpty() {
			queue = append(queue, right)
		}
	}
	return arrWidth
}
