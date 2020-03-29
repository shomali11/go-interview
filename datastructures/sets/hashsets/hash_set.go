package hashsets

// NewHashSet factory that creates a hash set
func NewHashSet(values ...string) *HashSet {
	set := HashSet{data: make(map[string]bool)}
	for _, value := range values {
		set.Add(value)
	}
	return &set
}

// HashSet datastructure
type HashSet struct {
	data map[string]bool
}

// Add adds values to the set
func (s *HashSet) Add(values ...string) {
	for _, value := range values {
		s.data[value] = true
	}
}

// Remove removes values from the set
func (s *HashSet) Remove(values ...string) {
	for _, value := range values {
		delete(s.data, value)
	}
}

// Contains checks if values are in the set
func (s *HashSet) Contains(values ...string) bool {
	for _, value := range values {
		if !s.data[value] {
			return false
		}
	}
	return true
}

// Merge the two sets
func (s *HashSet) Merge(sets ...*HashSet) {
	for _, set := range sets {
		for _, value := range set.List() {
			s.Add(value)
		}
	}
}

// Clear clears set
func (s *HashSet) Clear() {
	s.data = make(map[string]bool)
}

// List returns values
func (s *HashSet) List() []string {
	values := make([]string, 0)
	for key := range s.data {
		values = append(values, key)
	}
	return values
}

// Size returns size of the set
func (s *HashSet) Size() int {
	return len(s.data)
}
