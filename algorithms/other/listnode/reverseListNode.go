package listnode

func ReverseListNode(node *ListNode) *ListNode {
	if node == nil {
		return node
	}
	curr := node.Next
	node.Next = nil
	return reverse(curr, node)
	//node.Next = nil
}

func reverse(curr, prev *ListNode) (head *ListNode) {
	if curr == nil || prev == nil {
		head = prev
		return
	}
	tmp := curr.Next
	curr.Next = prev
	return reverse(tmp, curr)
}

func Edit(a *ListNode) {
	a.Next = &ListNode{99, nil}
}
