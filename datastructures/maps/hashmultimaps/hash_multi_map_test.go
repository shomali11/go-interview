package hashmultimaps

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMultiMap(t *testing.T) {
	multiMap := New()
	multiMap.Put("hello", 1)

	assert.Equal(t, multiMap.Contains("hello"), true)
	assert.True(t, reflect.DeepEqual(multiMap.GetValues("hello"), []interface{}{1}))
	assert.Equal(t, multiMap.Contains("unknown"), false)
	assert.True(t, reflect.DeepEqual(multiMap.GetValues("unknown"), []interface{}{}))
	assert.Equal(t, multiMap.ContainsAll("hello", "unknown"), false)
	assert.Equal(t, multiMap.ContainsAny("hello", "unknown"), true)
	assert.Equal(t, multiMap.Size(), 1)
	assert.Equal(t, multiMap.GetKeys()[0], "hello")

	multiMap.Clear()
	assert.Equal(t, multiMap.Size(), 0)

	multiMap.PutAll(111, "x", "x", "y", "z")
	assert.Equal(t, multiMap.Contains(111), true)
	assert.Equal(t, multiMap.ContainsAll(111), true)
	assert.Equal(t, multiMap.ContainsAny(111), true)
	assert.True(t, reflect.DeepEqual(multiMap.GetValues(111), []interface{}{"x", "x", "y", "z"}))
	assert.Equal(t, multiMap.Size(), 1)
	assert.Equal(t, multiMap.GetKeys()[0], 111)

	multiMap.Remove(111, "x")
	assert.Equal(t, multiMap.Contains(111), true)
	assert.Equal(t, multiMap.ContainsAll(111), true)
	assert.Equal(t, multiMap.ContainsAny(111), true)
	assert.True(t, reflect.DeepEqual(multiMap.GetValues(111), []interface{}{"x", "y", "z"}))
	assert.Equal(t, multiMap.Size(), 1)
	assert.Equal(t, multiMap.GetKeys()[0], 111)

	multiMap.RemoveKey(111)
	assert.Equal(t, multiMap.Size(), 0)
	assert.Equal(t, multiMap.Contains(111), false)
	assert.Equal(t, multiMap.ContainsAll(111), false)
	assert.Equal(t, multiMap.ContainsAny(111), false)
	assert.True(t, reflect.DeepEqual(multiMap.GetValues(111), []interface{}{}))
}

func TestHashMultiMapMerge(t *testing.T) {
	multiMap1 := New()
	multiMap1.Put(true, "true")

	multiMap2 := New()
	multiMap2.Put(true, 5)

	multiMap3 := New()
	multiMap3.Merge(multiMap1, multiMap2)

	values := multiMap3.GetValues(true)

	assert.Equal(t, multiMap3.Size(), 1)
	assert.True(t, reflect.DeepEqual(values, []interface{}{"true", 5}))
	assert.Equal(t, multiMap3.ContainsAll(true), true)
}
