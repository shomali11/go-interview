package armstrongs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPerfect(t *testing.T) {
	assert.True(t, IsArmstrong(0))
	assert.True(t, IsArmstrong(1))
	assert.True(t, IsArmstrong(153))
	assert.True(t, IsArmstrong(370))
	assert.True(t, IsArmstrong(371))
	assert.True(t, IsArmstrong(1634))
}
