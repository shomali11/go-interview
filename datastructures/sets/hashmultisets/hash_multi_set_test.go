package hashmultisets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMultiSet(t *testing.T) {
	set := New[string]()
	set.Add("hello")

	assert.Equal(t, set.Contains("hello"), true)
	assert.Equal(t, set.GetCount("hello"), 1)
	assert.Equal(t, set.Contains("unknown"), false)
	assert.Equal(t, set.GetCount("unknown"), 0)
	assert.Equal(t, set.ContainsAll("hello", "unknown"), false)
	assert.Equal(t, set.ContainsAny("hello", "unknown"), true)
	assert.Equal(t, set.Size(), 1)
	assert.Equal(t, set.GetValues()[0], "hello")

	set.Clear()
	assert.Equal(t, set.Size(), 0)

	set.Add("111")
	assert.Equal(t, set.Contains("111"), true)
	assert.Equal(t, set.ContainsAll("111"), true)
	assert.Equal(t, set.ContainsAny("111"), true)
	assert.Equal(t, set.GetCount("111"), 1)
	assert.Equal(t, set.Size(), 1)
	assert.Equal(t, set.GetValues()[0], "111")

	set.Remove("111")
	assert.Equal(t, set.Size(), 0)
	assert.Equal(t, set.Contains("111"), false)
	assert.Equal(t, set.ContainsAll("111"), false)
	assert.Equal(t, set.ContainsAny("111"), false)
	assert.Equal(t, set.GetCount("111"), 0)
}

func TestHashMultiSetMerge(t *testing.T) {
	set1 := New[int]()
	set1.Add(1, 2)

	set2 := New[int](1, 3)

	set3 := New[int]()
	set3.Merge(set1, set2)

	assert.Equal(t, set3.Size(), 3)
	assert.Equal(t, set3.GetCount(1), 2)
	assert.Equal(t, set3.GetCount(2), 1)
	assert.Equal(t, set3.GetCount(3), 1)
	assert.Equal(t, set3.ContainsAll(1, 2, 3), true)
}

func TestHashMultiSetTop(t *testing.T) {
	set := New[string]()
	set.IncrementBy("key1", 10)
	set.IncrementBy("key2", 15)

	pairs := set.GetTopValues()
	assert.Equal(t, len(pairs), 2)

	pair1 := pairs[0]
	assert.Equal(t, pair1.Key, "key2")
	assert.Equal(t, pair1.Count, 15)

	pair2 := pairs[1]
	assert.Equal(t, pair2.Key, "key1")
	assert.Equal(t, pair2.Count, 10)
}
