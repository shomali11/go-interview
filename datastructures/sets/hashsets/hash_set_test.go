package hashsets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashSet(t *testing.T) {
	set := NewHashSet()
	set.Add("hello")

	assert.Equal(t, set.Contains("hello"), true)
	assert.Equal(t, set.Contains("unknown"), false)
	assert.Equal(t, set.Size(), 1)
	assert.Equal(t, set.List()[0], "hello")

	set.Clear()
	assert.Equal(t, set.Size(), 0)

	set.Add("cool")
	assert.Equal(t, set.Contains("cool"), true)
	assert.Equal(t, set.Size(), 1)
	assert.Equal(t, set.List()[0], "cool")

	set.Remove("cool")
	assert.Equal(t, set.Size(), 0)
	assert.Equal(t, set.Contains("cool"), false)
}

func TestHashSetMerge(t *testing.T) {
	set1 := NewHashSet()
	set1.Add("hello", "cool")

	set2 := NewHashSet("hello", "sweet")

	set3 := NewHashSet()
	set3.Merge(set1, set2)

	assert.Equal(t, set3.Size(), 3)
	assert.Equal(t, set3.Contains("hello", "cool", "sweet"), true)
}
