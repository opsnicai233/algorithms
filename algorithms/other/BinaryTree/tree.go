package BinaryTree

import (
	"fmt"
)

type Data interface{}

var null struct{}

// Node 二叉树
type Node struct {
	Root  Data
	Left  *Node
	Right *Node
}

func (node *Node) String() string {
	if node.IsEmpty() {
		return fmt.Sprintf("%s", null)
	}
	return fmt.Sprintf("{Root: %v, Left: %v, Right: %v}", node.Root, node.Left, node.Right)
}

func (node *Node) GetData() Data {
	if node.IsEmpty() {
		return null
	}
	return node.Root
}

func (node *Node) GetNodeLeft() *Node {
	return node.Left
}

func (node *Node) GetNodeRight() *Node {
	return node.Right
}

func (node *Node) SetData(data Data) {
	node.Root = data
}

func (node *Node) SetLeft(n *Node) {
	node.Left = n
}

func (node *Node) SetRight(n *Node) {
	node.Right = n
}

func (node *Node) IsEmpty() bool {
	return node == nil || node.Root == nil
}
