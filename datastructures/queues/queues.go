package queues

import (
	"errors"
)

var (
	errEmptyQueue = errors.New("queue is empty")
)

// New factory to generate new Queues
func New(values ...interface{}) *Queue {
	Queue := Queue{make([]interface{}, 0, len(values))}
	Queue.Enqueue(values...)
	return &Queue
}

// Queue Queue structure
type Queue struct {
	array []interface{}
}

// Enqueue add to the Queue
func (q *Queue) Enqueue(values ...interface{}) {
	q.array = append(q.array, values...)
}

// IsEmpty checks if the Queue is empty
func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

// Size returns size of the Queue
func (q *Queue) Size() int {
	return len(q.array)
}

// Clear clears Queue
func (q *Queue) Clear() {
	q.array = nil
}

// Dequeue remove from the Queue
func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errEmptyQueue
	}

	value := q.array[0]
	q.array[0] = nil
	q.array = q.array[1:]
	return value, nil
}

// Peek returns front of the Queue
func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errEmptyQueue
	}

	value := q.array[0]
	return value, nil
}

// GetValues returns values
func (q *Queue) GetValues() []interface{} {
	values := make([]interface{}, 0, q.Size())
	for _, value := range q.array {
		values = append(values, value)
	}
	return values
}
