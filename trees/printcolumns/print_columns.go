package printcolumns

import (
	"fmt"

	"github.com/shomali11/go-interview/datastructures/maps/hashmultimaps"
	"github.com/shomali11/go-interview/datastructures/queues"
	"github.com/shomali11/go-interview/datastructures/trees"

	"golang.org/x/exp/constraints"
)

// PrintColumns prints a tree column by column
func PrintColumns[T constraints.Ordered](node *trees.BinaryNode[T]) {
	if node == nil {
		return
	}

	queue := queues.New[*trees.BinaryNode[T]]()
	queue.Enqueue(node)

	multimap := hashmultimaps.New[int, *trees.BinaryNode[T]]()
	multimap.Put(0, node)

	mapNodeIndex := make(map[*trees.BinaryNode[T]]int)
	mapNodeIndex[node] = 0

	for !queue.IsEmpty() {
		element, _ := queue.Dequeue()
		nodeIndex := mapNodeIndex[element]

		if element.Left != nil {
			queue.Enqueue(element.Left)
			multimap.Put(nodeIndex-1, element.Left)
			mapNodeIndex[element.Left] = nodeIndex - 1
		}

		if element.Right != nil {
			queue.Enqueue(element.Right)
			multimap.Put(nodeIndex+1, element.Right)
			mapNodeIndex[element.Right] = nodeIndex + 1
		}
	}

	keys := multimap.GetKeys()
	minIndex, maxIndex := minMax(keys...)
	for i := minIndex; i <= maxIndex; i++ {
		for _, value := range multimap.GetValues(i) {
			fmt.Print(value.Data, " ")
		}
		fmt.Println()
	}
}

func minMax[T constraints.Ordered](values ...T) (T, T) {
	first := values[0]
	min := first
	max := first

	for i := 1; i < len(values); i++ {
		current := values[i]
		if current < min {
			min = current
		}
		if current > max {
			max = current
		}
	}
	return min, max
}
