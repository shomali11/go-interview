package balancedbinarytrees

import (
	"math"

	"github.com/shomali11/go-interview/datastructures/trees"
)

// IsBalanced checks if a binary tree is balanced
func IsBalanced(node *trees.BinaryNode) bool {
	_, balanced := isBalanced(node)
	return balanced
}

func isBalanced(node *trees.BinaryNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	left, balanced := isBalanced(node.Left)
	if !balanced {
		return 0, false
	}

	right, balanced := isBalanced(node.Right)
	if !balanced {
		return 0, false
	}

	if math.Abs(float64(left-right)) > 1 {
		return 0, false
	}

	if left > right {
		return 1 + left, true
	}
	return 1 + right, true
}
