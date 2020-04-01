package countrotations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountRotations(t *testing.T) {
	assert.Equal(t, Count([]int{1, 2, 3, 4, 5}), 0)
	assert.Equal(t, Count([]int{2, 3, 4, 5, 1}), 1)
	assert.Equal(t, Count([]int{3, 4, 5, 1, 2}), 2)
	assert.Equal(t, Count([]int{4, 5, 1, 2, 3}), 3)
	assert.Equal(t, Count([]int{5, 1, 2, 3, 4}), 4)
}
