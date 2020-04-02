package detectcycles

import "github.com/shomali11/go-interview/datastructures/linkedlists/singlylinkedlists"

// ContainsCycle checks if the list contains a cycle
func ContainsCycle(head *singlylinkedlists.SLLNode) bool {
	fastPointer := head
	slowPointer := head
	for fastPointer != nil && fastPointer.Next != nil {
		fastPointer = fastPointer.Next.Next
		slowPointer = slowPointer.Next
		if slowPointer == fastPointer {
			return true
		}
	}
	return false
}
