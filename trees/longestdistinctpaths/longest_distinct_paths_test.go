package longestdistinctpaths

import (
	"testing"

	"github.com/shomali11/go-interview/datastructures/trees"
	"github.com/stretchr/testify/assert"
)

func TestLongestDistinctPath1(t *testing.T) {
	node4 := &trees.MultiNode{Data: 5, Children: make([]*trees.MultiNode, 0)}
	node5 := &trees.MultiNode{Data: 6, Children: make([]*trees.MultiNode, 0)}

	node6 := &trees.MultiNode{Data: 7, Children: make([]*trees.MultiNode, 0)}
	node7 := &trees.MultiNode{Data: 8, Children: make([]*trees.MultiNode, 0)}

	node1 := &trees.MultiNode{Data: 2, Children: []*trees.MultiNode{node4, node5}}
	node2 := &trees.MultiNode{Data: 3, Children: []*trees.MultiNode{node6, node7}}
	node3 := &trees.MultiNode{Data: 4, Children: make([]*trees.MultiNode, 0)}

	node := &trees.MultiNode{Data: 1, Children: []*trees.MultiNode{node1, node2, node3}}
	assert.Equal(t, LongestDistinctPath(node), 3)
}

func TestLongestDistinctPath2(t *testing.T) {
	node4 := &trees.MultiNode{Data: 1, Children: make([]*trees.MultiNode, 0)}
	node5 := &trees.MultiNode{Data: 1, Children: make([]*trees.MultiNode, 0)}

	node6 := &trees.MultiNode{Data: 1, Children: make([]*trees.MultiNode, 0)}
	node7 := &trees.MultiNode{Data: 1, Children: make([]*trees.MultiNode, 0)}

	node1 := &trees.MultiNode{Data: 2, Children: []*trees.MultiNode{node4, node5}}
	node2 := &trees.MultiNode{Data: 3, Children: []*trees.MultiNode{node6, node7}}
	node3 := &trees.MultiNode{Data: 4, Children: make([]*trees.MultiNode, 0)}

	node := &trees.MultiNode{Data: 1, Children: []*trees.MultiNode{node1, node2, node3}}
	assert.Equal(t, LongestDistinctPath(node), 2)
}

func TestLongestDistinctPath3(t *testing.T) {
	node4 := &trees.MultiNode{Data: 14, Children: make([]*trees.MultiNode, 0)}
	node5 := &trees.MultiNode{Data: 13, Children: make([]*trees.MultiNode, 0)}

	node6 := &trees.MultiNode{Data: 12, Children: make([]*trees.MultiNode, 0)}
	node7 := &trees.MultiNode{Data: 10, Children: make([]*trees.MultiNode, 0)}

	node1 := &trees.MultiNode{Data: 1, Children: []*trees.MultiNode{node4, node5}}
	node2 := &trees.MultiNode{Data: 1, Children: []*trees.MultiNode{node6, node7}}
	node3 := &trees.MultiNode{Data: 1, Children: make([]*trees.MultiNode, 0)}

	node := &trees.MultiNode{Data: 1, Children: []*trees.MultiNode{node1, node2, node3}}
	assert.Equal(t, LongestDistinctPath(node), 1)
}

func TestLongestDistinctPath4(t *testing.T) {
	assert.Equal(t, LongestDistinctPath(nil), 0)
}
