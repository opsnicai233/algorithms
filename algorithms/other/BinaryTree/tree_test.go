package BinaryTree

import (
	"gopkg.in/go-playground/assert.v1"
	"log"
	"testing"
)

var node = &Node{
	Root:  5,
	Left:  nil,
	Right: nil,
}

func TestNode_GetNodeLeft(t *testing.T) {
	node.SetData(55)
	log.Println(node.String())
	log.Println(node.GetNodeRight().String())
	data := node.GetData()
	assert.Equal(t, data, 55)

	n1 := node.GetNodeLeft()

	assert.Equal(t, n1, nil)

	node.SetLeft(&Node{
		Root:  12,
		Left:  nil,
		Right: nil,
	})

	n1 = node.GetNodeLeft()
	assert.Equal(t, n1, &Node{12, nil, nil})

	assert.Equal(t, new(Data), nil)
	assert.Equal(t, node.IsEmpty(), false)

	assert.Equal(t, n1.GetNodeLeft().IsEmpty(), true)

}
