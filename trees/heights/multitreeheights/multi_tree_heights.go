package multitreeheights

import (
	"github.com/shomali11/go-interview/datastructures/trees"
)

// Height returns a tree's height
func Height(node *trees.MultiNode) int {
	if node == nil {
		return 0
	}

	if len(node.Children) == 0 {
		return 1
	}

	max := Height(node.Children[0])
	for i := 1; i < len(node.Children); i++ {
		current := Height(node.Children[i])
		if max < current {
			max = current
		}
	}
	return 1 + max
}
