package detectintersections

import (
	"testing"

	"github.com/shomali11/go-interview/datastructures/linkedlists/singlylinkedlists"
	"github.com/stretchr/testify/assert"
)

func TestDoIntersect(t *testing.T) {
	emptyList1 := singlylinkedlists.New()
	emptyList2 := singlylinkedlists.New()
	assert.False(t, DoIntersect(emptyList1.GetHead(), emptyList2.GetHead()))

	nonEmptyList1 := singlylinkedlists.New("A", "B", "C")
	nonEmptyList2 := singlylinkedlists.New("A", "B", "C")
	assert.False(t, DoIntersect(nonEmptyList1.GetHead(), nonEmptyList2.GetHead()))

	node1 := &singlylinkedlists.SLLNode{Value: "X"}
	node2 := &singlylinkedlists.SLLNode{Value: "X"}
	node3 := &singlylinkedlists.SLLNode{Value: "X"}
	node4 := &singlylinkedlists.SLLNode{Value: "X"}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	node5 := &singlylinkedlists.SLLNode{Value: "X"}
	node6 := &singlylinkedlists.SLLNode{Value: "X"}

	node5.Next = node6
	node6.Next = node4
	assert.True(t, DoIntersect(node1, node5))
}
