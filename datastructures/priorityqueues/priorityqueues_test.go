package priorityqueues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxPriorityQueue(t *testing.T) {
	priorityQueue := NewMax()
	assert.Equal(t, priorityQueue.IsEmpty(), true)

	priorityQueue.Push(&PQElement{Value: "hello", Priority: 10})

	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 1)

	element, err := priorityQueue.Peek()
	assert.Equal(t, element.Value, "hello")
	assert.Equal(t, element.Priority, 10)
	assert.Nil(t, err)

	priorityQueue.Clear()
	assert.Equal(t, priorityQueue.IsEmpty(), true)
	assert.Equal(t, priorityQueue.Size(), 0)

	_, err = priorityQueue.Peek()
	assert.NotNil(t, err)

	priorityQueue.Push(&PQElement{Value: 5, Priority: 5})
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 1)

	priorityQueue.Push(&PQElement{Value: "sweet", Priority: 11})
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 2)

	priorityQueue.Push(&PQElement{Value: true, Priority: 1})
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 3)

	element, err = priorityQueue.Peek()
	assert.Equal(t, element.Value, "sweet")
	assert.Equal(t, element.Priority, 11)
	assert.Nil(t, err)

	element, err = priorityQueue.Pop()
	assert.Nil(t, err)
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 2)
	assert.Equal(t, element.Value, "sweet")
	assert.Equal(t, element.Priority, 11)

	element, err = priorityQueue.Pop()
	assert.Nil(t, err)
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 1)
	assert.Equal(t, element.Value, 5)
	assert.Equal(t, element.Priority, 5)

	element, err = priorityQueue.Pop()
	assert.Nil(t, err)
	assert.Equal(t, priorityQueue.IsEmpty(), true)
	assert.Equal(t, priorityQueue.Size(), 0)
	assert.Equal(t, element.Value, true)
	assert.Equal(t, element.Priority, 1)

	_, err = priorityQueue.Pop()
	assert.NotNil(t, err)
}

func TestMinPriorityQueue(t *testing.T) {
	priorityQueue := NewMin()
	assert.Equal(t, priorityQueue.IsEmpty(), true)

	priorityQueue.Push(&PQElement{Value: "hello", Priority: 10})

	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 1)

	element, err := priorityQueue.Peek()
	assert.Equal(t, element.Value, "hello")
	assert.Equal(t, element.Priority, 10)
	assert.Nil(t, err)

	priorityQueue.Clear()
	assert.Equal(t, priorityQueue.IsEmpty(), true)
	assert.Equal(t, priorityQueue.Size(), 0)

	_, err = priorityQueue.Peek()
	assert.NotNil(t, err)

	priorityQueue.Push(&PQElement{Value: 5, Priority: 5})
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 1)

	priorityQueue.Push(&PQElement{Value: "sweet", Priority: 11})
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 2)

	priorityQueue.Push(&PQElement{Value: true, Priority: 1})
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 3)

	element, err = priorityQueue.Peek()
	assert.Equal(t, element.Value, true)
	assert.Equal(t, element.Priority, 1)
	assert.Nil(t, err)

	element, err = priorityQueue.Pop()
	assert.Nil(t, err)
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 2)
	assert.Equal(t, element.Value, true)
	assert.Equal(t, element.Priority, 1)

	element, err = priorityQueue.Pop()
	assert.Nil(t, err)
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 1)
	assert.Equal(t, element.Value, 5)
	assert.Equal(t, element.Priority, 5)

	element, err = priorityQueue.Pop()
	assert.Nil(t, err)
	assert.Equal(t, priorityQueue.IsEmpty(), true)
	assert.Equal(t, priorityQueue.Size(), 0)
	assert.Equal(t, element.Value, "sweet")
	assert.Equal(t, element.Priority, 11)

	_, err = priorityQueue.Pop()
	assert.NotNil(t, err)
}
