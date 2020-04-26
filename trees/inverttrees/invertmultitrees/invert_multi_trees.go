package invertmultitrees

import "github.com/shomali11/go-interview/datastructures/trees"

// InvertTree inverts a binary tree
func InvertTree(node *trees.MultiNode) *trees.MultiNode {
	if node == nil {
		return nil
	}

	children := node.Children
	for i, j := 0, len(children)-1; i < j; i, j = i+1, j-1 {
		children[i], children[j] = InvertTree(children[j]), InvertTree(children[i])
	}
	return node
}
