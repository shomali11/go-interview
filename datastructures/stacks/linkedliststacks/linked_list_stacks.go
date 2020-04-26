package linkedliststacks

import (
	"errors"

	"github.com/shomali11/go-interview/datastructures/linkedlists/singlylinkedlists"
)

var (
	errEmptyStack = errors.New("stack is empty")
)

// New factory to generate new stacks
func New(values ...interface{}) *Stack {
	return &Stack{list: singlylinkedlists.New(values...)}
}

// Stack stack structure
type Stack struct {
	list *singlylinkedlists.SinglyLinkedList
}

// Push add to the stack
func (s *Stack) Push(values ...interface{}) {
	s.list.Add(values...)
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the stack
func (s *Stack) Size() int {
	return s.list.Size()
}

// Clear clears stack
func (s *Stack) Clear() {
	s.list.Clear()
}

// Pop remove from the stack
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errEmptyStack
	}
	return s.list.RemoveAt(s.Size() - 1)
}

// Peek returns top of the stack
func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errEmptyStack
	}
	return s.list.GetLastValue()
}

// GetValues returns values
func (s *Stack) GetValues() []interface{} {
	return s.list.GetValues()
}
