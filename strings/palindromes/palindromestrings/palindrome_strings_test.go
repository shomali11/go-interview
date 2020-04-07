package palindromestrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome("hannah"))
	assert.True(t, IsPalindrome("hanah"))
	assert.False(t, IsPalindrome("hello"))
}
