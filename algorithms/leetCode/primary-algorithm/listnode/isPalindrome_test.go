package listnode

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestReverseListNode(t *testing.T) {
	l1 := &ListNode{-1,
		&ListNode{5,
			&ListNode{7,
				&ListNode{8,
					&ListNode{10,
						&ListNode{12, nil},
					}}}}}
	node := ReverseListNode(l1)
	ListNode2Array(node)
}

func TestIsPalindrome(t *testing.T) {
	l1 := &ListNode{11,
		&ListNode{5,
			&ListNode{7,
				&ListNode{7,
					&ListNode{5,
						&ListNode{11, nil},
					}}}}}
	assert.Equal(t, IsPalindrome(l1), true)

	l2 := new(ListNode)
	assert.Equal(t, IsPalindrome(l2), true)

}
