package array

import (
	"log"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	arr := []int{0, 0, 1}
	MoveZeroes(arr)
	log.Println(arr)
}

func TestMoveZeroesPro(t *testing.T) {
	arr := []int{0, 0, 1}
	MoveZeroesPro(arr)
	arr = []int{0, 0, 1, 0, 2, 3, 0, 0}
	MoveZeroesPro(arr)
	log.Println(arr)
}
