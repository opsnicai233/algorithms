package array

import (
	"log"
	"testing"
)

func TestPlusOne(t *testing.T) {
	arr := []int{9, 9, 9, 9}
	log.Println(PlusOne(arr))
	arr = []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 9, 9}
	log.Println(PlusOne(arr))
	arr = []int{0}
	log.Println(PlusOne(arr))
}
