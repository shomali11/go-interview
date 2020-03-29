package words

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	dictionary = []string{"app", "apple", "applet", "let", "table", "tablet", "able", "t"}
)

func TestBreakWord(t *testing.T) {
	assert.Equal(t, BreakWord("applet", dictionary), "applet")
	assert.Equal(t, BreakWord("apples", dictionary), "")
	assert.Equal(t, BreakWord("boo", dictionary), "")
}

func TestExtractWords(t *testing.T) {
	assert.True(t, reflect.DeepEqual(ExtractWords("applet", dictionary), []string{"applet", "app let", "apple t"}))
	assert.True(t, reflect.DeepEqual(ExtractWords("apples", dictionary), []string{}))
	assert.True(t, reflect.DeepEqual(ExtractWords("boo", dictionary), []string{}))
}
