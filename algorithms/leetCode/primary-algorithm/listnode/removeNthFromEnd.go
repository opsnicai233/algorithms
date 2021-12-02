package listnode

// RemoveNthFromEnd 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	var nodes []*ListNode
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	nums := len(nodes)
	if nums == 1 {
		return nil
	}
	target := nodes[nums-n]
	if target.Next == nil {
		nodes[nums-n-1].Next = nil
	} else {
		// *target = *target.Next
		target.Val = target.Next.Val
		target.Next = target.Next.Next
	}

	return nodes[0]
}

// RemoveNthFromEnd2 倒数第N个节点到最后一个节点的距离应该和第N个节点到第一个节点的距离一样
func RemoveNthFromEndPro(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	i, fast, slow := 2, head.Next, head
	for i < n {
		fast = fast.Next
		i++
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	if n == 1 {
		slow.Next = fast.Next
	} else {
		slow.Val = slow.Next.Val
		slow.Next = slow.Next.Next
	}
	return head
}

func RemoveNthFromEndProMax(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	i, fast, slow := 2, head.Next, head
	for fast.Next != nil {
		if i >= n {
			slow = slow.Next
		}
		fast = fast.Next
		i++
	}

	if n == 1 {
		slow.Next = nil
	} else {
		slow.Val = slow.Next.Val
		slow.Next = slow.Next.Next
	}
	return head
}
