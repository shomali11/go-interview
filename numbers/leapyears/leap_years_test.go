package leapyears

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLeapYear(t *testing.T) {
	for year := 1804; year <= 1896; year = year + 4 {
		assert.True(t, IsLeapYear(year))
	}

	for year := 1904; year <= 2000; year = year + 4 {
		assert.True(t, IsLeapYear(year))
	}

	for year := 2004; year <= 2096; year = year + 4 {
		assert.True(t, IsLeapYear(year))
	}

	for year := 2104; year <= 2196; year = year + 4 {
		assert.True(t, IsLeapYear(year))
	}

	for year := 2204; year <= 2296; year = year + 4 {
		assert.True(t, IsLeapYear(year))
	}

	for year := 2304; year <= 2400; year = year + 4 {
		assert.True(t, IsLeapYear(year))
	}

	assert.False(t, IsLeapYear(1900))
	assert.False(t, IsLeapYear(2100))
	assert.False(t, IsLeapYear(2200))
	assert.False(t, IsLeapYear(2300))
}
