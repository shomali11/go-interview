package printcolumns

import (
	"testing"

	"github.com/shomali11/go-interview/datastructures/trees"
)

func TestPrintColumns(t *testing.T) {
	node4 := &trees.BinaryNode{Data: 5}
	node5 := &trees.BinaryNode{Data: 6}

	node6 := &trees.BinaryNode{Data: 7}
	node7 := &trees.BinaryNode{Data: 8}

	node1 := &trees.BinaryNode{Data: 2, Left: node4, Right: node5}
	node2 := &trees.BinaryNode{Data: 3, Left: node6, Right: node7}

	node := &trees.BinaryNode{Data: 1, Left: node1, Right: node2}
	PrintColumns(node)
}
