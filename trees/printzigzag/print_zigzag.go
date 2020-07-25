package printzigzag

import (
	"fmt"

	"github.com/shomali11/go-interview/datastructures/linkedlists/doublylinkedlists"
	"github.com/shomali11/go-interview/datastructures/queues"
	"github.com/shomali11/go-interview/datastructures/trees"
)

// PrintZigZag prints a tree zig zag
func PrintZigZag(node *trees.MultiNode) {
	if node == nil {
		return
	}

	queue := queues.New()
	queue.Enqueue(node)

	list := doublylinkedlists.New()

	isForwardDirection := true
	for !queue.IsEmpty() {
		for size := queue.Size(); size > 0; size-- {
			element, _ := queue.Dequeue()
			node := element.(*trees.MultiNode)
			fmt.Print(node.Data, " ")

			for _, child := range node.Children {
				list.Add(child)
			}
		}

		if isForwardDirection {
			for _, value := range list.GetValues() {
				queue.Enqueue(value)
			}
		} else {
			for _, value := range list.GetReverseValues() {
				queue.Enqueue(value)
			}
		}

		list.Clear()
		isForwardDirection = !isForwardDirection
		fmt.Println()
	}
}
