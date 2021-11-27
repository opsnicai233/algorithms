package BinaryTree

// 默认 arr[0] 为root节点
func ArrayToBinaryTree(arr []Data) *Node {
	var nodes []*Node
	for _, data := range arr {
		node := new(Node)
		if data != nil {
			node.SetData(data)
		}
		nodes = append(nodes, node)
	}
	i := 1
	var queue []*Node
	queue = append(queue, nodes[0])
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if i < len(nodes) {
			if node.GetNodeLeft().IsEmpty() {
				node.SetLeft(nodes[i])
				queue = append(queue, nodes[i])
			}
		} else {
			break
		}

		i += 1
		if i < len(nodes) {
			if node.GetNodeRight().IsEmpty() {
				node.SetRight(nodes[i])
				queue = append(queue, nodes[i])
			}
		} else {
			break
		}
		i += 1
	}

	return nodes[0]
}
