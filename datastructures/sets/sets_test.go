package sets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	set := NewSet()
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
