package queues

import (
	"errors"
)

var (
	errEmptyQueue = errors.New("queue is empty")
)

// New factory to generate new Queues
func New(values ...interface{}) *Queue {
	Queue := Queue{}
	Queue.Enqueue(values...)
	return &Queue
}

// Queue Queue structure
type Queue struct {
	array []interface{}
}

// Enqueue add to the Queue
func (s *Queue) Enqueue(values ...interface{}) {
	s.array = append(s.array, values...)
}

// IsEmpty checks if the Queue is empty
func (s *Queue) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the Queue
func (s *Queue) Size() int {
	return len(s.array)
}

// Clear clears Queue
func (s *Queue) Clear() {
	s.array = nil
}

// Dequeue remove from the Queue
func (s *Queue) Dequeue() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errEmptyQueue
	}

	value := s.array[0]
	s.array[0] = nil
	s.array = s.array[1:]
	return value, nil
}

// Peek returns front of the Queue
func (s *Queue) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errEmptyQueue
	}

	value := s.array[0]
	return value, nil
}
