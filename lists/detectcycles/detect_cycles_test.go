package detectcycles

import (
	"testing"

	"github.com/shomali11/go-interview/datastructures/linkedlists/singlylinkedlists"
	"github.com/stretchr/testify/assert"
)

func TestContainsCycle(t *testing.T) {
	emptyList := singlylinkedlists.New()
	assert.False(t, ContainsCycle(emptyList.GetHead()))

	nonEmptyList := singlylinkedlists.New("A", "B", "C")
	assert.False(t, ContainsCycle(nonEmptyList.GetHead()))

	node1 := &singlylinkedlists.SLLNode{Value: "X"}
	node2 := &singlylinkedlists.SLLNode{Value: "X"}
	node3 := &singlylinkedlists.SLLNode{Value: "X"}
	node4 := &singlylinkedlists.SLLNode{Value: "X"}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2
	assert.True(t, ContainsCycle(node1))
}
