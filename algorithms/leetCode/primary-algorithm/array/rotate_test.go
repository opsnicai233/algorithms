package array

import (
	"log"
	"testing"
)

func TestRotate(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	Rotate(arr, 3)
	log.Println(arr)
}

func TestRotate2(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateProMax(arr, 3)
	log.Println(arr)
}
