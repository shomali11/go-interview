package countdigits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDigits(t *testing.T) {
	assert.Equal(t, CountDigits(0), 1)
	assert.Equal(t, CountDigits(1), 1)
	assert.Equal(t, CountDigits(19), 2)
	assert.Equal(t, CountDigits(123), 3)
}
