package listnode

// 给你一个链表的头节点 head ，判断链表中是否有环。

func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	return check(head, head.Next, 2)
}

func check(curr *ListNode, next *ListNode, js int) bool {
	for counter := 0; counter < js; counter++ {
		if next == nil {
			return false
		}
		if next == curr {
			return true
		} else {
			next = next.Next
		}
	}
	return check(curr.Next, next, js)
}

func HasCyclePro(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && slow != nil {
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}
