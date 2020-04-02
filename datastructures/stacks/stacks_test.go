package stacks

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := New()
	assert.Equal(t, stack.IsEmpty(), true)

	stack.Push("hello")

	assert.True(t, reflect.DeepEqual(stack.GetValues(), []interface{}{"hello"}))
	assert.Equal(t, stack.IsEmpty(), false)
	assert.Equal(t, stack.Size(), 1)

	value, err := stack.Peek()
	assert.Equal(t, value, "hello")
	assert.Nil(t, err)

	stack.Clear()
	assert.True(t, reflect.DeepEqual(stack.GetValues(), []interface{}{}))
	assert.Equal(t, stack.IsEmpty(), true)
	assert.Equal(t, stack.Size(), 0)

	_, err = stack.Peek()
	assert.NotNil(t, err)

	stack.Push(111)
	assert.True(t, reflect.DeepEqual(stack.GetValues(), []interface{}{111}))
	assert.Equal(t, stack.IsEmpty(), false)
	assert.Equal(t, stack.Size(), 1)

	stack.Push(true)
	assert.True(t, reflect.DeepEqual(stack.GetValues(), []interface{}{111, true}))
	assert.Equal(t, stack.IsEmpty(), false)
	assert.Equal(t, stack.Size(), 2)

	value, err = stack.Pop()
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(stack.GetValues(), []interface{}{111}))
	assert.Equal(t, stack.IsEmpty(), false)
	assert.Equal(t, stack.Size(), 1)
	assert.Equal(t, value, true)

	value, err = stack.Pop()
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(stack.GetValues(), []interface{}{}))
	assert.Equal(t, stack.IsEmpty(), true)
	assert.Equal(t, stack.Size(), 0)
	assert.Equal(t, value, 111)

	_, err = stack.Pop()
	assert.NotNil(t, err)
}
