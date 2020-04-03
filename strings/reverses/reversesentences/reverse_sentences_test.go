package reversesentences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, ReverseSentence("hello"), "hello")
	assert.Equal(t, ReverseSentence("this is a sentence"), "sentence a is this")
	assert.Equal(t, ReverseSentence("      hello      "), "      hello      ")
	assert.Equal(t, ReverseSentence("   this is a sentence"), "sentence a is this   ")
	assert.Equal(t, ReverseSentence("What is happening?"), "happening? is What")
	assert.Equal(t, ReverseSentence("Hello! How are you?"), "you? are How Hello!")
}
