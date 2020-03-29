package stacks

import (
	"errors"
)

const (
	empty = ""
)

// New factory to generate new stacks
func New(values ...string) *Stack {
	stack := Stack{array: make([]string, 0)}
	stack.Push(values...)
	return &stack
}

// Stack stack structure
type Stack struct {
	array []string
}

// Push add to the stack
func (s *Stack) Push(values ...string) {
	for _, value := range values {
		s.array = append(s.array, value)
	}
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.array) == 0
}

// Size returns size of the stack
func (s *Stack) Size() int {
	return len(s.array)
}

// Clear clears stack
func (s *Stack) Clear() {
	s.array = make([]string, 0)
}

// Pop remove from the stack
func (s *Stack) Pop() (string, error) {
	if len(s.array) == 0 {
		return empty, errors.New("stack is empty")
	}

	value := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return value, nil
}

// Peek returns top of the stack
func (s *Stack) Peek() (string, error) {
	if len(s.array) == 0 {
		return empty, errors.New("stack is empty")
	}

	value := s.array[len(s.array)-1]
	return value, nil
}
