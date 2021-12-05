package listnode

import "log"

func ListNode2Array(node *ListNode) {
	var nums []int
	for node != nil {
		nums = append(nums, node.Val)
		node = node.Next
	}
	log.Println(nums)
}
