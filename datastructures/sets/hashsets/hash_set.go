package hashsets

// New factory that creates a hash set
func New(values ...interface{}) *HashSet {
	set := HashSet{data: make(map[interface{}]struct{}, len(values))}
	set.Add(values...)
	return &set
}

// HashSet datastructure
type HashSet struct {
	data map[interface{}]struct{}
}

// Add adds values to the set
func (s *HashSet) Add(values ...interface{}) {
	for _, value := range values {
		s.data[value] = struct{}{}
	}
}

// Remove removes values from the set
func (s *HashSet) Remove(values ...interface{}) {
	for _, value := range values {
		delete(s.data, value)
	}
}

// Contains checks if the value is in the set
func (s *HashSet) Contains(value interface{}) bool {
	_, exists := s.data[value]
	return exists
}

// ContainsAll checks if all values are in the set
func (s *HashSet) ContainsAll(values ...interface{}) bool {
	for _, value := range values {
		if !s.Contains(value) {
			return false
		}
	}
	return true
}

// ContainsAny checks if any of the values are in the set
func (s *HashSet) ContainsAny(values ...interface{}) bool {
	for _, value := range values {
		if s.Contains(value) {
			return true
		}
	}
	return false
}

// Merge the two sets
func (s *HashSet) Merge(sets ...*HashSet) {
	for _, set := range sets {
		for _, value := range set.GetValues() {
			s.Add(value)
		}
	}
}

// Clear clears set
func (s *HashSet) Clear() {
	s.data = make(map[interface{}]struct{})
}

// GetValues returns values
func (s *HashSet) GetValues() []interface{} {
	values := make([]interface{}, 0, s.Size())
	for key := range s.data {
		values = append(values, key)
	}
	return values
}

// IsEmpty checks if the set is empty
func (s *HashSet) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the set
func (s *HashSet) Size() int {
	return len(s.data)
}

// Common set functions

// Copy makes an identical copy of the set
func (s *HashSet) Copy() *HashSet {
	return New(s.GetValues()...)
}

// Union makes a set that has all of the elements in either of two sets
func (s *HashSet) Union(ss *HashSet) *HashSet {
	new := s.Copy()
	new.Merge(ss)
	return new
}

// Intersection makes a set that has only the elements common to both of two sets
func (s *HashSet) Intersection(ss *HashSet) *HashSet {
	new := s.Copy()
	for _, v := range new.GetValues() {
		if !ss.Contains(v) {
			new.Remove(v)
		}
	}
	return new
}

// SymmetricDifference makes a set that has elements that are in one of two sets, but not both
func (s *HashSet) SymmetricDifference(ss *HashSet) *HashSet {
	new := &HashSet{make(map[interface{}]struct{}, s.Size())}
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
func (s *HashSet) Subtraction(ss *HashSet) *HashSet {
	new := s.Copy()
	for _, v := range ss.GetValues() {
		new.Remove(v)
	}
	return new
}
