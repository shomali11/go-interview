package balancedmultitrees

import (
	"testing"

	"github.com/shomali11/go-interview/datastructures/trees"
	"github.com/stretchr/testify/assert"
)

func TestIsBalanced_EmptyNode(t *testing.T) {
	assert.True(t, IsBalanced(nil))
}

func TestIsBalanced_RootNode(t *testing.T) {
	node := &trees.MultiNode{Data: 1}

	assert.True(t, IsBalanced(node))
}

func TestIsBalanced_HalfTree(t *testing.T) {
	node2 := &trees.MultiNode{Data: 2}
	node := &trees.MultiNode{Data: 1, Children: []*trees.MultiNode{node2}}

	assert.True(t, IsBalanced(node))
}

func TestIsBalanced_FullTree(t *testing.T) {
	node2 := &trees.MultiNode{Data: 2}
	node3 := &trees.MultiNode{Data: 3}
	node := &trees.MultiNode{Data: 1, Children: []*trees.MultiNode{node2, node3}}

	assert.True(t, IsBalanced(node))
}

func TestIsBalanced_NotBalancedTree(t *testing.T) {
	node2 := &trees.MultiNode{Data: 2}
	node3 := &trees.MultiNode{Data: 3, Children: []*trees.MultiNode{node2}}
	node4 := &trees.MultiNode{Data: 4, Children: []*trees.MultiNode{node3}}
	node5 := &trees.MultiNode{Data: 5}
	node := &trees.MultiNode{Data: 1, Children: []*trees.MultiNode{node4, node5}}

	assert.False(t, IsBalanced(node))
}
