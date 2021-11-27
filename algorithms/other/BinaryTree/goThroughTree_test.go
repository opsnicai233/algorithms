package BinaryTree

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestGoDeepThroughTree(t *testing.T) {
	var node = &Node{
		Root:  5,
		Left:  nil,
		Right: nil,
	}

	node.SetLeft(&Node{2, nil, nil})
	node.SetRight(&Node{3, nil, nil})
	node.GetNodeLeft().SetLeft(&Node{4, nil, nil})
	node.GetNodeLeft().SetRight(&Node{7, nil, nil})

	GoDeepThroughTree(node)
	assert.Equal(t, arrDeep, []interface{}{5, 2, 4, null, null, 7, null, null, 3, null, null})

}

func TestGoWidthThroughTree(t *testing.T) {
	var node = &Node{
		Root:  5,
		Left:  nil,
		Right: nil,
	}
	node.SetLeft(&Node{2, nil, nil})
	node.SetRight(&Node{3, nil, nil})
	node.GetNodeLeft().SetLeft(&Node{4, nil, nil})
	node.GetNodeLeft().SetRight(&Node{7, nil, nil})

	assert.Equal(t, GoWidthThroughTree(node), []interface{}{
		5, 2, 3, 4, 7, null, null, null, null, null, null,
	})

	var n = &Node{5, nil, nil}
	n.SetRight(&Node{3, &Node{7, nil, nil}, nil})
	assert.Equal(t, GoWidthThroughTree(n), []interface{}{
		5, null, 3, 7, null, null, null,
	})
}
