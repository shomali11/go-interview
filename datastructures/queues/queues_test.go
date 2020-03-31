package queues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := New()
	assert.Equal(t, queue.IsEmpty(), true)

	queue.Enqueue("hello")

	assert.Equal(t, queue.IsEmpty(), false)
	assert.Equal(t, queue.Size(), 1)

	value, err := queue.Peek()
	assert.Equal(t, value, "hello")
	assert.Nil(t, err)

	queue.Clear()
	assert.Equal(t, queue.IsEmpty(), true)
	assert.Equal(t, queue.Size(), 0)

	_, err = queue.Peek()
	assert.NotNil(t, err)

	queue.Enqueue(111)
	assert.Equal(t, queue.IsEmpty(), false)
	assert.Equal(t, queue.Size(), 1)

	queue.Enqueue(true)
	assert.Equal(t, queue.IsEmpty(), false)
	assert.Equal(t, queue.Size(), 2)

	value, err = queue.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, queue.IsEmpty(), false)
	assert.Equal(t, queue.Size(), 1)
	assert.Equal(t, value, 111)

	value, err = queue.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, queue.IsEmpty(), true)
	assert.Equal(t, queue.Size(), 0)
	assert.Equal(t, value, true)

	_, err = queue.Dequeue()
	assert.NotNil(t, err)
}
