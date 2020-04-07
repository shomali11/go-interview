package runningmedians

import "github.com/shomali11/go-interview/datastructures/priorityqueues"

// New generates a Running Median
func New() *RunningMedian {
	maxCompare := func(i, j interface{}) bool {
		return i.(int) > j.(int)
	}

	minCompare := func(i, j interface{}) bool {
		return i.(int) < j.(int)
	}
	return &RunningMedian{minPQ: priorityqueues.New(minCompare), maxPQ: priorityqueues.New(maxCompare)}
}

// RunningMedian keeps track of the running median
type RunningMedian struct {
	minPQ *priorityqueues.PriorityQueue
	maxPQ *priorityqueues.PriorityQueue
}

// Add adds a number to the running median calculations
func (rm *RunningMedian) Add(number int) {
	rm.minPQ.Push(number)
	rm.maxPQ.Push(pop(rm.minPQ))

	if rm.minPQ.Size() < rm.maxPQ.Size() {
		rm.minPQ.Push(pop(rm.maxPQ))
	}
}

// GetMedian calculates median
func (rm *RunningMedian) GetMedian() float64 {
	if rm.minPQ.Size() > rm.maxPQ.Size() {
		return peek(rm.minPQ)
	}
	return (peek(rm.minPQ) + peek(rm.maxPQ)) / 2.0
}

func pop(pq *priorityqueues.PriorityQueue) int {
	value, _ := pq.Pop()
	return value.(int)
}

func peek(pq *priorityqueues.PriorityQueue) float64 {
	value, _ := pq.Peek()
	return float64(value.(int))
}
