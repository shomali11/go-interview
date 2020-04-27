package minmaxstacks

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Clear(t *testing.T) {
	compare := func(i, j interface{}) bool {
		return i.(string) < j.(string)
	}

	stack := New(compare)
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
	compare := func(i, j interface{}) bool {
		return i.(string) < j.(string)
	}

	stack := New(compare)

	stack.Push("hello", "abc", "xyz")
	assert.True(t, reflect.DeepEqual(stack.GetValues(), []interface{}{"hello", "abc", "xyz"}))
}

func TestStack_Peek(t *testing.T) {
	compare := func(i, j interface{}) bool {
		return i.(string) < j.(string)
	}

	stack := New(compare)

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
	compare := func(i, j interface{}) bool {
		return i.(int) < j.(int)
	}

	stack := New(compare)

	stack.Push(111, 222)

	value, err := stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, stack.IsEmpty(), false)
	assert.Equal(t, stack.Size(), 1)
	assert.Equal(t, value, 222)

	value, err = stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, stack.IsEmpty(), true)
	assert.Equal(t, stack.Size(), 0)
	assert.Equal(t, value, 111)

	_, err = stack.Pop()
	assert.NotNil(t, err)
}

func TestStack_GetMinMax(t *testing.T) {
	compare := func(i, j interface{}) bool {
		return i.(int) < j.(int)
	}

	stack := New(compare)

	stack.Push(111, 222, 55, 333)

	min, max, err := stack.GetMinMax()
	assert.Nil(t, err)
	assert.Equal(t, min, 55)
	assert.Equal(t, max, 333)

	stack.Pop()

	min, max, err = stack.GetMinMax()
	assert.Nil(t, err)
	assert.Equal(t, min, 55)
	assert.Equal(t, max, 222)

	stack.Pop()

	min, max, err = stack.GetMinMax()
	assert.Nil(t, err)
	assert.Equal(t, min, 111)
	assert.Equal(t, max, 222)

	stack.Pop()

	min, max, err = stack.GetMinMax()
	assert.Nil(t, err)
	assert.Equal(t, min, 111)
	assert.Equal(t, max, 111)

	stack.Pop()

	_, _, err = stack.GetMinMax()
	assert.NotNil(t, err)
}
