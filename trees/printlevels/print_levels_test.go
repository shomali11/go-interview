package printlevels

import (
	"testing"

	"github.com/shomali11/go-interview/datastructures/trees"
)

func TestPrintLevels(t *testing.T) {
	node4 := &trees.MultiNode{Data: 5, Children: make([]*trees.MultiNode, 0)}
	node5 := &trees.MultiNode{Data: 6, Children: make([]*trees.MultiNode, 0)}

	node6 := &trees.MultiNode{Data: 7, Children: make([]*trees.MultiNode, 0)}
	node7 := &trees.MultiNode{Data: 8, Children: make([]*trees.MultiNode, 0)}

	node1 := &trees.MultiNode{Data: 2, Children: []*trees.MultiNode{node4, node5}}
	node2 := &trees.MultiNode{Data: 3, Children: []*trees.MultiNode{node6, node7}}
	node3 := &trees.MultiNode{Data: 4, Children: make([]*trees.MultiNode, 0)}

	node := &trees.MultiNode{Data: 1, Children: []*trees.MultiNode{node1, node2, node3}}
	PrintLevels(node)
}
