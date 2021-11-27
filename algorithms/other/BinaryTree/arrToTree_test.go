package BinaryTree

import (
	"log"
	"testing"
)

func TestArrayToBinaryTree(t *testing.T) {
	arr := []Data{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
	}
	node := ArrayToBinaryTree(arr)
	//log.Println(node.String())
	log.Println(GoWidthThroughTree(node))

	arr1 := []Data{
		"a", "b", "c", nil, "d", nil, "e", "f", nil,
	}
	node = ArrayToBinaryTree(arr1)
	log.Println(node.String())
	log.Println(GoWidthThroughTree(node))

	arr2 := []Data{
		"a",
	}
	node = ArrayToBinaryTree(arr2)
	log.Println(node.String())
	log.Println(GoWidthThroughTree(node))
}
