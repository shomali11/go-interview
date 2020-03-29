package hashmultisets

import "sort"

// New factory that creates a new Hash Multi Set
func New(values ...string) *HashMultiSet {
	set := HashMultiSet{data: make(map[string]int)}
	set.Add(values...)
	return &set
}

// MultiSetPair a set's key/count pair
type MultiSetPair struct {
	Key   string
	Count int
}

// HashMultiSet a data structure representing a set with counts
type HashMultiSet struct {
	data map[string]int
}

// Merge merge multiple sets
func (s *HashMultiSet) Merge(sets ...*HashMultiSet) {
	for _, set := range sets {
		for _, value := range set.List() {
			s.IncrementBy(value, set.GetCount(value))
		}
	}
}

// Add adds a value to the set
func (s *HashMultiSet) Add(values ...string) {
	for _, value := range values {
		s.IncrementBy(value, 1)
	}
}

// IncrementBy increments a value's count by a number
func (s *HashMultiSet) IncrementBy(value string, count int) {
	existingCount := s.data[value]
	s.data[value] = existingCount + count
}

// List returns a list of the set's values
func (s *HashMultiSet) List() []string {
	values := make([]string, 0)
	for key := range s.data {
		values = append(values, key)
	}
	return values
}

// Contains checks if a value is in the set
func (s *HashMultiSet) Contains(values ...string) bool {
	for _, value := range values {
		_, exists := s.data[value]
		if !exists {
			return false
		}
	}
	return true
}

// GetCount returns count associated with the value
func (s *HashMultiSet) GetCount(value string) int {
	return s.data[value]
}

// Remove removes a value
func (s *HashMultiSet) Remove(values ...string) {
	for _, value := range values {
		delete(s.data, value)
	}
}

// Clear clears the set
func (s *HashMultiSet) Clear() {
	s.data = make(map[string]int)
}

// IsEmpty checks if the set is empty
func (s *HashMultiSet) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns size of the set
func (s *HashMultiSet) Size() int {
	return len(s.data)
}

// Top clears the set
func (s *HashMultiSet) Top() []MultiSetPair {
	setPairs := make([]MultiSetPair, 0)
	for key, count := range s.data {
		setPairs = append(setPairs, MultiSetPair{Key: key, Count: count})
	}

	sort.SliceStable(setPairs, func(i, j int) bool {
		return setPairs[i].Count > setPairs[j].Count
	})
	return setPairs
}
