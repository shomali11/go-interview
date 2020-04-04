package excels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	assert.Equal(t, Encode(1), "A")
	assert.Equal(t, Encode(2), "B")
	assert.Equal(t, Encode(26), "Z")
	assert.Equal(t, Encode(27), "AA")
	assert.Equal(t, Encode(100), "CV")
}

func TestDecode(t *testing.T) {
	assert.Equal(t, Decode("A"), 1)
	assert.Equal(t, Decode("B"), 2)
	assert.Equal(t, Decode("Z"), 26)
	assert.Equal(t, Decode("AA"), 27)
	assert.Equal(t, Decode("CV"), 100)
}
