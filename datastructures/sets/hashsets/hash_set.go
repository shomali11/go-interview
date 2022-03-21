package hashsets

// New factory that creates a hash set
func New[T comparable](values ...T) *HashSet[T] {
	set := HashSet[T]{data: make(map[T]struct{}, len(values))}
	set.Add(values...)
	return &set
}

// HashSet datastructure
type HashSet[T comparable] struct {
	data map[T]struct{}
}

// Add adds values to the set
func (s *HashSet[T]) Add(values ...T) {
	for _, value := range values {
		s.data[value] = struct{}{}
	}
}

// Remove removes values from the set
func (s *HashSet[T]) Remove(values ...T) {
	for _, value := range values {
		delete(s.data, value)
	}
}

// Contains checks if the value is in the set
func (s *HashSet[T]) Contains(value T) bool {
	_, exists := s.data[value]
	return exists
}

// ContainsAll checks if all values are in the set
func (s *HashSet[T]) ContainsAll(values ...T) bool {
	for _, value := range values {
		if !s.Contains(value) {
			return false
		}
	}
	return true
}

// ContainsAny checks if any of the values are in the set
func (s *HashSet[T]) ContainsAny(values ...T) bool {
	for _, value := range values {
		if s.Contains(value) {
			return true
		}
	}
	return false
}

// Merge the two sets
func (s *HashSet[T]) Merge(sets ...*HashSet[T]) {
	for _, set := range sets {
		for _, value := range set.GetValues() {
			s.Add(value)
		}
	}
}

// Clear clears set
func (s *HashSet[T]) Clear() {
	s.data = make(map[T]struct{})
}

// GetValues returns values
func (s *HashSet[T]) GetValues() []T {
	values := make([]T, 0, s.Size())
	for key := range s.data {
		values = append(values, key)
	}
	return values
}

// IsEmpty checks if the set is empty
func (s *HashSet[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the set
func (s *HashSet[T]) Size() int {
	return len(s.data)
}

// Common set functions

// Copy makes an identical copy of the set
func (s *HashSet[T]) Copy() *HashSet[T] {
	return New[T](s.GetValues()...)
}

// Union makes a set that has all of the elements in either of two sets
func (s *HashSet[T]) Union(ss *HashSet[T]) *HashSet[T] {
	new := s.Copy()
	new.Merge(ss)
	return new
}

// Intersection makes a set that has only the elements common to both of two sets
func (s *HashSet[T]) Intersection(ss *HashSet[T]) *HashSet[T] {
	new := s.Copy()
	for _, v := range new.GetValues() {
		if !ss.Contains(v) {
			new.Remove(v)
		}
	}
	return new
}

// SymmetricDifference makes a set that has elements that are in one of two sets, but not both
func (s *HashSet[T]) SymmetricDifference(ss *HashSet[T]) *HashSet[T] {
	new := &HashSet[T]{make(map[T]struct{}, s.Size())}
	for _, v := range s.GetValues() {
		if !ss.Contains(v) {
			new.Add(v)
		}
	}
	for _, v := range ss.GetValues() {
		if !s.Contains(v) {
			new.Add(v)
		}
	}
	return new
}

// Subtraction makes a set with the elements that are in the first set, but not the second
func (s *HashSet[T]) Subtraction(ss *HashSet[T]) *HashSet[T] {
	new := s.Copy()
	for _, v := range ss.GetValues() {
		new.Remove(v)
	}
	return new
}
