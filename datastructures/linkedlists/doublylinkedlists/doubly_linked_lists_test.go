package doublylinkedlists

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedList(t *testing.T) {
	list := New()

	_, err := list.GetFirstValue()
	assert.NotNil(t, err)

	_, err = list.GetLastValue()
	assert.NotNil(t, err)

	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{}))
	assert.True(t, reflect.DeepEqual(list.GetReverseValues(), []interface{}{}))
	assert.Equal(t, list.IsEmpty(), true)
	assert.Equal(t, list.Size(), 0)

	err = list.InsertAt(3, "hello")
	assert.NotNil(t, err)

	list.Add("hello")

	value, err := list.GetFirstValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "hello")

	value, err = list.GetLastValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "hello")

	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{"hello"}))
	assert.True(t, reflect.DeepEqual(list.GetReverseValues(), []interface{}{"hello"}))
	assert.Equal(t, list.GetIndexOf("hello"), 0)
	assert.Equal(t, list.IsEmpty(), false)
	assert.Equal(t, list.Size(), 1)

	list.InsertAt(0, "cool")

	value, err = list.GetFirstValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "cool")

	value, err = list.GetLastValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "hello")

	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{"cool", "hello"}))
	assert.True(t, reflect.DeepEqual(list.GetReverseValues(), []interface{}{"hello", "cool"}))
	assert.Equal(t, list.GetIndexOf("cool"), 0)
	assert.Equal(t, list.GetIndexOf("hello"), 1)
	assert.Equal(t, list.IsEmpty(), false)
	assert.Equal(t, list.Size(), 2)

	list.InsertAt(1, "awesome")

	value, err = list.GetFirstValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "cool")

	value, err = list.GetLastValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "hello")

	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{"cool", "awesome", "hello"}))
	assert.True(t, reflect.DeepEqual(list.GetReverseValues(), []interface{}{"hello", "awesome", "cool"}))
	assert.Equal(t, list.GetIndexOf("cool"), 0)
	assert.Equal(t, list.GetIndexOf("awesome"), 1)
	assert.Equal(t, list.GetIndexOf("hello"), 2)
	assert.Equal(t, list.IsEmpty(), false)
	assert.Equal(t, list.Size(), 3)

	list.InsertAt(2, "fantastic")

	value, err = list.GetFirstValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "cool")

	value, err = list.GetLastValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "hello")

	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{"cool", "awesome", "fantastic", "hello"}))
	assert.True(t, reflect.DeepEqual(list.GetReverseValues(), []interface{}{"hello", "fantastic", "awesome", "cool"}))
	assert.Equal(t, list.GetIndexOf("cool"), 0)
	assert.Equal(t, list.GetIndexOf("awesome"), 1)
	assert.Equal(t, list.GetIndexOf("fantastic"), 2)
	assert.Equal(t, list.GetIndexOf("hello"), 3)
	assert.Equal(t, list.IsEmpty(), false)
	assert.Equal(t, list.Size(), 4)

	value, err = list.RemoveAt(3)
	assert.Equal(t, value, "hello")
	assert.Nil(t, err)
	assert.Equal(t, list.GetIndexOf("hello"), -1)
	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{"cool", "awesome", "fantastic"}))
	assert.True(t, reflect.DeepEqual(list.GetReverseValues(), []interface{}{"fantastic", "awesome", "cool"}))
	assert.Equal(t, list.GetIndexOf("cool"), 0)
	assert.Equal(t, list.GetIndexOf("awesome"), 1)
	assert.Equal(t, list.GetIndexOf("fantastic"), 2)
	assert.Equal(t, list.IsEmpty(), false)
	assert.Equal(t, list.Size(), 3)

	_, err = list.RemoveAt(3)
	assert.NotNil(t, err)

	list.Clear()
	assert.Equal(t, list.GetIndexOf("hello"), -1)
	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{}))
	assert.True(t, reflect.DeepEqual(list.GetReverseValues(), []interface{}{}))
	assert.Equal(t, list.IsEmpty(), true)
	assert.Equal(t, list.Size(), 0)

	list.Add("1", "2", "3", "4")
	assert.Equal(t, list.GetIndexOf("1"), 0)
	assert.Equal(t, list.GetIndexOf("2"), 1)
	assert.Equal(t, list.GetIndexOf("3"), 2)
	assert.Equal(t, list.GetIndexOf("4"), 3)

	value, err = list.GetFirstValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "1")

	value, err = list.GetLastValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "4")

	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{"1", "2", "3", "4"}))
	assert.True(t, reflect.DeepEqual(list.GetReverseValues(), []interface{}{"4", "3", "2", "1"}))
	assert.Equal(t, list.IsEmpty(), false)
	assert.Equal(t, list.Size(), 4)

	value, err = list.RemoveAt(2)
	assert.Equal(t, value, "3")
	assert.Nil(t, err)

	assert.Equal(t, list.GetIndexOf("1"), 0)
	assert.Equal(t, list.GetIndexOf("2"), 1)
	assert.Equal(t, list.GetIndexOf("4"), 2)
	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{"1", "2", "4"}))
	assert.True(t, reflect.DeepEqual(list.GetReverseValues(), []interface{}{"4", "2", "1"}))
	assert.Equal(t, list.IsEmpty(), false)
	assert.Equal(t, list.Size(), 3)

	list.Clear()
	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{}))
	assert.True(t, reflect.DeepEqual(list.GetReverseValues(), []interface{}{}))
	assert.Equal(t, list.IsEmpty(), true)
	assert.Equal(t, list.Size(), 0)

	list.Add("a", "b", "c", "d")
	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{"a", "b", "c", "d"}))
	assert.True(t, reflect.DeepEqual(list.GetReverseValues(), []interface{}{"d", "c", "b", "a"}))

	value, err = list.RemoveAt(3)
	assert.Equal(t, value, "d")
	assert.Nil(t, err)

	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{"a", "b", "c"}))
	assert.True(t, reflect.DeepEqual(list.GetReverseValues(), []interface{}{"c", "b", "a"}))
	assert.Equal(t, list.IsEmpty(), false)
	assert.Equal(t, list.Size(), 3)

	for !list.IsEmpty() {
		list.RemoveAt(0)
	}
	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{}))
	assert.True(t, reflect.DeepEqual(list.GetReverseValues(), []interface{}{}))
	assert.Equal(t, list.IsEmpty(), true)
	assert.Equal(t, list.Size(), 0)

	list.Clear()
	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{}))
	assert.True(t, reflect.DeepEqual(list.GetReverseValues(), []interface{}{}))
	assert.Equal(t, list.IsEmpty(), true)
	assert.Equal(t, list.Size(), 0)

	_, err = list.RemoveAt(0)
	assert.NotNil(t, err)
}
