package hashsets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashSet(t *testing.T) {
	set := New()
	assert.Equal(t, set.IsEmpty(), true)

	set.Add("hello")

	assert.Equal(t, set.Contains("hello"), true)
	assert.Equal(t, set.Contains("unknown"), false)
	assert.Equal(t, set.ContainsAll("hello", "unknown"), false)
	assert.Equal(t, set.ContainsAny("hello", "unknown"), true)
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

func TestHashSetCopy(t *testing.T) {
	set1 := New()
	set1.Add("a", "b", "c")
	set2 := set1.Copy()
	assert.Equal(t, set2.Size(), 3)
	assert.Equal(t, set2.ContainsAll("a", "b", "c"), true)
}

func TestHashSetUnion(t *testing.T) {
	set1 := New()
	set1.Add("a", "b")
	set2 := New()
	set2.Add("c", "d")
	set3 := set1.Union(set2)
	assert.Equal(t, set3.Size(), 4)
	assert.Equal(t, set3.ContainsAll("a", "b", "c", "d"), true)
}

func TestHashSetIntersection(t *testing.T) {
	set1 := New()
	set1.Add("a", "b", "c", "d", "e")
	set2 := New()
	set2.Add("c", "d", "e", "f", "g")
	set3 := set1.Intersection(set2)
	assert.Equal(t, set3.Size(), 3)
	assert.Equal(t, set3.ContainsAll("c", "d", "e"), true)
}

func TestHashSetSymmetricDifference(t *testing.T) {
	set1 := New()
	set1.Add("a", "b", "c", "d", "e")
	set2 := New()
	set2.Add("c", "d", "e", "f", "g")
	set3 := set1.SymmetricDifference(set2)
	assert.Equal(t, set3.Size(), 4)
	assert.Equal(t, set3.ContainsAll("a", "b", "f", "g"), true)
}

func TestHashSetSubtraction(t *testing.T) {
	set1 := New()
	set1.Add("a", "b", "c", "d", "e")
	set2 := New()
	set2.Add("c", "d", "e", "f", "g")
	set3 := set1.Subtraction(set2)
	assert.Equal(t, set3.Size(), 2)
	assert.Equal(t, set3.ContainsAll("a", "b"), true)
}
