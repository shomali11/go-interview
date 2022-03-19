package singlylinkedlists

import (
	"errors"
)

var (
	errEmptyList        = errors.New("list is empty")
	errIndexOutOfBounds = errors.New("index is out of bounds")
)

// New factory to generate new singly linked lists
func New[T comparable](values ...T) *SinglyLinkedList[T] {
	list := SinglyLinkedList[T]{}
	list.Add(values...)
	return &list
}

// SLLNode singly linked list node
type SLLNode[T comparable] struct {
	Value T
	Next  *SLLNode[T]
}

// SinglyLinkedList singly linked list structure
type SinglyLinkedList[T comparable] struct {
	count int
	head  *SLLNode[T]
	tail  *SLLNode[T]
}

// IsEmpty checks if the list is empty
func (s *SinglyLinkedList[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the list
func (s *SinglyLinkedList[T]) Size() int {
	return s.count
}

// Clear clears list
func (s *SinglyLinkedList[T]) Clear() {
	current := s.head
	for current != nil {
		temp := current
		current = current.Next
		temp.Next = nil
		temp = nil
	}
	current = nil
	s.count = 0
	s.head = nil
	s.tail = nil
}

// GetValues returns values
func (s *SinglyLinkedList[T]) GetValues() []T {
	values := make([]T, 0, s.Size())
	current := s.head
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	current = nil
	return values
}

// GetIndexOf returns the index of the first occurence
func (s *SinglyLinkedList[T]) GetIndexOf(value T) int {
	index := 0
	current := s.head
	for current != nil {
		if value == current.Value {
			return index
		}
		current = current.Next
		index++
	}
	current = nil
	return -1
}

// Add add to the list
func (s *SinglyLinkedList[T]) Add(values ...T) {
	for _, value := range values {
		s.InsertAt(s.Size(), value)
	}
}

// InsertAt insert value at specific index in the list
func (s *SinglyLinkedList[T]) InsertAt(index int, value T) error {
	if index < 0 || index > s.count {
		return errIndexOutOfBounds
	}

	element := &SLLNode[T]{Value: value}
	if s.IsEmpty() {
		s.head = element
		s.tail = element
		s.count++
		element = nil
		return nil
	}

	if index == 0 {
		element.Next = s.head
		s.head = element
		s.count++
		element = nil
		return nil
	}

	if index == s.count {
		s.tail.Next = element
		s.tail = element
		s.count++
		element = nil
		return nil
	}

	current, err := s.getNode(index - 1)
	if err != nil {
		return err
	}

	element.Next = current.Next
	current.Next = element
	s.count++
	current = nil
	element = nil
	return nil
}

// RemoveAt remove value from the list at specific index
func (s *SinglyLinkedList[T]) RemoveAt(index int) (res T, err error) {
	if index < 0 || index >= s.count {
		return res, errIndexOutOfBounds
	}

	if s.count == 1 {
		res = s.head.Value
		s.head = nil
		s.tail = nil
		s.count--
		return res, nil
	}

	if index == 0 {
		current := s.head
		s.head = current.Next
		current.Next = nil
		res = current.Value
		s.count--
		current = nil
		return res, nil
	}

	current, err := s.getNode(index - 1)
	if err != nil {
		return res, err
	}

	res = current.Next.Value
	current.Next = current.Next.Next
	if index == s.count-1 {
		s.tail = current
	}
	s.count--
	current = nil
	return res, nil
}

// GetValueAt returns value at a specific index in the list
func (s *SinglyLinkedList[T]) GetValueAt(index int) (res T, err error) {
	current, err := s.getNode(index)
	if err != nil {
		return res, err
	}
	return current.Value, nil
}

// GetFirstValue returns first value in the list
func (s *SinglyLinkedList[T]) GetFirstValue() (res T, err error) {
	if s.IsEmpty() {
		return res, errEmptyList
	}
	return s.head.Value, nil
}

// GetLastValue returns last value in the list
func (s *SinglyLinkedList[T]) GetLastValue() (res T, err error) {
	if s.IsEmpty() {
		return res, errEmptyList
	}
	return s.tail.Value, nil
}

// GetHead returns head node of the list
func (s *SinglyLinkedList[T]) GetHead() *SLLNode[T] {
	return s.head
}

func (s *SinglyLinkedList[T]) getNode(index int) (*SLLNode[T], error) {
	if index < 0 || index >= s.count {
		return nil, errIndexOutOfBounds
	}

	current := s.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current, nil
}
