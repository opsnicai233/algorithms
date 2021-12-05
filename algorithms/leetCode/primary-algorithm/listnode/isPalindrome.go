package listnode

func IsPalindrome(head *ListNode) bool {
	nodeLen, tmp, node := 0, head, head
	for tmp != nil {
		nodeLen++
		tmp = tmp.Next
	}
	for i := 0; i < nodeLen/2-1; i++ {
		node = node.Next
	}
	tmp = node.Next
	node.Next = nil
	tail := ReverseListNode(tmp)

	for i := 0; i < nodeLen/2; i++ {
		if head.Val != tail.Val {
			return false
		}
		head, tail = head.Next, tail.Next
	}
	return true
}

func ReverseListNode(node *ListNode) (tail *ListNode) {
	if node == nil || node.Next == nil {
		tail = node
		return
	}
	prev, curr := node, node.Next
	prev.Next = nil
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev, curr = curr, tmp
	}
	tail = prev
	return
}
