package longestdistinctpaths

import (
	"github.com/shomali11/go-interview/datastructures/sets/hashmultisets"
	"github.com/shomali11/go-interview/datastructures/trees"
)

// LongestDistinctPath returns the length of the longest distinct path
func LongestDistinctPath(node *trees.MultiNode) int {
	multiSet := hashmultisets.New()
	return longestDistinctPath(node, multiSet)
}

func longestDistinctPath(node *trees.MultiNode, multiSet *hashmultisets.HashMultiSet) int {
	if node == nil || multiSet.Contains(node.Data) {
		return multiSet.Size()
	}

	multiSet.Add(node.Data)

	maxPath := 0
	for _, child := range node.Children {
		length := longestDistinctPath(child, multiSet)
		if length > maxPath {
			maxPath = length
		}
	}

	if len(node.Children) == 0 {
		maxPath = multiSet.Size()
	}

	multiSet.IncrementBy(node.Data, -1)
	if multiSet.GetCount(node.Data) == 0 {
		multiSet.Remove(node.Data)
	}
	return maxPath
}
