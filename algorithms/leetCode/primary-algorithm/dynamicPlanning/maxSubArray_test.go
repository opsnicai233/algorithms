package dynamicPlanning

import (
	"log"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	log.Println(MaxSubArray(nums))
}
