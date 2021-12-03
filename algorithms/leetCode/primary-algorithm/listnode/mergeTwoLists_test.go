package listnode

import (
	"log"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{-1,
		&ListNode{5,
			&ListNode{7,
				&ListNode{8,
					&ListNode{10,
						&ListNode{12, nil},
					}}}}}

	l2 := &ListNode{1,
		&ListNode{3,
			&ListNode{4,
				&ListNode{6,
					&ListNode{9,
						&ListNode{10, nil},
					}}}}}
	log.Println(MergeTwoLists(l1, l2))
}

func TestMergeTwoListsPro(t *testing.T) {
	l1 := &ListNode{-1,
		&ListNode{5,
			&ListNode{7,
				&ListNode{8,
					&ListNode{10,
						&ListNode{12, nil},
					}}}}}

	l2 := &ListNode{1,
		&ListNode{3,
			&ListNode{4,
				&ListNode{6,
					&ListNode{9,
						&ListNode{10, nil},
					}}}}}
	result := MergeTwoListsPro(l1, l2)
	ListNode2Array(result)
	ListNode2Array(l1)
	ListNode2Array(l2)
}

func TestMergeTwoLists2(t *testing.T) {
	l1 := &ListNode{-1,
		&ListNode{5,
			&ListNode{7,
				&ListNode{8,
					&ListNode{10,
						&ListNode{12, nil},
					}}}}}

	l2 := &ListNode{1,
		&ListNode{3,
			&ListNode{4,
				&ListNode{6,
					&ListNode{9,
						&ListNode{10, nil},
					}}}}}
	result := MergeTwoLists2(l1, l2)
	ListNode2Array(result)
	ListNode2Array(l1)
	ListNode2Array(l2)
}

func ListNode2Array(node *ListNode) {
	var nums []int
	for node != nil {
		nums = append(nums, node.Val)
		node = node.Next
	}
	log.Println(nums)
}
