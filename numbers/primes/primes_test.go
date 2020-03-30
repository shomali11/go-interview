package primes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	assert.True(t, IsPrime(7))
	assert.True(t, IsPrime(31))
	assert.False(t, IsPrime(4))
}
