package duplicates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicates(t *testing.T) {
	assert.Equal(t, ContainsDuplicates(1, 2, 3), false)
	assert.Equal(t, ContainsDuplicates(1, 2, 1), true)
	assert.Equal(t, ContainsDuplicates(true, false), false)
	assert.Equal(t, ContainsDuplicates(true, false, true), true)
	assert.Equal(t, ContainsDuplicates("1", "2", "3"), false)
	assert.Equal(t, ContainsDuplicates("1", "2", "1"), true)
}
