package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	assert.Equal(t, Sort(""), "")
	assert.Equal(t, Sort("a"), "a")
	assert.Equal(t, Sort("abc"), "abc")
	assert.Equal(t, Sort("cba"), "abc")
}
