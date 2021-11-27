package BinaryTree

import (
	"log"
	"testing"
)

func TestGoDeep(t *testing.T) {
	arr := []int{5, 3, 4, Null, 1, 2, Null, 7, Null}
	log.Println(GoDeep(0, 0, 11, arr))
	arr = []int{5}
	log.Println(GoDeep(0, 0, 11, arr))
}
