package wordbreakers

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	dictionary = []string{"app", "apple", "pie", "applet", "let", "table", "tablet", "able", "t"}
)

func TestBreakWord(t *testing.T) {
	assert.Equal(t, BreakWord("applepie", dictionary), "apple pie")
	assert.Equal(t, BreakWord("applet", dictionary), "applet")
	assert.Equal(t, BreakWord("apples", dictionary), "")
	assert.Equal(t, BreakWord("boo", dictionary), "")
}

func TestExtractWords(t *testing.T) {
	values := ExtractWords("applet", dictionary)
	assert.True(t, contains(values, "applet"))
	assert.True(t, contains(values, "app let"))
	assert.True(t, contains(values, "apple t"))
	assert.True(t, reflect.DeepEqual(ExtractWords("apples", dictionary), []string{}))
	assert.True(t, reflect.DeepEqual(ExtractWords("boo", dictionary), []string{}))
}

func contains(values []string, value string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}
