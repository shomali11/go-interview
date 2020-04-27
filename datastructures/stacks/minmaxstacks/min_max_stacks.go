package minmaxstacks

import (
	"errors"

	"github.com/shomali11/go-interview/datastructures/stacks/slicestacks"
)

var (
	errEmptyStack = errors.New("stack is empty")
)

// New factory to generate new stacks
func New(compare func(i, j interface{}) bool, values ...interface{}) *Stack {
	stack := Stack{
		valuesStack: slicestacks.New(),
		minMaxStack: slicestacks.New(),
		compare:     compare,
	}
	stack.Push(values...)
	return &stack
}

// Stack stack structure
type Stack struct {
	valuesStack *slicestacks.Stack
	minMaxStack *slicestacks.Stack
	compare     func(i, j interface{}) bool
}

// Push add to the stack
func (s *Stack) Push(values ...interface{}) {
	if len(values) == 0 {
		return
	}

	var min interface{}
	var max interface{}

	if s.IsEmpty() {
		min = values[0]
		max = values[0]
	} else {
		minMaxInterface, _ := s.minMaxStack.Peek()
		minMax := minMaxInterface.([]interface{})
		min = minMax[0]
		max = minMax[1]
	}

	for i := 0; i < len(values); i++ {
		value := values[i]

		if s.compare(value, min) {
			min = value
		}

		if s.compare(max, value) {
			max = value
		}

		s.valuesStack.Push(value)
		s.minMaxStack.Push([]interface{}{min, max})
	}
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the stack
func (s *Stack) Size() int {
	return s.valuesStack.Size()
}

// Clear clears stack
func (s *Stack) Clear() {
	s.valuesStack.Clear()
	s.minMaxStack.Clear()
}

// Pop remove from the stack
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errEmptyStack
	}

	s.minMaxStack.Pop()
	return s.valuesStack.Pop()
}

// Peek returns top of the stack
func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errEmptyStack
	}

	return s.valuesStack.Peek()
}

// GetValues returns values
func (s *Stack) GetValues() []interface{} {
	return s.valuesStack.GetValues()
}

// GetMinMax returns the min and max values
func (s *Stack) GetMinMax() (interface{}, interface{}, error) {
	if s.IsEmpty() {
		return nil, nil, errEmptyStack
	}

	minMaxInterface, err := s.minMaxStack.Peek()
	minMax := minMaxInterface.([]interface{})
	return minMax[0], minMax[1], err
}
