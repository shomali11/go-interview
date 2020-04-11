package singlylinkedlists

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinked_GetFirstValue(t *testing.T) {
	list := New()

	_, err := list.GetFirstValue()
	assert.NotNil(t, err)

	list.Add("abc")

	value, err := list.GetFirstValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "abc")

	list.Add("def")

	value, err = list.GetFirstValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "abc")

	list.InsertAt(0, "xyz")

	value, err = list.GetFirstValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "xyz")
}

func TestSinglyLinked_GetLastValue(t *testing.T) {
	list := New()

	_, err := list.GetLastValue()
	assert.NotNil(t, err)

	list.Add("abc")

	value, err := list.GetLastValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "abc")

	list.Add("def")

	value, err = list.GetLastValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "def")

	list.InsertAt(2, "xyz")

	value, err = list.GetLastValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "xyz")
}

func TestSinglyLinked_GetValues(t *testing.T) {
	list := New()
	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{}))

	list.Add("a", "b", "c")
	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{"a", "b", "c"}))

	list.InsertAt(1, "d")
	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{"a", "d", "b", "c"}))
}

func TestSinglyLinkedList_IsEmpty(t *testing.T) {
	list := New()
	assert.Equal(t, list.IsEmpty(), true)

	list.Add("hello")
	assert.Equal(t, list.IsEmpty(), false)

	list.RemoveAt(0)
	assert.Equal(t, list.IsEmpty(), true)
}

func TestSinglyLinkedList_Size(t *testing.T) {
	list := New()
	assert.Equal(t, list.Size(), 0)

	list.Add("hello")
	assert.Equal(t, list.Size(), 1)

	list.Add("abc")
	assert.Equal(t, list.Size(), 2)

	list.RemoveAt(0)
	assert.Equal(t, list.Size(), 1)

	list.RemoveAt(0)
	assert.Equal(t, list.Size(), 0)
}

func TestSinglyLinkedList_GetIndexOf(t *testing.T) {
	list := New()
	assert.Equal(t, list.GetIndexOf("hello"), -1)

	list.Add("hello")
	assert.Equal(t, list.GetIndexOf("hello"), 0)

	list.Add("abc")
	assert.Equal(t, list.GetIndexOf("abc"), 1)

	list.Add("abc")
	assert.Equal(t, list.GetIndexOf("abc"), 1)
}

func TestSinglyLinkedList_InsertAt(t *testing.T) {
	list := New()

	err := list.InsertAt(3, "hello")
	assert.NotNil(t, err)

	err = list.InsertAt(0, "hello")
	assert.Nil(t, err)
}

func TestSinglyLinkedList_RemoveAt(t *testing.T) {
	list := New()

	_, err := list.RemoveAt(0)
	assert.NotNil(t, err)

	list.Add("hello")
	value, err := list.RemoveAt(0)
	assert.Equal(t, value, "hello")
	assert.Nil(t, err)
}

func TestSinglyLinkedList_Clear(t *testing.T) {
	list := New()

	list.Clear()
	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{}))
	assert.Equal(t, list.IsEmpty(), true)
	assert.Equal(t, list.Size(), 0)

	list.Add("hello")

	list.Clear()
	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{}))
	assert.Equal(t, list.IsEmpty(), true)
	assert.Equal(t, list.Size(), 0)
}
