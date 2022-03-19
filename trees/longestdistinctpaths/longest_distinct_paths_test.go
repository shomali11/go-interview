package longestdistinctpaths

import (
	"testing"

	"github.com/shomali11/go-interview/datastructures/trees"
	"github.com/stretchr/testify/assert"
)

func TestLongestDistinctPath1(t *testing.T) {
	node4 := &trees.MultiNode[int]{Data: 5, Children: make([]*trees.MultiNode[int], 0)}
	node5 := &trees.MultiNode[int]{Data: 6, Children: make([]*trees.MultiNode[int], 0)}

	node6 := &trees.MultiNode[int]{Data: 7, Children: make([]*trees.MultiNode[int], 0)}
	node7 := &trees.MultiNode[int]{Data: 8, Children: make([]*trees.MultiNode[int], 0)}

	node1 := &trees.MultiNode[int]{Data: 2, Children: []*trees.MultiNode[int]{node4, node5}}
	node2 := &trees.MultiNode[int]{Data: 3, Children: []*trees.MultiNode[int]{node6, node7}}
	node3 := &trees.MultiNode[int]{Data: 4, Children: make([]*trees.MultiNode[int], 0)}

	node := &trees.MultiNode[int]{Data: 1, Children: []*trees.MultiNode[int]{node1, node2, node3}}
	assert.Equal(t, LongestDistinctPath(node), 3)
}

func TestLongestDistinctPath2(t *testing.T) {
	node4 := &trees.MultiNode[int]{Data: 1, Children: make([]*trees.MultiNode[int], 0)}
	node5 := &trees.MultiNode[int]{Data: 1, Children: make([]*trees.MultiNode[int], 0)}

	node6 := &trees.MultiNode[int]{Data: 1, Children: make([]*trees.MultiNode[int], 0)}
	node7 := &trees.MultiNode[int]{Data: 1, Children: make([]*trees.MultiNode[int], 0)}

	node1 := &trees.MultiNode[int]{Data: 2, Children: []*trees.MultiNode[int]{node4, node5}}
	node2 := &trees.MultiNode[int]{Data: 3, Children: []*trees.MultiNode[int]{node6, node7}}
	node3 := &trees.MultiNode[int]{Data: 4, Children: make([]*trees.MultiNode[int], 0)}

	node := &trees.MultiNode[int]{Data: 1, Children: []*trees.MultiNode[int]{node1, node2, node3}}
	assert.Equal(t, LongestDistinctPath(node), 2)
}

func TestLongestDistinctPath3(t *testing.T) {
	node4 := &trees.MultiNode[int]{Data: 14, Children: make([]*trees.MultiNode[int], 0)}
	node5 := &trees.MultiNode[int]{Data: 13, Children: make([]*trees.MultiNode[int], 0)}

	node6 := &trees.MultiNode[int]{Data: 12, Children: make([]*trees.MultiNode[int], 0)}
	node7 := &trees.MultiNode[int]{Data: 10, Children: make([]*trees.MultiNode[int], 0)}

	node1 := &trees.MultiNode[int]{Data: 1, Children: []*trees.MultiNode[int]{node4, node5}}
	node2 := &trees.MultiNode[int]{Data: 1, Children: []*trees.MultiNode[int]{node6, node7}}
	node3 := &trees.MultiNode[int]{Data: 1, Children: make([]*trees.MultiNode[int], 0)}

	node := &trees.MultiNode[int]{Data: 1, Children: []*trees.MultiNode[int]{node1, node2, node3}}
	assert.Equal(t, LongestDistinctPath(node), 1)
}

func TestLongestDistinctPath4(t *testing.T) {
	assert.Equal(t, LongestDistinctPath[int](nil), 0)
}
