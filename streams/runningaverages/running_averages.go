package runningaverages

// RunningAverage keeps track of the running average
type RunningAverage struct {
	sum   int
	count int
}

// Add adds a number to the running average calculations
func (ra *RunningAverage) Add(number int) {
	ra.sum += number
	ra.count++
}

// GetAverage calculates the running average
func (ra *RunningAverage) GetAverage() float64 {
	return float64(ra.sum) / float64(ra.count)
}

// GetSum returns the running sum
func (ra *RunningAverage) GetSum() int {
	return ra.sum
}

// GetCount returns the running count
func (ra *RunningAverage) GetCount() int {
	return ra.count
}
