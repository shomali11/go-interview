package minrotations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.Equal(t, Min([]int{1, 2, 3, 4, 5}), 0)
	assert.Equal(t, Min([]int{2, 3, 4, 5, 1}), 4)
	assert.Equal(t, Min([]int{3, 4, 5, 1, 2}), 3)
	assert.Equal(t, Min([]int{4, 5, 1, 2, 3}), 2)
	assert.Equal(t, Min([]int{5, 1, 2, 3, 4}), 1)
}
