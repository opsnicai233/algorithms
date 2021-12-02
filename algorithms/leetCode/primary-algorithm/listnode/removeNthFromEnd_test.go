package listnode

import (
	"log"
	"testing"
)

func TestRemoveNthFromEnd2(t *testing.T) {
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
	log.Println(RemoveNthFromEndPro(node, 1))
}

func TestRemoveNthFromEndProMax(t *testing.T) {
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
	log.Println(RemoveNthFromEndProMax(node, 2))
}
