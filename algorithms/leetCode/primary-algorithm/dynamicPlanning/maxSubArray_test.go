package dynamicPlanning

import (
	"log"
	"testing"
)

func TestMaxSubArray2(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	log.Println(MaxSubArray(nums))
}

func TestMaxSubArrayPro(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	log.Println(MaxSubArrayPro(nums))
}
