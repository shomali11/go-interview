package queues

import (
	"errors"
)

var (
	errEmptyQueue = errors.New("queue is empty")
)

// New factory to generate new Queues
func New[T any](values ...T) *Queue[T] {
	Queue := Queue[T]{make([]T, 0, len(values))}
	Queue.Enqueue(values...)
	return &Queue
}

// Queue Queue structure
type Queue[T any] struct {
	array []T
}

// Enqueue add to the Queue
func (q *Queue[T]) Enqueue(values ...T) {
	q.array = append(q.array, values...)
}

// IsEmpty checks if the Queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

// Size returns size of the Queue
func (q *Queue[T]) Size() int {
	return len(q.array)
}

// Clear clears Queue
func (q *Queue[T]) Clear() {
	q.array = nil
}

// Dequeue remove from the Queue
func (q *Queue[T]) Dequeue() (res T, err error) {
	if q.IsEmpty() {
		return res, errEmptyQueue
	}

	res = q.array[0]
	q.array = q.array[1:]
	return res, nil
}

// Peek returns front of the Queue
func (q *Queue[T]) Peek() (res T, err error) {
	if q.IsEmpty() {
		return res, errEmptyQueue
	}

	res = q.array[0]
	return res, nil
}

// GetValues returns values
func (q *Queue[T]) GetValues() []T {
	values := make([]T, 0, q.Size())
	values = append(values, q.array...)
	return values
}
