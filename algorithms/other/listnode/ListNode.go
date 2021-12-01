package listnode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	return fmt.Sprintf("ListNode{%d, %s}", node.Val, node.Next)
}
