package countprimes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPrimes(t *testing.T) {
	assert.Equal(t, CountPrimes(1), 0)
	assert.Equal(t, CountPrimes(2), 1)
	assert.Equal(t, CountPrimes(3), 2)
	assert.Equal(t, CountPrimes(4), 2)
	assert.Equal(t, CountPrimes(10), 4)
}
