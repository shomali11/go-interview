package linkedliststacks

import (
	"errors"

	"github.com/shomali11/go-interview/datastructures/linkedlists/singlylinkedlists"
)

var (
	errEmptyStack = errors.New("stack is empty")
)

// New factory to generate new stacks
func New[T comparable](values ...T) *Stack[T] {
	return &Stack[T]{list: singlylinkedlists.New[T](values...)}
}

// Stack stack structure
type Stack[T comparable] struct {
	list *singlylinkedlists.SinglyLinkedList[T]
}

// Push add to the stack
func (s *Stack[T]) Push(values ...T) {
	s.list.Add(values...)
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the stack
func (s *Stack[T]) Size() int {
	return s.list.Size()
}

// Clear clears stack
func (s *Stack[T]) Clear() {
	s.list.Clear()
}

// Pop remove from the stack
func (s *Stack[T]) Pop() (res T, err error) {
	if s.IsEmpty() {
		return res, errEmptyStack
	}
	return s.list.RemoveAt(s.Size() - 1)
}

// Peek returns top of the stack
func (s *Stack[T]) Peek() (res T, err error) {
	if s.IsEmpty() {
		return res, errEmptyStack
	}
	return s.list.GetLastValue()
}

// GetValues returns values
func (s *Stack[T]) GetValues() []T {
	return s.list.GetValues()
}
