package sort

import (
	"log"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := []int{
		345, 12, 654, 6, 4, 32, 9, 56, 58, 9, 4, 21, 1, 0,
	}
	log.Println(InsertSort(arr))
}
