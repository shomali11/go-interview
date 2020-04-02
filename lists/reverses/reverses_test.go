package reverses

import (
	"testing"

	"github.com/shomali11/go-interview/datastructures/linkedlists/singlylinkedlists"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	node1 := &singlylinkedlists.SLLNode{Value: "1"}
	node2 := &singlylinkedlists.SLLNode{Value: "2"}
	node3 := &singlylinkedlists.SLLNode{Value: "3"}

	node1.Next = node2
	node2.Next = node3
	node3.Next = nil

	newHead := Reverse(node1)

	assert.Equal(t, newHead.Value, "3")
	assert.NotNil(t, newHead.Next)
	assert.Equal(t, newHead.Next.Value, "2")
	assert.NotNil(t, newHead.Next.Next)
	assert.Equal(t, newHead.Next.Next.Value, "1")
	assert.Nil(t, newHead.Next.Next.Next)
}
