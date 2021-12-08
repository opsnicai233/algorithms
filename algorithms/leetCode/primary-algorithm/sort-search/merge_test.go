package sort_search

import (
	"log"
	"testing"
)

func TestMerge(t *testing.T) {
	nums1 := []int{3, 5, 0, 0, 0, 0, 0}
	nums2 := []int{1, 2, 4, 5, 9}
	Merge(nums1, 2, nums2, 5)
	log.Println(nums1)
}

func TestMerge2(t *testing.T) {
	nums1 := []int{3, 5, 0, 0, 0, 0, 0}
	nums2 := []int{1, 2, 4, 5, 9}
	Merge2(nums1, 2, nums2, 5)
	log.Println(nums1)
}
