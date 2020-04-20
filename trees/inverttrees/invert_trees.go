package inverttrees

import "github.com/shomali11/go-interview/datastructures/trees"

// InvertTree inverts a binary tree
func InvertTree(node *trees.BinaryNode) *trees.BinaryNode {
	if node == nil {
		return nil
	}

	node.Left = InvertTree(node.Right)
	node.Right = InvertTree(node.Left)
	return node
}
