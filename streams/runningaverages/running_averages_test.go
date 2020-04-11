package runningaverages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunningAverage(t *testing.T) {
	runningAverage := RunningAverage{}

	runningAverage.Add(1)
	assert.Equal(t, runningAverage.GetAverage(), float64(1))
	assert.Equal(t, runningAverage.GetSum(), 1)
	assert.Equal(t, runningAverage.GetCount(), 1)

	runningAverage.Add(2)
	assert.Equal(t, runningAverage.GetAverage(), float64(1.5))
	assert.Equal(t, runningAverage.GetSum(), 3)
	assert.Equal(t, runningAverage.GetCount(), 2)

	runningAverage.Add(3)
	assert.Equal(t, runningAverage.GetAverage(), float64(2))
	assert.Equal(t, runningAverage.GetSum(), 6)
	assert.Equal(t, runningAverage.GetCount(), 3)
}
