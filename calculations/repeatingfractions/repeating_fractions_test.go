package repeatingfractions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	assert.Equal(t, Divide(0, 0), "ERROR")
	assert.Equal(t, Divide(2, 1), "2")
	assert.Equal(t, Divide(1, 2), "0.5")
	assert.Equal(t, Divide(0, 3), "0")
	assert.Equal(t, Divide(2, 3), "0.(6)")
	assert.Equal(t, Divide(10, 3), "3.(3)")
	assert.Equal(t, Divide(22, 7), "3.(142857)")
	assert.Equal(t, Divide(100, 145), "0.(6896551724137931034482758620)")
}
