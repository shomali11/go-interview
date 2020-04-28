package binarytreeheights

import (
	"testing"

	"github.com/shomali11/go-interview/datastructures/trees"
	"github.com/stretchr/testify/assert"
)

func TestHeight_EmptyNode(t *testing.T) {
	assert.Equal(t, Height(nil), 0)
}

func TestHeight_RootNode(t *testing.T) {
	node := &trees.BinaryNode{Data: 1}

	assert.Equal(t, Height(node), 1)
}

func TestHeight_HalfTree(t *testing.T) {
	node2 := &trees.BinaryNode{Data: 2}
	node := &trees.BinaryNode{Data: 1, Left: node2}

	assert.Equal(t, Height(node), 2)
}

func TestHeight_FullTree(t *testing.T) {
	node2 := &trees.BinaryNode{Data: 2}
	node3 := &trees.BinaryNode{Data: 3}
	node := &trees.BinaryNode{Data: 1, Left: node2, Right: node3}

	assert.Equal(t, Height(node), 2)
}

func TestHeight_LongerTree(t *testing.T) {
	node2 := &trees.BinaryNode{Data: 2}
	node3 := &trees.BinaryNode{Data: 3, Left: node2}
	node := &trees.BinaryNode{Data: 1, Right: node3}

	assert.Equal(t, Height(node), 3)
}
