package palindromestrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome(12321))
	assert.True(t, IsPalindrome(8998))
	assert.False(t, IsPalindrome(12345))
}
