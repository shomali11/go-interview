package priorityqueues

import (
	"container/heap"
	"errors"
)

var (
	errEmptyQueue = errors.New("queue is empty")
)

// PQElement is an element in a priority queue.
type PQElement struct {
	Value    string
	Priority int

	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// NewMax factory to generate new max priority queues
func NewMax(elements ...*PQElement) *PriorityQueue {
	return New(true, elements...)
}

// NewMin factory to generate new min priority queues
func NewMin(elements ...*PQElement) *PriorityQueue {
	return New(false, elements...)
}

// New factory to generate new priority queues
func New(max bool, elements ...*PQElement) *PriorityQueue {
	priorityQueue := PriorityQueue{pq: &heapArray{max: max}}
	heap.Init(priorityQueue.pq)
	priorityQueue.Push(elements...)
	return &priorityQueue
}

// PriorityQueue Priority Queue structure
type PriorityQueue struct {
	pq *heapArray
}

// Push pushes to the Priority Queue
func (s *PriorityQueue) Push(elements ...*PQElement) {
	for _, element := range elements {
		heap.Push(s.pq, element)
	}
}

// IsEmpty checks if the Priority Queue is empty
func (s *PriorityQueue) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the Priority Queue
func (s *PriorityQueue) Size() int {
	return len(s.pq.array)
}

// Clear clears the Priority Queue
func (s *PriorityQueue) Clear() {
	s.pq.array = nil
}

// Pop removes from the Priority Queue
func (s *PriorityQueue) Pop() (*PQElement, error) {
	if s.IsEmpty() {
		return nil, errEmptyQueue
	}

	element := heap.Pop(s.pq).(*PQElement)
	return element, nil
}

// Peek returns top of the Priority Queue
func (s *PriorityQueue) Peek() (*PQElement, error) {
	if s.IsEmpty() {
		return nil, errEmptyQueue
	}

	element := s.pq.array[0]
	return element, nil
}
