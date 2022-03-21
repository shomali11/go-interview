package doublylinkedlists

import (
	"errors"
	"reflect"
)

var (
	errEmptyList        = errors.New("list is empty")
	errIndexOutOfBounds = errors.New("index is out of bounds")
)

// New factory to generate new doubly linked lists
func New[T any](values ...T) *DoublyLinkedList[T] {
	list := DoublyLinkedList[T]{}
	list.Add(values...)
	return &list
}

// DLLNode doubly linked list node
type DLLNode[T any] struct {
	Value    T
	Previous *DLLNode[T]
	Next     *DLLNode[T]
}

// DoublyLinkedList doubly linked list structure
type DoublyLinkedList[T any] struct {
	count int
	head  *DLLNode[T]
	tail  *DLLNode[T]
}

// IsEmpty checks if the list is empty
func (s *DoublyLinkedList[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the list
func (s *DoublyLinkedList[T]) Size() int {
	return s.count
}

// Clear clears list
func (s *DoublyLinkedList[T]) Clear() {
	current := s.head
	for current != nil {
		temp := current
		current = current.Next
		temp.Previous = nil
		temp.Next = nil
		temp = nil
	}
	current = nil
	s.count = 0
	s.head = nil
	s.tail = nil
}

// GetValues returns values from head to tail
func (s *DoublyLinkedList[T]) GetValues() []T {
	values := make([]T, 0, s.Size())
	current := s.head
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	current = nil
	return values
}

// GetReverseValues returns values from tail to head
func (s *DoublyLinkedList[T]) GetReverseValues() []T {
	values := make([]T, 0, s.Size())
	current := s.tail
	for current != nil {
		values = append(values, current.Value)
		current = current.Previous
	}
	current = nil
	return values
}

// GetIndexOf returns the index of the first occurence
func (s *DoublyLinkedList[T]) GetIndexOf(value T) int {
	index := 0
	current := s.head
	for current != nil {
		if reflect.DeepEqual(value, current.Value) {
			return index
		}
		current = current.Next
		index++
	}
	current = nil
	return -1
}

// GetLastIndexOf returns the index of the last occurence
func (s *DoublyLinkedList[T]) GetLastIndexOf(value T) int {
	index := s.Size() - 1
	current := s.tail
	for current != nil {
		if reflect.DeepEqual(value, current.Value) {
			return index
		}
		current = current.Previous
		index--
	}
	current = nil
	return -1
}

// Add add to the list
func (s *DoublyLinkedList[T]) Add(values ...T) {
	for _, value := range values {
		s.InsertAt(s.Size(), value)
	}
}

// InsertAt insert value at specific index in the list
func (s *DoublyLinkedList[T]) InsertAt(index int, value T) error {
	if index < 0 || index > s.count {
		return errIndexOutOfBounds
	}

	element := &DLLNode[T]{Value: value}
	if s.IsEmpty() {
		s.head = element
		s.tail = element
		s.count++
		element = nil
		return nil
	}

	if index == 0 {
		element.Next = s.head
		s.head.Previous = element
		s.head = element
		s.count++
		element = nil
		return nil
	}

	if index == s.count {
		s.tail.Next = element
		element.Previous = s.tail
		s.tail = element
		s.count++
		element = nil
		return nil
	}

	current, err := s.getNode(index)
	if err != nil {
		return err
	}

	current.Previous.Next = element
	element.Previous = current.Previous
	element.Next = current
	current.Previous = element
	s.count++
	current = nil
	element = nil
	return nil
}

// RemoveAt remove value from the list at specific index
func (s *DoublyLinkedList[T]) RemoveAt(index int) (res T, err error) {
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
		current.Next.Previous = nil
		current.Next = nil
		res = current.Value
		s.count--
		current = nil
		return res, nil
	}

	if index == s.count-1 {
		current := s.tail
		s.tail = current.Previous
		current.Previous.Next = nil
		current.Previous = nil
		res = current.Value
		s.count--
		current = nil
		return res, nil
	}

	current, err := s.getNode(index)
	if err != nil {
		return res, err
	}

	res = current.Value
	current.Previous.Next = current.Next
	current.Next.Previous = current.Previous
	s.count--
	current = nil
	return res, nil
}

// GetValueAt returns value at a specific index in the list
func (s *DoublyLinkedList[T]) GetValueAt(index int) (res T, err error) {
	current, err := s.getNode(index)
	if err != nil {
		return res, err
	}
	return current.Value, nil
}

// GetFirstValue returns first value in the list
func (s *DoublyLinkedList[T]) GetFirstValue() (res T, err error) {
	if s.IsEmpty() {
		return res, errEmptyList
	}
	return s.head.Value, nil
}

// GetLastValue returns last value in the list
func (s *DoublyLinkedList[T]) GetLastValue() (res T, err error) {
	if s.IsEmpty() {
		return res, errEmptyList
	}
	return s.tail.Value, nil
}

// GetHead returns head node of the list
func (s *DoublyLinkedList[T]) GetHead() *DLLNode[T] {
	return s.head
}

// GetTail returns tail node of the list
func (s *DoublyLinkedList[T]) GetTail() *DLLNode[T] {
	return s.tail
}

func (s *DoublyLinkedList[T]) getNode(index int) (*DLLNode[T], error) {
	if index < 0 || index >= s.count {
		return nil, errIndexOutOfBounds
	}

	var current *DLLNode[T]
	if index-0 <= s.count-1-index {
		current = s.head
		for i := 0; i < index; i++ {
			current = current.Next
		}
	} else {
		current = s.tail
		for i := 0; i < s.count-1-index; i++ {
			current = current.Previous
		}
	}
	return current, nil
}
