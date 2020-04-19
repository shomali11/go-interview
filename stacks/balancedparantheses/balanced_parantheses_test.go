package balancedparantheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalancedParantheses(t *testing.T) {
	assert.True(t, IsBalancedParantheses(""))
	assert.True(t, IsBalancedParantheses(" "))
	assert.True(t, IsBalancedParantheses("{}"))
	assert.True(t, IsBalancedParantheses("[]"))
	assert.True(t, IsBalancedParantheses("()"))
	assert.True(t, IsBalancedParantheses("{()}"))
	assert.True(t, IsBalancedParantheses("[{()}]"))
	assert.True(t, IsBalancedParantheses("[](){}"))
	assert.True(t, IsBalancedParantheses("[()]{}"))
	assert.True(t, IsBalancedParantheses("[ (X)]{   }"))

	assert.False(t, IsBalancedParantheses("{"))
	assert.False(t, IsBalancedParantheses("("))
	assert.False(t, IsBalancedParantheses("["))
	assert.False(t, IsBalancedParantheses("}"))
	assert.False(t, IsBalancedParantheses(")"))
	assert.False(t, IsBalancedParantheses("]"))
	assert.False(t, IsBalancedParantheses("}{"))
	assert.False(t, IsBalancedParantheses(")("))
	assert.False(t, IsBalancedParantheses("]["))
	assert.False(t, IsBalancedParantheses("{[}]"))
	assert.False(t, IsBalancedParantheses("(((()))"))
}
