package minmaxstacks

import (
	"errors"

	"github.com/shomali11/go-interview/datastructures/stacks/slicestacks"
)

var (
	errEmptyStack = errors.New("stack is empty")
)

// New factory to generate new stacks
func New[T any](compare func(i, j T) bool, values ...T) *Stack[T] {
	stack := Stack[T]{
		valuesStack: slicestacks.New[T](),
		minMaxStack: slicestacks.New[[]T](),
		compare:     compare,
	}
	stack.Push(values...)
	return &stack
}

// Stack stack structure
type Stack[T any] struct {
	valuesStack *slicestacks.Stack[T]
	minMaxStack *slicestacks.Stack[[]T]
	compare     func(i, j T) bool
}

// Push add to the stack
func (s *Stack[T]) Push(values ...T) {
	if len(values) == 0 {
		return
	}

	var min T
	var max T

	if s.IsEmpty() {
		min = values[0]
		max = values[0]
	} else {
		minMaxInterface, _ := s.minMaxStack.Peek()
		min = minMaxInterface[0]
		max = minMaxInterface[1]
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
		s.minMaxStack.Push([]T{min, max})
	}
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the stack
func (s *Stack[T]) Size() int {
	return s.valuesStack.Size()
}

// Clear clears stack
func (s *Stack[T]) Clear() {
	s.valuesStack.Clear()
	s.minMaxStack.Clear()
}

// Pop remove from the stack
func (s *Stack[T]) Pop() (res T, err error) {
	if s.IsEmpty() {
		return res, errEmptyStack
	}

	s.minMaxStack.Pop()
	return s.valuesStack.Pop()
}

// Peek returns top of the stack
func (s *Stack[T]) Peek() (res T, err error) {
	if s.IsEmpty() {
		return res, errEmptyStack
	}

	return s.valuesStack.Peek()
}

// GetValues returns values
func (s *Stack[T]) GetValues() []T {
	return s.valuesStack.GetValues()
}

// GetMinMax returns the min and max values
func (s *Stack[T]) GetMinMax() (min T, max T, err error) {
	if s.IsEmpty() {
		return min, max, errEmptyStack
	}

	minMaxInterface, err := s.minMaxStack.Peek()
	minMax := minMaxInterface
	return minMax[0], minMax[1], err
}
