package hashmultimaps

// New factory that creates a new Hash Multi Map
func New[K, V comparable]() *HashMultiMap[K, V] {
	multiMap := HashMultiMap[K, V]{data: make(map[K][]V)}
	return &multiMap
}

// HashMultiMap a data structure representing a map of keys with lists of values
type HashMultiMap[K, V comparable] struct {
	data map[K][]V
}

// Merge merge multiple multi maps
func (s *HashMultiMap[K, V]) Merge(maps ...*HashMultiMap[K, V]) {
	for _, multiMap := range maps {
		for _, key := range multiMap.GetKeys() {
			values := multiMap.GetValues(key)
			s.PutAll(key, values...)
		}
	}
}

// Put key/value pair to the multi map
func (s *HashMultiMap[K, V]) Put(key K, value V) {
	values, ok := s.data[key]
	if !ok {
		values = make([]V, 0)
	}
	values = append(values, value)
	s.data[key] = values
}

// PutAll put key/values to the multi map
func (s *HashMultiMap[K, V]) PutAll(key K, values ...V) {
	for _, value := range values {
		s.Put(key, value)
	}
}

// GetKeys returns a list of the multi map's keys
func (s *HashMultiMap[K, V]) GetKeys() []K {
	keys := make([]K, 0, s.Size())
	for key := range s.data {
		keys = append(keys, key)
	}
	return keys
}

// Contains checks if a key is in the multi map
func (s *HashMultiMap[K, V]) Contains(key K) bool {
	_, exists := s.data[key]
	return exists
}

// ContainsAll checks if all keys are in the multi map
func (s *HashMultiMap[K, V]) ContainsAll(keys ...K) bool {
	for _, key := range keys {
		if !s.Contains(key) {
			return false
		}
	}
	return true
}

// ContainsAny checks if any keys are in the multi map
func (s *HashMultiMap[K, V]) ContainsAny(keys ...K) bool {
	for _, key := range keys {
		if s.Contains(key) {
			return true
		}
	}
	return false
}

// GetValues returns values associated with the key
func (s *HashMultiMap[K, V]) GetValues(key K) []V {
	values, ok := s.data[key]
	if !ok {
		return make([]V, 0)
	}
	return values
}

// RemoveKey removes a key and all its values
func (s *HashMultiMap[K, V]) RemoveKey(keys ...K) {
	for _, key := range keys {
		delete(s.data, key)
	}
}

// Remove removes a value from a key's values
func (s *HashMultiMap[K, V]) Remove(key K, value V) {
	values, ok := s.data[key]
	if !ok {
		return
	}

	index := getIndex(value, values...)
	if index == -1 {
		return
	}

	newValues := make([]V, 0)
	newValues = append(newValues, values[:index]...)
	s.data[key] = append(newValues, values[index+1:]...)
}

// Clear clears the multiMap
func (s *HashMultiMap[K, V]) Clear() {
	s.data = make(map[K][]V)
}

// IsEmpty checks if the multiMap is empty
func (s *HashMultiMap[K, V]) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the multiMap
func (s *HashMultiMap[K, V]) Size() int {
	return len(s.data)
}

func getIndex[V comparable](value V, values ...V) int {
	for i, v := range values {
		if v == value {
			return i
		}
	}
	return -1
}
