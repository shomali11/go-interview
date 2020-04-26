package slicestacks

import (
	"errors"
)

var (
	errEmptyStack = errors.New("stack is empty")
)

// New factory to generate new stacks
func New(values ...interface{}) *Stack {
	stack := Stack{make([]interface{}, 0, len(values))}
	stack.Push(values...)
	return &stack
}

// Stack stack structure
type Stack struct {
	array []interface{}
}

// Push add to the stack
func (s *Stack) Push(values ...interface{}) {
	s.array = append(s.array, values...)
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the stack
func (s *Stack) Size() int {
	return len(s.array)
}

// Clear clears stack
func (s *Stack) Clear() {
	s.array = nil
}

// Pop remove from the stack
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errEmptyStack
	}

	size := s.Size()
	value := s.array[size-1]
	s.array[size-1] = nil
	s.array = s.array[:size-1]
	return value, nil
}

// Peek returns top of the stack
func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errEmptyStack
	}

	value := s.array[s.Size()-1]
	return value, nil
}

// GetValues returns values
func (s *Stack) GetValues() []interface{} {
	values := make([]interface{}, 0, s.Size())
	for _, value := range s.array {
		values = append(values, value)
	}
	return values
}
