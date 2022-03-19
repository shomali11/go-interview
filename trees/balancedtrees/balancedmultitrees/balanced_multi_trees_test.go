package balancedmultitrees

import (
	"testing"

	"github.com/shomali11/go-interview/datastructures/trees"
	"github.com/stretchr/testify/assert"
)

func TestIsBalanced_EmptyNode(t *testing.T) {
	assert.True(t, IsBalanced[int](nil))
}

func TestIsBalanced_RootNode(t *testing.T) {
	node := &trees.MultiNode[int]{Data: 1}

	assert.True(t, IsBalanced(node))
}

func TestIsBalanced_HalfTree(t *testing.T) {
	node2 := &trees.MultiNode[int]{Data: 2}
	node := &trees.MultiNode[int]{Data: 1, Children: []*trees.MultiNode[int]{node2}}

	assert.True(t, IsBalanced(node))
}

func TestIsBalanced_FullTree(t *testing.T) {
	node2 := &trees.MultiNode[int]{Data: 2}
	node3 := &trees.MultiNode[int]{Data: 3}
	node := &trees.MultiNode[int]{Data: 1, Children: []*trees.MultiNode[int]{node2, node3}}

	assert.True(t, IsBalanced(node))
}

func TestIsBalanced_NotBalancedTree(t *testing.T) {
	node2 := &trees.MultiNode[int]{Data: 2}
	node3 := &trees.MultiNode[int]{Data: 3, Children: []*trees.MultiNode[int]{node2}}
	node4 := &trees.MultiNode[int]{Data: 4, Children: []*trees.MultiNode[int]{node3}}
	node5 := &trees.MultiNode[int]{Data: 5}
	node := &trees.MultiNode[int]{Data: 1, Children: []*trees.MultiNode[int]{node4, node5}}

	assert.False(t, IsBalanced(node))
}
