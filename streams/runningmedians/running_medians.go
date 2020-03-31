package runningmedians

import "github.com/shomali11/go-interview/datastructures/priorityqueues"

// New generates a Running Median
func New() *RunningMedian {
	return &RunningMedian{minPQ: priorityqueues.NewMin(), maxPQ: priorityqueues.NewMax()}
}

// RunningMedian keeps track of the running median
type RunningMedian struct {
	minPQ *priorityqueues.PriorityQueue
	maxPQ *priorityqueues.PriorityQueue
}

// Add adds a number to the running median calculations
func (rm *RunningMedian) Add(number int) {
	push(rm.minPQ, number)
	push(rm.maxPQ, pop(rm.minPQ))

	if rm.minPQ.Size() < rm.maxPQ.Size() {
		push(rm.minPQ, pop(rm.maxPQ))
	}
}

// GetMedian calculates median
func (rm *RunningMedian) GetMedian() float64 {
	if rm.minPQ.Size() > rm.maxPQ.Size() {
		return peek(rm.minPQ)
	}
	return (peek(rm.minPQ) + peek(rm.maxPQ)) / 2.0
}

func push(pq *priorityqueues.PriorityQueue, value int) {
	pq.Push(&priorityqueues.PQElement{Priority: value})
}

func pop(pq *priorityqueues.PriorityQueue) int {
	element, _ := pq.Pop()
	return element.Priority
}

func peek(pq *priorityqueues.PriorityQueue) float64 {
	element, _ := pq.Peek()
	return float64(element.Priority)
}
