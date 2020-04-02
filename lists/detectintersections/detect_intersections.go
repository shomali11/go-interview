package detectintersections

import "github.com/shomali11/go-interview/datastructures/linkedlists/singlylinkedlists"

// DoIntersect checks if the two lists intersect
func DoIntersect(head1 *singlylinkedlists.SLLNode, head2 *singlylinkedlists.SLLNode) bool {
	head1ListCount := getCountNodes(head1)
	head2ListCount := getCountNodes(head2)

	current1 := getStartingNode(head1, head1ListCount-head2ListCount)
	current2 := getStartingNode(head2, head2ListCount-head1ListCount)

	for current1 != nil {
		if current1 == current2 {
			return true
		}
		current1 = current1.Next
		current2 = current2.Next
	}
	return false
}

func getStartingNode(head *singlylinkedlists.SLLNode, moves int) *singlylinkedlists.SLLNode {
	current := head
	for moves > 0 {
		current = current.Next
		moves--
	}
	return current
}

func getCountNodes(head *singlylinkedlists.SLLNode) int {
	count := 0
	current := head
	for current != nil {
		current = current.Next
		count++
	}
	return count
}
