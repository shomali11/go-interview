package invertbinarytrees

import "github.com/shomali11/go-interview/datastructures/trees"

// InvertTree inverts a binary tree
func InvertTree[T any](node *trees.BinaryNode[T]) *trees.BinaryNode[T] {
	if node == nil {
		return nil
	}

	node.Left, node.Right = InvertTree(node.Right), InvertTree(node.Left)
	return node
}
