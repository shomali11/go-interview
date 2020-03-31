package movingaverages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingAverage(t *testing.T) {
	movingAverage := New(3)

	movingAverage.Add(1)
	assert.Equal(t, movingAverage.GetAverage(), float64(1))
	assert.Equal(t, movingAverage.GetSum(), 1)

	movingAverage.Add(2)
	assert.Equal(t, movingAverage.GetAverage(), float64(1.5))
	assert.Equal(t, movingAverage.GetSum(), 3)

	movingAverage.Add(3)
	assert.Equal(t, movingAverage.GetAverage(), float64(2))
	assert.Equal(t, movingAverage.GetSum(), 6)

	movingAverage.Add(4)
	assert.Equal(t, movingAverage.GetAverage(), float64(3))
	assert.Equal(t, movingAverage.GetSum(), 9)

	movingAverage.Add(5)
	assert.Equal(t, movingAverage.GetAverage(), float64(4))
	assert.Equal(t, movingAverage.GetSum(), 12)
}
