package addbinaries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	value, err := Add("11", "")
	assert.Nil(t, err)
	assert.Equal(t, value, "11")

	value, err = Add("11", "1")
	assert.Nil(t, err)
	assert.Equal(t, value, "100")

	value, err = Add("10", "11")
	assert.Nil(t, err)
	assert.Equal(t, value, "101")

	_, err = Add("a", "11")
	assert.NotNil(t, err)
}
