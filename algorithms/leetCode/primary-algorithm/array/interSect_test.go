package array

import (
	"log"
	"sort"
	"testing"
	"time"
)

func TestIntersect(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	Intersect(nums1, nums2)
}

func TestInterSectSorted(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	sort.Ints(nums1)
	sort.Ints(nums2)
	now := time.Now()
	log.Println(InterSectSorted(nums2, nums1))
	log.Println(time.Since(now))

}
