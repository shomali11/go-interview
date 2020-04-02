package hashsets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashSet(t *testing.T) {
	set := New()
	assert.Equal(t, set.IsEmpty(), true)

	set.Add("hello")

	assert.Equal(t, set.Contains("hello"), true)
	assert.Equal(t, set.Contains("unknown"), false)
	assert.Equal(t, set.ContainsAll("hello" , "unknown"), false)
	assert.Equal(t, set.ContainsAny("hello" , "unknown"), true)
	assert.Equal(t, set.IsEmpty(), false)
	assert.Equal(t, set.Size(), 1)
	assert.Equal(t, set.GetValues()[0], "hello")

	set.Clear()
	assert.Equal(t, set.IsEmpty(), true)
	assert.Equal(t, set.Size(), 0)

	set.Add("cool")
	assert.Equal(t, set.Contains("cool"), true)
	assert.Equal(t, set.ContainsAll("cool"), true)
	assert.Equal(t, set.ContainsAny("cool"), true)
	assert.Equal(t, set.Size(), 1)
	assert.Equal(t, set.GetValues()[0], "cool")

	set.Remove("cool")
	assert.Equal(t, set.Size(), 0)
	assert.Equal(t, set.Contains("cool"), false)
	assert.Equal(t, set.ContainsAll("cool"), false)
	assert.Equal(t, set.ContainsAny("cool"), false)
}

func TestHashSetMerge(t *testing.T) {
	set1 := New()
	set1.Add("hello", "cool")

	set2 := New("hello", "sweet")

	set3 := New()
	set3.Merge(set1, set2)

	assert.Equal(t, set3.Size(), 3)
	assert.Equal(t, set3.ContainsAll("hello", "cool", "sweet"), true)
}
