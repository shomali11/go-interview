package hashmultisets

import "sort"

// New factory that creates a new Hash Multi Set
func New[T comparable](values ...T) *HashMultiSet[T] {
	set := HashMultiSet[T]{data: make(map[T]int, len(values))}
	set.Add(values...)
	return &set
}

// MultiSetPair a set's key/count pair
type MultiSetPair[T comparable] struct {
	Key   T
	Count int
}

// HashMultiSet a data structure representing a set with counts
type HashMultiSet[T comparable] struct {
	data map[T]int
}

// Merge merge multiple sets
func (s *HashMultiSet[T]) Merge(sets ...*HashMultiSet[T]) {
	for _, set := range sets {
		for _, value := range set.GetValues() {
			s.IncrementBy(value, set.GetCount(value))
		}
	}
}

// Add adds a value to the set
func (s *HashMultiSet[T]) Add(values ...T) {
	for _, value := range values {
		s.IncrementBy(value, 1)
	}
}

// IncrementBy increments a value's count by a number
func (s *HashMultiSet[T]) IncrementBy(value T, count int) {
	existingCount := s.data[value]
	s.data[value] = existingCount + count
}

// GetValues returns a list of the set's values
func (s *HashMultiSet[T]) GetValues() []T {
	values := make([]T, 0, s.Size())
	for key := range s.data {
		values = append(values, key)
	}
	return values
}

// Contains checks if a value is in the set
func (s *HashMultiSet[T]) Contains(value T) bool {
	_, exists := s.data[value]
	return exists
}

// ContainsAll checks if all values are in the set
func (s *HashMultiSet[T]) ContainsAll(values ...T) bool {
	for _, value := range values {
		if !s.Contains(value) {
			return false
		}
	}
	return true
}

// ContainsAny checks if any values are in the set
func (s *HashMultiSet[T]) ContainsAny(values ...T) bool {
	for _, value := range values {
		if s.Contains(value) {
			return true
		}
	}
	return false
}

// GetCount returns count associated with the value
func (s *HashMultiSet[T]) GetCount(value T) int {
	return s.data[value]
}

// Remove removes a value
func (s *HashMultiSet[T]) Remove(values ...T) {
	for _, value := range values {
		delete(s.data, value)
	}
}

// Clear clears the set
func (s *HashMultiSet[T]) Clear() {
	s.data = make(map[T]int)
}

// IsEmpty checks if the set is empty
func (s *HashMultiSet[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the set
func (s *HashMultiSet[T]) Size() int {
	return len(s.data)
}

// GetTopValues returns values ordered in descending order
func (s *HashMultiSet[T]) GetTopValues() []MultiSetPair[T] {
	setPairs := make([]MultiSetPair[T], 0, s.Size())
	for key, count := range s.data {
		setPairs = append(setPairs, MultiSetPair[T]{Key: key, Count: count})
	}

	sort.SliceStable(setPairs, func(i, j int) bool {
		return setPairs[i].Count > setPairs[j].Count
	})
	return setPairs
}
