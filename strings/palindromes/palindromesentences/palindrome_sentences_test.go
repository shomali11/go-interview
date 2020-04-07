package palindromesentences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome("hannah"))
	assert.True(t, IsPalindrome("too hot to hoot."))
	assert.False(t, IsPalindrome("hello"))
}
