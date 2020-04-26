package linkedliststacks

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Clear(t *testing.T) {
	stack := New()
	assert.Equal(t, stack.Size(), 0)
	assert.Equal(t, stack.IsEmpty(), true)

	stack.Push("hello")
	assert.Equal(t, stack.IsEmpty(), false)
	assert.Equal(t, stack.Size(), 1)

	stack.Clear()
	assert.Equal(t, stack.IsEmpty(), true)
	assert.Equal(t, stack.Size(), 0)
}

func TestStack_GetValues(t *testing.T) {
	stack := New()

	stack.Push("hello", "abc", "xyz")
	assert.True(t, reflect.DeepEqual(stack.GetValues(), []interface{}{"hello", "abc", "xyz"}))
}

func TestStack_Peek(t *testing.T) {
	stack := New()

	stack.Push("hello")
	value, err := stack.Peek()
	assert.Equal(t, value, "hello")
	assert.Nil(t, err)

	stack.Push("abc")
	value, err = stack.Peek()
	assert.Equal(t, value, "abc")
	assert.Nil(t, err)
}

func TestStack_Pop(t *testing.T) {
	stack := New()

	stack.Push(111, true)

	value, err := stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, stack.IsEmpty(), false)
	assert.Equal(t, stack.Size(), 1)
	assert.Equal(t, value, true)

	value, err = stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, stack.IsEmpty(), true)
	assert.Equal(t, stack.Size(), 0)
	assert.Equal(t, value, 111)

	_, err = stack.Pop()
	assert.NotNil(t, err)
}
