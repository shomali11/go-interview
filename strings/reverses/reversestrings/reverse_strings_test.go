package reversestrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, ReverseString("hello"), "olleh")
}
