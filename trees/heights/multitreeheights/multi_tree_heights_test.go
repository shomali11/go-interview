package multitreeheights

import (
	"testing"

	"github.com/shomali11/go-interview/datastructures/trees"
	"github.com/stretchr/testify/assert"
)

func TestHeight_EmptyNode(t *testing.T) {
	assert.Equal(t, Height(nil), 0)
}

func TestHeight_RootNode(t *testing.T) {
	node := &trees.MultiNode{Data: 1}

	assert.Equal(t, Height(node), 1)
}

func TestHeight_HalfTree(t *testing.T) {
	node2 := &trees.MultiNode{Data: 2}
	node := &trees.MultiNode{Data: 1, Children: []*trees.MultiNode{node2}}

	assert.Equal(t, Height(node), 2)
}

func TestHeight_FullTree(t *testing.T) {
	node2 := &trees.MultiNode{Data: 2}
	node3 := &trees.MultiNode{Data: 3}
	node := &trees.MultiNode{Data: 1, Children: []*trees.MultiNode{node2, node3}}

	assert.Equal(t, Height(node), 2)
}

func TestHeight_NotBalancedTree(t *testing.T) {
	node2 := &trees.MultiNode{Data: 2}
	node3 := &trees.MultiNode{Data: 3, Children: []*trees.MultiNode{node2}}
	node4 := &trees.MultiNode{Data: 4, Children: []*trees.MultiNode{node3}}
	node5 := &trees.MultiNode{Data: 5}
	node := &trees.MultiNode{Data: 1, Children: []*trees.MultiNode{node4, node5}}

	assert.Equal(t, Height(node), 4)
}
