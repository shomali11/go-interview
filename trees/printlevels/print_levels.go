package printlevels

import (
	"fmt"

	"github.com/shomali11/go-interview/datastructures/queues"
	"github.com/shomali11/go-interview/datastructures/trees"
)

// PrintLevels prints a tree level by level
func PrintLevels(node *trees.MultiNode) {
	if node == nil {
		return
	}

	queue := queues.New()
	queue.Enqueue(node)

	for !queue.IsEmpty() {
		for size := queue.Size(); size > 0; size-- {
			element, _ := queue.Dequeue()
			node := element.(*trees.MultiNode)
			fmt.Print(node.Data, " ")

			for _, child := range node.Children {
				queue.Enqueue(child)
			}
		}
		fmt.Println()
	}
}
