package binarytreeheights

import (
	"github.com/shomali11/go-interview/datastructures/trees"
)

// Height returns a tree's height
func Height[T any](node *trees.BinaryNode[T]) int {
	if node == nil {
		return 0
	}

	left := Height(node.Left)
	right := Height(node.Right)

	if left > right {
		return 1 + left
	}
	return 1 + right
}
