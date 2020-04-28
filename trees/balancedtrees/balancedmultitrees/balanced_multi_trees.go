package balancedmultitrees

import (
	"math"

	"github.com/shomali11/go-interview/datastructures/trees"
)

// IsBalanced checks if a multi tree is balanced
func IsBalanced(node *trees.MultiNode) bool {
	_, balanced := isBalanced(node)
	return balanced
}

func isBalanced(node *trees.MultiNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	if len(node.Children) == 0 {
		return 1, true
	}

	height, balanced := isBalanced(node.Children[0])
	if !balanced {
		return 0, false
	}

	min := height
	max := height

	for i := 1; i < len(node.Children); i++ {
		height, balanced := isBalanced(node.Children[i])
		if !balanced {
			return 0, false
		}

		if height < min {
			min = height
		}

		if height > max {
			max = height
		}
	}

	if math.Abs(float64(min-max)) > 1 {
		return 0, false
	}
	return 1 + max, true
}
