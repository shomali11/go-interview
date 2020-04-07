package priorityqueues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Element struct {
	Value    interface{}
	Priority int
}

func TestMaxStringPriorityQueue(t *testing.T) {
	compare := func(i, j interface{}) bool {
		return i.(string) > j.(string)
	}

	priorityQueue := New(compare)
	assert.Equal(t, priorityQueue.IsEmpty(), true)

	priorityQueue.Push("hello")

	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 1)

	element, err := priorityQueue.Peek()
	assert.Equal(t, element, "hello")
	assert.Nil(t, err)

	priorityQueue.Remove(element)
	assert.Equal(t, priorityQueue.Size(), 0)
	assert.Nil(t, err)

	priorityQueue.Clear()
	assert.Equal(t, priorityQueue.IsEmpty(), true)
	assert.Equal(t, priorityQueue.Size(), 0)

	_, err = priorityQueue.Peek()
	assert.NotNil(t, err)

	priorityQueue.Push("abc")
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 1)

	priorityQueue.Push("sweet")
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 2)

	priorityQueue.Push("xyz")
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 3)

	element, err = priorityQueue.Peek()
	assert.Equal(t, element, "xyz")
	assert.Nil(t, err)

	element, err = priorityQueue.Pop()
	assert.Nil(t, err)
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 2)
	assert.Equal(t, element, "xyz")

	element, err = priorityQueue.Pop()
	assert.Nil(t, err)
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 1)
	assert.Equal(t, element, "sweet")

	element, err = priorityQueue.Pop()
	assert.Nil(t, err)
	assert.Equal(t, priorityQueue.IsEmpty(), true)
	assert.Equal(t, priorityQueue.Size(), 0)
	assert.Equal(t, element, "abc")

	_, err = priorityQueue.Pop()
	assert.NotNil(t, err)
}

func TestMaxPriorityQueue(t *testing.T) {
	compare := func(i, j interface{}) bool {
		return i.(*Element).Priority > j.(*Element).Priority
	}

	priorityQueue := New(compare)
	assert.Equal(t, priorityQueue.IsEmpty(), true)

	priorityQueue.Push(&Element{Value: "hello", Priority: 10})

	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 1)

	e, err := priorityQueue.Peek()
	element := e.(*Element)
	assert.Equal(t, element.Value, "hello")
	assert.Equal(t, element.Priority, 10)
	assert.Nil(t, err)

	priorityQueue.Remove(element)
	assert.Equal(t, priorityQueue.Size(), 0)
	assert.Nil(t, err)

	priorityQueue.Clear()
	assert.Equal(t, priorityQueue.IsEmpty(), true)
	assert.Equal(t, priorityQueue.Size(), 0)

	_, err = priorityQueue.Peek()
	assert.NotNil(t, err)

	priorityQueue.Push(&Element{Value: 5, Priority: 5})
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 1)

	priorityQueue.Push(&Element{Value: "sweet", Priority: 11})
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 2)

	priorityQueue.Push(&Element{Value: true, Priority: 1})
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 3)

	e, err = priorityQueue.Peek()
	element = e.(*Element)
	assert.Equal(t, element.Value, "sweet")
	assert.Equal(t, element.Priority, 11)
	assert.Nil(t, err)

	e, err = priorityQueue.Pop()
	element = e.(*Element)
	assert.Nil(t, err)
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 2)
	assert.Equal(t, element.Value, "sweet")
	assert.Equal(t, element.Priority, 11)

	e, err = priorityQueue.Pop()
	element = e.(*Element)
	assert.Nil(t, err)
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 1)
	assert.Equal(t, element.Value, 5)
	assert.Equal(t, element.Priority, 5)

	e, err = priorityQueue.Pop()
	element = e.(*Element)
	assert.Nil(t, err)
	assert.Equal(t, priorityQueue.IsEmpty(), true)
	assert.Equal(t, priorityQueue.Size(), 0)
	assert.Equal(t, element.Value, true)
	assert.Equal(t, element.Priority, 1)

	_, err = priorityQueue.Pop()
	assert.NotNil(t, err)
}

func TestMinPriorityQueue(t *testing.T) {
	compare := func(i, j interface{}) bool {
		return i.(*Element).Priority < j.(*Element).Priority
	}

	priorityQueue := New(compare)
	assert.Equal(t, priorityQueue.IsEmpty(), true)

	priorityQueue.Push(&Element{Value: "hello", Priority: 10})

	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 1)

	e, err := priorityQueue.Peek()
	element := e.(*Element)
	assert.Equal(t, element.Value, "hello")
	assert.Equal(t, element.Priority, 10)
	assert.Nil(t, err)

	priorityQueue.Remove(element)
	assert.Equal(t, priorityQueue.Size(), 0)
	assert.Nil(t, err)

	priorityQueue.Clear()
	assert.Equal(t, priorityQueue.IsEmpty(), true)
	assert.Equal(t, priorityQueue.Size(), 0)

	_, err = priorityQueue.Peek()
	assert.NotNil(t, err)

	priorityQueue.Push(&Element{Value: 5, Priority: 5})
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 1)

	priorityQueue.Push(&Element{Value: "sweet", Priority: 11})
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 2)

	priorityQueue.Push(&Element{Value: true, Priority: 1})
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 3)

	e, err = priorityQueue.Peek()
	element = e.(*Element)
	assert.Equal(t, element.Value, true)
	assert.Equal(t, element.Priority, 1)
	assert.Nil(t, err)

	e, err = priorityQueue.Pop()
	element = e.(*Element)
	assert.Nil(t, err)
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 2)
	assert.Equal(t, element.Value, true)
	assert.Equal(t, element.Priority, 1)

	e, err = priorityQueue.Pop()
	element = e.(*Element)
	assert.Nil(t, err)
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 1)
	assert.Equal(t, element.Value, 5)
	assert.Equal(t, element.Priority, 5)

	e, err = priorityQueue.Pop()
	element = e.(*Element)
	assert.Nil(t, err)
	assert.Equal(t, priorityQueue.IsEmpty(), true)
	assert.Equal(t, priorityQueue.Size(), 0)
	assert.Equal(t, element.Value, "sweet")
	assert.Equal(t, element.Priority, 11)

	_, err = priorityQueue.Pop()
	assert.NotNil(t, err)
}
