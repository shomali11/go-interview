package invertmultitrees

import (
	"testing"

	"github.com/shomali11/go-interview/datastructures/trees"
	"github.com/stretchr/testify/assert"
)

func TestInvertTrees(t *testing.T) {
	node2 := &trees.MultiNode{Data: 2}
	node3 := &trees.MultiNode{Data: 3}
	node := &trees.MultiNode{Data: 1, Children: []*trees.MultiNode{node2, node3}}

	node = InvertTree(node)

	assert.Equal(t, node.Children[0].Data, 3)
	assert.Equal(t, node.Children[1].Data, 2)
}
