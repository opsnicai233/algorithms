package listnode

import (
	"log"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	node := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	target := node.Next
	DeleteNode(target)
	log.Println(node)
}
