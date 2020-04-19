package printcolumns

import (
	"fmt"

	"github.com/shomali11/go-interview/datastructures/maps/hashmultimaps"
	"github.com/shomali11/go-interview/datastructures/queues"
	"github.com/shomali11/go-interview/datastructures/trees"
)

// PrintColumns prints a tree column by column
func PrintColumns(node *trees.BinaryNode) {
	if node == nil {
		return
	}

	queue := queues.New()
	queue.Enqueue(node)

	multimap := hashmultimaps.New()
	multimap.Put(0, node)

	mapNodeIndex := make(map[*trees.BinaryNode]int)
	mapNodeIndex[node] = 0

	for !queue.IsEmpty() {
		element, _ := queue.Dequeue()
		node := element.(*trees.BinaryNode)
		nodeIndex := mapNodeIndex[node]

		if node.Left != nil {
			queue.Enqueue(node.Left)
			multimap.Put(nodeIndex-1, node.Left)
			mapNodeIndex[node.Left] = nodeIndex - 1
		}

		if node.Right != nil {
			queue.Enqueue(node.Right)
			multimap.Put(nodeIndex+1, node.Right)
			mapNodeIndex[node.Right] = nodeIndex + 1
		}
	}

	keys := multimap.GetKeys()
	minIndex, maxIndex := minMax(keys...)
	for i := minIndex; i <= maxIndex; i++ {
		for _, value := range multimap.GetValues(i) {
			fmt.Print(value.(*trees.BinaryNode).Data, " ")
		}
		fmt.Println()
	}
}

func minMax(values ...interface{}) (int, int) {
	first := values[0].(int)
	min := first
	max := first

	for i := 1; i < len(values); i++ {
		current := values[i].(int)
		if current < min {
			min = current
		}
		if current > max {
			max = current
		}
	}
	return min, max
}
