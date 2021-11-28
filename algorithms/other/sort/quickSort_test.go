package sort

import (
	"gopkg.in/go-playground/assert.v1"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	//arr := []int{}
	//log.Println(QuickSort(arr))
	//
	//arr = []int{2}
	//log.Println(QuickSort(arr))

	arr := []int{66, 7, 56, 4, 6, 7, 6, 9, 2, 97, 3, -33, 9, 10, 24}
	arrtest := QuickSort(arr)
	sort.Ints(arr)
	assert.Equal(t, arrtest, arr)
}
