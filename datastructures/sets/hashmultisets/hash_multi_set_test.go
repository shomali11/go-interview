package hashmultisets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMultiSet(t *testing.T) {
	set := New()
	set.Add("hello")

	assert.Equal(t, set.Contains("hello"), true)
	assert.Equal(t, set.GetCount("hello"), 1)
	assert.Equal(t, set.Contains("unknown"), false)
	assert.Equal(t, set.GetCount("unknown"), 0)
	assert.Equal(t, set.Size(), 1)
	assert.Equal(t, set.List()[0], "hello")

	set.Clear()
	assert.Equal(t, set.Size(), 0)

	set.Add("cool")
	assert.Equal(t, set.Contains("cool"), true)
	assert.Equal(t, set.GetCount("cool"), 1)
	assert.Equal(t, set.Size(), 1)
	assert.Equal(t, set.List()[0], "cool")

	set.Remove("cool")
	assert.Equal(t, set.Size(), 0)
	assert.Equal(t, set.Contains("cool"), false)
	assert.Equal(t, set.GetCount("cool"), 0)
}

func TestHashMultiSetMerge(t *testing.T) {
	set1 := New()
	set1.Add("hello", "cool")

	set2 := New("hello", "sweet")

	set3 := New()
	set3.Merge(set1, set2)

	assert.Equal(t, set3.Size(), 3)
	assert.Equal(t, set3.GetCount("hello"), 2)
	assert.Equal(t, set3.GetCount("cool"), 1)
	assert.Equal(t, set3.GetCount("sweet"), 1)
	assert.Equal(t, set3.Contains("hello", "cool", "sweet"), true)
}

func TestHashMultiSetTop(t *testing.T) {
	set := New()
	set.IncrementBy("key1", 10)
	set.IncrementBy("key2", 15)

	pairs := set.Top()
	assert.Equal(t, len(pairs), 2)

	pair1 := pairs[0]
	assert.Equal(t, pair1.Key, "key2")
	assert.Equal(t, pair1.Count, 15)

	pair2 := pairs[1]
	assert.Equal(t, pair2.Key, "key1")
	assert.Equal(t, pair2.Count, 10)
}
