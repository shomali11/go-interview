package hashmultimaps

// New factory that creates a new Hash Multi Map
func New() *HashMultiMap {
	multiMap := HashMultiMap{data: make(map[interface{}][]interface{})}
	return &multiMap
}

// HashMultiMap a data structure representing a map of keys with lists of values
type HashMultiMap struct {
	data map[interface{}][]interface{}
}

// Merge merge multiple multi maps
func (s *HashMultiMap) Merge(maps ...*HashMultiMap) {
	for _, multiMap := range maps {
		for _, key := range multiMap.GetKeys() {
			values := multiMap.GetValues(key)
			s.PutAll(key, values...)
		}
	}
}

// Put key/value pair to the multi map
func (s *HashMultiMap) Put(key interface{}, value interface{}) {
	values, ok := s.data[key]
	if !ok {
		values = make([]interface{}, 0)
	}
	values = append(values, value)
	s.data[key] = values
}

// PutAll put key/values to the multi map
func (s *HashMultiMap) PutAll(key interface{}, values ...interface{}) {
	for _, value := range values {
		s.Put(key, value)
	}
}

// GetKeys returns a list of the multi map's keys
func (s *HashMultiMap) GetKeys() []interface{} {
	keys := make([]interface{}, 0, s.Size())
	for key := range s.data {
		keys = append(keys, key)
	}
	return keys
}

// Contains checks if a key is in the multi map
func (s *HashMultiMap) Contains(key interface{}) bool {
	_, exists := s.data[key]
	return exists
}

// ContainsAll checks if all keys are in the multi map
func (s *HashMultiMap) ContainsAll(keys ...interface{}) bool {
	for _, key := range keys {
		if !s.Contains(key) {
			return false
		}
	}
	return true
}

// ContainsAny checks if any keys are in the multi map
func (s *HashMultiMap) ContainsAny(keys ...interface{}) bool {
	for _, key := range keys {
		if s.Contains(key) {
			return true
		}
	}
	return false
}

// GetValues returns values associated with the key
func (s *HashMultiMap) GetValues(key interface{}) []interface{} {
	values, ok := s.data[key]
	if !ok {
		return make([]interface{}, 0)
	}
	return values
}

// RemoveKey removes a key and all its values
func (s *HashMultiMap) RemoveKey(keys ...interface{}) {
	for _, key := range keys {
		delete(s.data, key)
	}
}

// Remove removes a value from a key's values
func (s *HashMultiMap) Remove(key interface{}, value interface{}) {
	values, ok := s.data[key]
	if !ok {
		return
	}

	index := getIndex(value, values...)
	if index == -1 {
		return
	}

	newValues := make([]interface{}, 0)
	newValues = append(newValues, values[:index]...)
	s.data[key] = append(newValues, values[index+1:]...)
}

// Clear clears the multiMap
func (s *HashMultiMap) Clear() {
	s.data = make(map[interface{}][]interface{})
}

// IsEmpty checks if the multiMap is empty
func (s *HashMultiMap) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the multiMap
func (s *HashMultiMap) Size() int {
	return len(s.data)
}

func getIndex(value interface{}, values ...interface{}) int {
	for i, v := range values {
		if v == value {
			return i
		}
	}
	return -1
}
