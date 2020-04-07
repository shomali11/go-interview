package priorityqueues

import (
	"container/heap"
	"errors"

	"github.com/shomali11/go-interview/datastructures/maps/hashmultimaps"
)

var (
	errEmptyQueue = errors.New("queue is empty")
)

// New factory to generate new priority queues
func New(compare func(i, j interface{}) bool, values ...interface{}) *PriorityQueue {
	priorityQueue := PriorityQueue{pq: &heapArray{compare: compare}, multiMap: hashmultimaps.New()}
	heap.Init(priorityQueue.pq)
	priorityQueue.Push(values...)
	return &priorityQueue
}

// PriorityQueue Priority Queue structure
type PriorityQueue struct {
	pq       *heapArray
	multiMap *hashmultimaps.HashMultiMap
}

// Push pushes to the Priority Queue
func (s *PriorityQueue) Push(values ...interface{}) {
	for _, value := range values {
		element := &pqElement{value: value}
		heap.Push(s.pq, element)
		s.multiMap.Put(value, element)
	}
}

// Contains checks if the value exists in the Priority Queue
func (s *PriorityQueue) Contains(value interface{}) bool {
	return s.multiMap.Contains(value)
}

// Remove removes from the Priority Queue
func (s *PriorityQueue) Remove(values ...interface{}) {
	for _, value := range values {
		elements := s.multiMap.GetValues(value)
		if len(elements) == 0 {
			continue
		}

		element := elements[0].(*pqElement)
		heap.Remove(s.pq, element.index)
		s.multiMap.Remove(value, element)
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
func (s *PriorityQueue) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errEmptyQueue
	}

	element := heap.Pop(s.pq).(*pqElement)
	return element.value, nil
}

// Peek returns top of the Priority Queue
func (s *PriorityQueue) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errEmptyQueue
	}

	element := s.pq.array[0]
	return element.value, nil
}
