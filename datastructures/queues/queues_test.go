package queues

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := New()
	assert.Equal(t, queue.IsEmpty(), true)

	queue.Enqueue("hello")

	assert.True(t, reflect.DeepEqual(queue.GetValues(), []interface{}{"hello"}))
	assert.Equal(t, queue.IsEmpty(), false)
	assert.Equal(t, queue.Size(), 1)

	value, err := queue.Peek()
	assert.Equal(t, value, "hello")
	assert.Nil(t, err)

	queue.Clear()
	assert.True(t, reflect.DeepEqual(queue.GetValues(), []interface{}{}))
	assert.Equal(t, queue.IsEmpty(), true)
	assert.Equal(t, queue.Size(), 0)

	_, err = queue.Peek()
	assert.NotNil(t, err)

	queue.Enqueue(111)
	assert.True(t, reflect.DeepEqual(queue.GetValues(), []interface{}{111}))
	assert.Equal(t, queue.IsEmpty(), false)
	assert.Equal(t, queue.Size(), 1)

	queue.Enqueue(true)
	assert.True(t, reflect.DeepEqual(queue.GetValues(), []interface{}{111, true}))
	assert.Equal(t, queue.IsEmpty(), false)
	assert.Equal(t, queue.Size(), 2)

	value, err = queue.Dequeue()
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(queue.GetValues(), []interface{}{true}))
	assert.Equal(t, queue.IsEmpty(), false)
	assert.Equal(t, queue.Size(), 1)
	assert.Equal(t, value, 111)

	value, err = queue.Dequeue()
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(queue.GetValues(), []interface{}{}))
	assert.Equal(t, queue.IsEmpty(), true)
	assert.Equal(t, queue.Size(), 0)
	assert.Equal(t, value, true)

	_, err = queue.Dequeue()
	assert.NotNil(t, err)
}
