package runningmedians

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunningMedian(t *testing.T) {
	runningMedian := New()

	runningMedian.Add(1)
	assert.Equal(t, runningMedian.GetMedian(), float64(1))

	runningMedian.Add(2)
	assert.Equal(t, runningMedian.GetMedian(), float64(1.5))

	runningMedian.Add(3)
	assert.Equal(t, runningMedian.GetMedian(), float64(2))
}
