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
func New[T comparable](compare func(i, j T) bool, values ...T) *PriorityQueue[T] {
	priorityQueue := PriorityQueue[T]{
		pq:       &heapArray[T]{compare: compare},
		multiMap: hashmultimaps.New[T, *pqElement[T]](),
	}
	heap.Init(priorityQueue.pq)
	priorityQueue.Push(values...)
	return &priorityQueue
}

// PriorityQueue Priority Queue structure
type PriorityQueue[T comparable] struct {
	pq       *heapArray[T]
	multiMap *hashmultimaps.HashMultiMap[T, *pqElement[T]]
}

// Push pushes to the Priority Queue
func (s *PriorityQueue[T]) Push(values ...T) {
	for _, value := range values {
		element := &pqElement[T]{value: value}
		heap.Push(s.pq, element)
		s.multiMap.Put(value, element)
	}
}

// Contains checks if the value exists in the Priority Queue
func (s *PriorityQueue[T]) Contains(value T) bool {
	return s.multiMap.Contains(value)
}

// Remove removes from the Priority Queue
func (s *PriorityQueue[T]) Remove(values ...T) {
	for _, value := range values {
		elements := s.multiMap.GetValues(value)
		if len(elements) == 0 {
			continue
		}

		element := elements[0]
		heap.Remove(s.pq, element.index)
		s.multiMap.Remove(value, element)
	}
}

// IsEmpty checks if the Priority Queue is empty
func (s *PriorityQueue[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the Priority Queue
func (s *PriorityQueue[T]) Size() int {
	return len(s.pq.array)
}

// Clear clears the Priority Queue
func (s *PriorityQueue[T]) Clear() {
	s.pq.array = nil
}

// Pop removes from the Priority Queue
func (s *PriorityQueue[T]) Pop() (res T, err error) {
	if s.IsEmpty() {
		return res, errEmptyQueue
	}

	element := heap.Pop(s.pq).(*pqElement[T])
	return element.value, nil
}

// Peek returns top of the Priority Queue
func (s *PriorityQueue[T]) Peek() (res T, err error) {
	if s.IsEmpty() {
		return res, errEmptyQueue
	}

	element := s.pq.array[0]
	return element.value, nil
}
