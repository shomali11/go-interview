package queues

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_Clear(t *testing.T) {
	queue := New()
	assert.Equal(t, queue.Size(), 0)
	assert.Equal(t, queue.IsEmpty(), true)

	queue.Enqueue("hello")
	assert.Equal(t, queue.IsEmpty(), false)
	assert.Equal(t, queue.Size(), 1)

	queue.Clear()
	assert.Equal(t, queue.IsEmpty(), true)
	assert.Equal(t, queue.Size(), 0)
}

func TestQueue_GetValues(t *testing.T) {
	queue := New()

	queue.Enqueue("hello", "abc", "xyz")
	assert.True(t, reflect.DeepEqual(queue.GetValues(), []interface{}{"hello", "abc", "xyz"}))
}

func TestQueue_Peek(t *testing.T) {
	queue := New()

	queue.Enqueue("hello")
	value, err := queue.Peek()
	assert.Equal(t, value, "hello")
	assert.Nil(t, err)

	queue.Enqueue("abc")
	value, err = queue.Peek()
	assert.Equal(t, value, "hello")
	assert.Nil(t, err)
}

func TestQueue_Dequeue(t *testing.T) {
	queue := New()

	queue.Enqueue(111, true)

	value, err := queue.Dequeue()
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
