package slicestacks

import (
	"errors"
)

var (
	errEmptyStack = errors.New("stack is empty")
)

// New factory to generate new stacks
func New[T any](values ...T) *Stack[T] {
	stack := Stack[T]{make([]T, 0, len(values))}
	stack.Push(values...)
	return &stack
}

// Stack stack structure
type Stack[T any] struct {
	array []T
}

// Push add to the stack
func (s *Stack[T]) Push(values ...T) {
	s.array = append(s.array, values...)
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the stack
func (s *Stack[T]) Size() int {
	return len(s.array)
}

// Clear clears stack
func (s *Stack[T]) Clear() {
	s.array = nil
}

// Pop remove from the stack
func (s *Stack[T]) Pop() (res T, err error) {
	if s.IsEmpty() {
		return res, errEmptyStack
	}

	size := s.Size()
	value := s.array[size-1]
	s.array = s.array[:size-1]
	return value, nil
}

// Peek returns top of the stack
func (s *Stack[T]) Peek() (res T, err error) {
	if s.IsEmpty() {
		return res, errEmptyStack
	}

	value := s.array[s.Size()-1]
	return value, nil
}

// GetValues returns values
func (s *Stack[T]) GetValues() []T {
	values := make([]T, 0, s.Size())
	for _, value := range s.array {
		values = append(values, value)
	}
	return values
}
