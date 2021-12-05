package listnode

// MergeTwoLists 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	target := new(ListNode)
	putNode(list1, list2, target)
	return target
}

// 并没有使用 list1 list2 中的节点进行拼接
func putNode(curr1, curr2, result *ListNode) {
	if curr1 == nil && curr2 == nil {
		result = nil
	}
	if curr1 == nil {
		result = curr2
		return
	}
	if curr2 == nil {
		result = curr1
		return
	}
	if curr1.Val < curr2.Val {
		result.Val = curr1.Val
		result.Next = new(ListNode)
		putNode(curr1.Next, curr2, result.Next)
	}
	if curr1.Val > curr2.Val {
		result.Val = curr2.Val
		result.Next = new(ListNode)
		putNode(curr1, curr2.Next, result.Next)
	}
	if curr1.Val == curr2.Val {
		result.Val = curr1.Val
		result.Next = new(ListNode)
		result.Next.Val = curr2.Val
		result.Next.Next = new(ListNode)
		putNode(curr1.Next, curr2.Next, result.Next.Next)
	}
	return
}

func MergeTwoListsPro(list1 *ListNode, list2 *ListNode) *ListNode {
	return putNodePro(list1, list2)
}

// 尝试使用 list1、list2 中的节点进行拼接
// 递归写法
func putNodePro(curr1, curr2 *ListNode) (head *ListNode) {
	if curr1 == nil && curr2 == nil {
		head = nil
		return
	}
	if curr1 == nil {
		head = curr2
		return
	}
	if curr2 == nil {
		head = curr1
		return
	}
	if curr1.Val < curr2.Val {
		head = curr1
		head.Next = putNodePro(head.Next, curr2)
		//return
	}
	if curr1.Val > curr2.Val {
		head = curr2
		head.Next = putNodePro(curr1, head.Next)
		//return
	}
	if curr1.Val == curr2.Val {
		tmp1 := curr1.Next
		head = curr1
		head.Next = curr2
		head.Next.Next = putNodePro(tmp1, head.Next.Next)
		//return
	}
	return
}

func MergeTwoListsPro2(list1 *ListNode, list2 *ListNode) *ListNode {
	return putNodeProMax(list1, list2)
}

func putNodeProMax(curr1, curr2 *ListNode) (head *ListNode) {
	if curr1 == nil && curr2 == nil {
		head = nil
		return
	}
	if curr1 == nil {
		head = curr2
		return
	}
	if curr2 == nil {
		head = curr1
		return
	}
	if curr1.Val < curr2.Val {
		head = curr1
		head.Next = putNodeProMax(head.Next, curr2)
		//return
	} else {
		head = curr2
		head.Next = putNodeProMax(curr1, head.Next)
	}
	return
}

// MergeTwoLists2 尝试使用 list1、list2 中的节点进行拼接
// 迭代写法
func MergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	var head = &ListNode{-1, nil}
	result := head
	for list1 != nil || list2 != nil {
		if list1 == nil {
			head.Next = list2
			break
		}
		if list2 == nil {
			head.Next = list1
			break
		}
		if list1.Val < list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}
	return result.Next
}
