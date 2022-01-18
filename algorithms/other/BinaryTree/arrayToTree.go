package BinaryTree

func ArrayToTree(array []int) *Tree {
	n := len(array)
	if n <= 0 {
		return nil
	}
	var nodes []*Tree
	for i := 0; i < n; i++ {
		if array[i] != Null {
			nodes = append(nodes, &Tree{array[i], nil, nil})
		} else {
			nodes = append(nodes, nil)
		}
	}
	queue := []*Tree{nodes[0]}

	i := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr == nil {
			i += 2
			continue
		}
		i += 1
		if i < n {
			curr.Left = nodes[i]
			queue = append(queue, nodes[i])
		}

		i += 1
		if i < n {
			curr.Right = nodes[i]
			queue = append(queue, nodes[i])
		}
	}
	return nodes[0]
}
