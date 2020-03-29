package sets

// NewSet factory that creates a set
func NewSet(values ...string) *Set {
	set := Set{data: make(map[string]bool)}
	for _, value := range values {
		set.Add(value)
	}
	return &set
}

// Set datastructure
type Set struct {
	data map[string]bool
}

// Add adds a value to the set
func (s *Set) Add(value string) {
	s.data[value] = true
}

// Remove removes a vlaue from the set
func (s *Set) Remove(value string) {
	delete(s.data, value)
}

// Contains checks if a value is in the set
func (s *Set) Contains(value string) bool {
	return s.data[value]
}

// Merge the two sets
func (s *Set) Merge(set *Set) {
	for _, value := range set.List() {
		s.Add(value)
	}
}

// Clear clears set
func (s *Set) Clear() {
	s.data = make(map[string]bool)
}

// List returns values
func (s *Set) List() []string {
	values := make([]string, 0)
	for key := range s.data {
		values = append(values, key)
	}
	return values
}

// Size returns size of the set
func (s *Set) Size() int {
	return len(s.data)
}
