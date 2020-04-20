package minmeetingrooms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMeetingRooms(t *testing.T) {
	meetings := [][]int{
		{0, 30}, {5, 10}, {15, 20},
	}
	assert.Equal(t, MinMeetingRooms(meetings), 2)
}
