package doublylinkedlists

import (
	"errors"
)

var (
	errEmptyList        = errors.New("list is empty")
	errIndexOutOfBounds = errors.New("index is out of bounds")
)

// New factory to generate new doubly linked lists
func New(values ...interface{}) *DoublyLinkedList {
	list := DoublyLinkedList{}
	list.Add(values...)
	return &list
}

// DLLNode doubly linked list node
type DLLNode struct {
	Value    interface{}
	Previous *DLLNode
	Next     *DLLNode
}

// DoublyLinkedList doubly linked list structure
type DoublyLinkedList struct {
	count int
	head  *DLLNode
	tail  *DLLNode
}

// IsEmpty checks if the list is empty
func (s *DoublyLinkedList) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the list
func (s *DoublyLinkedList) Size() int {
	return s.count
}

// Clear clears list
func (s *DoublyLinkedList) Clear() {
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
func (s *DoublyLinkedList) GetValues() []interface{} {
	values := make([]interface{}, 0, s.Size())
	current := s.head
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	current = nil
	return values
}

// GetReverseValues returns values from tail to head
func (s *DoublyLinkedList) GetReverseValues() []interface{} {
	values := make([]interface{}, 0, s.Size())
	current := s.tail
	for current != nil {
		values = append(values, current.Value)
		current = current.Previous
	}
	current = nil
	return values
}

// GetIndexOf returns the index of the first occurence
func (s *DoublyLinkedList) GetIndexOf(value interface{}) int {
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

// GetLastIndexOf returns the index of the last occurence
func (s *DoublyLinkedList) GetLastIndexOf(value interface{}) int {
	index := s.Size() - 1
	current := s.tail
	for current != nil {
		if value == current.Value {
			return index
		}
		current = current.Previous
		index--
	}
	current = nil
	return -1
}

// Add add to the list
func (s *DoublyLinkedList) Add(values ...interface{}) {
	for _, value := range values {
		s.InsertAt(s.Size(), value)
	}
}

// InsertAt insert value at specific index in the list
func (s *DoublyLinkedList) InsertAt(index int, value interface{}) error {
	if index < 0 || index > s.count {
		return errIndexOutOfBounds
	}

	element := &DLLNode{Value: value}
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
func (s *DoublyLinkedList) RemoveAt(index int) (interface{}, error) {
	if index < 0 || index >= s.count {
		return nil, errIndexOutOfBounds
	}

	if s.count == 1 {
		value := s.head.Value
		s.head = nil
		s.tail = nil
		s.count--
		return value, nil
	}

	if index == 0 {
		current := s.head
		s.head = current.Next
		current.Next.Previous = nil
		current.Next = nil
		value := current.Value
		s.count--
		current = nil
		return value, nil
	}

	if index == s.count-1 {
		current := s.tail
		s.tail = current.Previous
		current.Previous.Next = nil
		current.Previous = nil
		value := current.Value
		s.count--
		current = nil
		return value, nil
	}

	current, err := s.getNode(index)
	if err != nil {
		return nil, err
	}

	value := current.Value
	current.Previous.Next = current.Next
	current.Next.Previous = current.Previous
	s.count--
	current = nil
	return value, nil
}

// GetValueAt returns value at a specific index in the list
func (s *DoublyLinkedList) GetValueAt(index int) (interface{}, error) {
	current, err := s.getNode(index)
	if err != nil {
		return nil, err
	}
	return current.Value, nil
}

// GetFirstValue returns first value in the list
func (s *DoublyLinkedList) GetFirstValue() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errEmptyList
	}
	return s.head.Value, nil
}

// GetLastValue returns last value in the list
func (s *DoublyLinkedList) GetLastValue() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errEmptyList
	}
	return s.tail.Value, nil
}

// GetHead returns head node of the list
func (s *DoublyLinkedList) GetHead() *DLLNode {
	return s.head
}

// GetTail returns tail node of the list
func (s *DoublyLinkedList) GetTail() *DLLNode {
	return s.tail
}

func (s *DoublyLinkedList) getNode(index int) (*DLLNode, error) {
	if index < 0 || index >= s.count {
		return nil, errIndexOutOfBounds
	}

	var current *DLLNode
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
