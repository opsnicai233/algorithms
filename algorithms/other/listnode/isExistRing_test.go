package listnode

import (
	"log"
	"testing"
)

func TestIsExistRing(t *testing.T) {
	node := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	node.Next.Next.Next.Next = node.Next.Next.Next
	log.Println(IsExistRing(node))
}

func TestIsExistRingPro(t *testing.T) {
	node := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	node.Next.Next.Next.Next = node.Next.Next
	log.Println(IsExistRingPro(node))
}
