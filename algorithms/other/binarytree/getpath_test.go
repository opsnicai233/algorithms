package binarytree

import (
	"log"
	"testing"
)

func TestGoDeep(t *testing.T) {
	arr := []int{5, 3, 4, null, 1, 2, null, 7, null}
	log.Println(GoDeep(0, 0, 11, arr))
	arr = []int{5}
	log.Println(GoDeep(0, 0, 11, arr))
}
