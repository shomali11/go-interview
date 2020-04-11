package priorityqueues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxStringPriorityQueue_IsEmpty(t *testing.T) {
	compare := func(i, j interface{}) bool {
		return i.(string) > j.(string)
	}

	priorityQueue := New(compare)
	assert.Equal(t, priorityQueue.IsEmpty(), true)
	assert.Equal(t, priorityQueue.Size(), 0)

	priorityQueue.Push("hello")
	assert.Equal(t, priorityQueue.IsEmpty(), false)
	assert.Equal(t, priorityQueue.Size(), 1)
}

func TestMaxStringPriorityQueue_Peek(t *testing.T) {
	compare := func(i, j interface{}) bool {
		return i.(string) > j.(string)
	}

	priorityQueue := New(compare)
	_, err := priorityQueue.Peek()
	assert.NotNil(t, err)

	priorityQueue.Push("hello")
	element, err := priorityQueue.Peek()
	assert.Equal(t, element, "hello")
	assert.Nil(t, err)

	priorityQueue.Remove(element)
	assert.Equal(t, priorityQueue.Size(), 0)

	_, err = priorityQueue.Peek()
	assert.NotNil(t, err)
}

func TestMaxStringPriorityQueue_Clear(t *testing.T) {
	compare := func(i, j interface{}) bool {
		return i.(string) > j.(string)
	}

	priorityQueue := New(compare)
	assert.Equal(t, priorityQueue.IsEmpty(), true)
	assert.Equal(t, priorityQueue.Size(), 0)

	priorityQueue.Push("hello")

	priorityQueue.Clear()
	assert.Equal(t, priorityQueue.IsEmpty(), true)
	assert.Equal(t, priorityQueue.Size(), 0)
}

func TestMaxStringPriorityQueue_Pop(t *testing.T) {
	compare := func(i, j interface{}) bool {
		return i.(string) > j.(string)
	}

	priorityQueue := New(compare)
	priorityQueue.Push("abc")
	priorityQueue.Push("sweet", "xyz")

	element, err := priorityQueue.Pop()
	assert.Nil(t, err)
	assert.Equal(t, priorityQueue.Size(), 2)
	assert.Equal(t, element, "xyz")

	element, err = priorityQueue.Pop()
	assert.Nil(t, err)
	assert.Equal(t, priorityQueue.Size(), 1)
	assert.Equal(t, element, "sweet")

	element, err = priorityQueue.Pop()
	assert.Nil(t, err)
	assert.Equal(t, priorityQueue.Size(), 0)
	assert.Equal(t, element, "abc")

	_, err = priorityQueue.Pop()
	assert.NotNil(t, err)
}
