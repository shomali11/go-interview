package hashmultimaps

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMultiMap_GetValues(t *testing.T) {
	multiMap := New[string, int]()
	multiMap.Put("hello", 1)

	assert.True(t, reflect.DeepEqual(multiMap.GetValues("hello"), []int{1}))
	assert.True(t, reflect.DeepEqual(multiMap.GetValues("unknown"), []int{}))
}

func TestHashMultiMap_Contains(t *testing.T) {
	multiMap := New[string, int]()
	multiMap.Put("hello", 1)

	assert.Equal(t, multiMap.Contains("hello"), true)
	assert.Equal(t, multiMap.Contains("unknown"), false)
	assert.Equal(t, multiMap.ContainsAll("hello", "unknown"), false)
	assert.Equal(t, multiMap.ContainsAny("hello", "unknown"), true)
}

func TestHashMultiMap_Clear(t *testing.T) {
	multiMap := New[string, int]()
	assert.Equal(t, multiMap.IsEmpty(), true)
	assert.Equal(t, multiMap.Size(), 0)

	multiMap.Put("hello", 1)
	assert.Equal(t, multiMap.IsEmpty(), false)
	assert.Equal(t, multiMap.Size(), 1)

	multiMap.Clear()
	assert.Equal(t, multiMap.IsEmpty(), true)
	assert.Equal(t, multiMap.Size(), 0)
}

func TestHashMultiMap_PutAll(t *testing.T) {
	multiMap := New[int, string]()

	multiMap.PutAll(111, "x", "x", "y", "z")
	assert.True(t, reflect.DeepEqual(multiMap.GetValues(111), []string{"x", "x", "y", "z"}))
	assert.Equal(t, multiMap.Size(), 1)
	assert.Equal(t, multiMap.GetKeys()[0], 111)
}

func TestHashMultiMap_Remove(t *testing.T) {
	multiMap := New[int, string]()

	multiMap.PutAll(111, "x", "x", "y", "z")
	multiMap.Remove(111, "x")

	assert.True(t, reflect.DeepEqual(multiMap.GetValues(111), []string{"x", "y", "z"}))
	assert.Equal(t, multiMap.Size(), 1)
	assert.Equal(t, multiMap.GetKeys()[0], 111)
}

func TestHashMultiMap_RemoveKey(t *testing.T) {
	multiMap := New[int, string]()

	multiMap.PutAll(111, "x", "x", "y", "z")
	multiMap.RemoveKey(111)

	assert.Equal(t, multiMap.Size(), 0)
	assert.True(t, reflect.DeepEqual(multiMap.GetValues(111), []string{}))
}

func TestHashMultiMapMerge(t *testing.T) {
	multiMap1 := New[bool, string]()
	multiMap1.Put(true, "true")

	multiMap2 := New[bool, string]()
	multiMap2.Put(true, "false")

	multiMap3 := New[bool, string]()
	multiMap3.Merge(multiMap1, multiMap2)

	values := multiMap3.GetValues(true)

	assert.Equal(t, multiMap3.Size(), 1)
	assert.True(t, reflect.DeepEqual(values, []string{"true", "false"}))
	assert.Equal(t, multiMap3.ContainsAll(true), true)
}
