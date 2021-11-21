package array

import (
	"log"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	arr := []int{0,0,1,1,1,2,2,3,3,4}
	log.Println(RemoveDuplicates(arr))

	arr = []int{1,1,2}
	log.Println(RemoveDuplicates(arr))
}
