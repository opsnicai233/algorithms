package listnode

import (
	"log"
	"testing"
)

func TestReverseListNode(t *testing.T) {
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

	log.Println(ReverseListNode(node))
}

func TestEdit(t *testing.T) {
	a := &ListNode{0, nil}
	Edit(a)
	log.Println(a)
}
