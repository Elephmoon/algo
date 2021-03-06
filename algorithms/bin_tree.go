package algorithms

import (
	"errors"
	"fmt"
)

var (
	errNilTree = errors.New("tree is nill")
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) Insert(value int) error {
	if n == nil {
		return errNilTree
	}
	if n.value > value {
		if n.left == nil {
			n.left = &Node{value: value}
			return nil
		}
		return n.left.Insert(value)
	}

	if n.value < value {
		if n.right == nil {
			n.right = &Node{value: value}
		}
		return n.right.Insert(value)
	}

	return nil
}

func (n *Node) Min() int {
	if n.left == nil {
		return n.value
	}
	return n.left.Min()
}

func (n *Node) Max() int {
	if n.right == nil {
		return n.value
	}
	return n.right.Max()
}

func (n *Node) Inorder() {
	if n == nil {
		return
	}
	n.left.Inorder()
	fmt.Println(n.value)
	n.right.Inorder()
}
