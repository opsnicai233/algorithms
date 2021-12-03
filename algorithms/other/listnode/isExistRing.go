package listnode

import "math"

// IsExistRing 判断一个链表是否存在环， 若存在返回重复节点的index（index最小值为0）,若不存在，返回-1
func IsExistRing(node *ListNode) int {
	if node.Next == nil {
		return -1
	}
	return check(node, node.Next, 0, 1)
}

// 这种写法只能得到一个概率性结果,并不是逻辑正确的，它希望在curr进入环之前js已经比环的节点数大，
// 否则返回的结果是环内某个节点的index，这个index可能比正确结果大
func check(curr *ListNode, next *ListNode, i int, js int) int {
	for counter := 0; counter < js; counter++ {
		if next == nil {
			return -1
		}
		if next == curr {
			return i
		} else {
			next = next.Next
		}
	}
	return check(curr.Next, next, i+1, js+1)
}

func IsExistRingPro(node *ListNode) int {
	if node.Next == nil {
		return -1
	}
	return checkpro(node, node.Next, 0, 1, math.MaxInt)
}

// 依旧是一个概率性算法，bug在于curr在没有循环到最小的index之前，next已经遍历环几圈了
func checkpro(curr *ListNode, next *ListNode, i int, js int, result int) int {
	for counter := 0; counter < js; counter++ {
		if next == nil {
			return -1
		}
		if next == curr {
			// 找result的最小值
			if result > i {
				result = i
			} else {
				return result
			}
		} else {
			next = next.Next
		}
	}
	return checkpro(curr.Next, next, i+1, js+1, result)
}
