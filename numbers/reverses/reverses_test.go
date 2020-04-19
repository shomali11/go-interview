package reverses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, Reverse(12321), 12321)
	assert.Equal(t, Reverse(8998), 8998)
	assert.Equal(t, Reverse(12345), 54321)
}
