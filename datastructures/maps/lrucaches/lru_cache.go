package lrucaches

import (
	"container/list"
)

// New factory that creates a new LRU Cache
func New[K comparable, V any](capacity int) *LRUCache[K, V] {
	multiMap := LRUCache[K, V]{
		capacity:   capacity,
		linkedList: list.New(),
		hashMap:    make(map[K]*list.Element),
	}
	return &multiMap
}

// LRUEntry holds a key/value pair
type LRUEntry[K comparable, V any] struct {
	Key   K
	Value V
}

// LRUCache a data structure representing a map of keys with lists of values
type LRUCache[K comparable, V any] struct {
	capacity   int
	linkedList *list.List
	hashMap    map[K]*list.Element
}

// Merge merge multiple lru caches
func (s *LRUCache[K, V]) Merge(caches ...*LRUCache[K, V]) {
	for _, cache := range caches {
		for _, key := range cache.GetKeys() {
			value, ok := cache.GetValue(key)
			if !ok {
				continue
			}

			s.Put(key, value)
		}
	}
}

// Put key/value pair into the cache
func (s *LRUCache[K, V]) Put(key K, value V) {
	if s.capacity == s.Size() {
		element := s.linkedList.Back()
		if element == nil {
			return
		}
		s.remove(element.Value.(*LRUEntry[K, V]).Key)
	}

	lruEntry := &LRUEntry[K, V]{Key: key, Value: value}
	listElement := s.linkedList.PushFront(lruEntry)
	s.hashMap[key] = listElement
}

// GetKeys returns a list of the cache's keys in most recently used order
func (s *LRUCache[K, V]) GetKeys() []K {
	keys := make([]K, 0, s.Size())
	currentListElement := s.linkedList.Front()
	for currentListElement != nil {
		entry := currentListElement.Value.(*LRUEntry[K, V])
		keys = append(keys, entry.Key)
		currentListElement = currentListElement.Next()
	}
	return keys
}

// GetEntries returns a list of the cache's entries in most recently used order
func (s *LRUCache[K, V]) GetEntries() []*LRUEntry[K, V] {
	entries := make([]*LRUEntry[K, V], 0, s.Size())
	currentListElement := s.linkedList.Front()
	for currentListElement != nil {
		entry := currentListElement.Value.(*LRUEntry[K, V])
		entries = append(entries, entry)
		currentListElement = currentListElement.Next()
	}
	return entries
}

// Contains checks if a key is in cache
func (s *LRUCache[K, V]) Contains(key K) bool {
	_, exists := s.hashMap[key]
	return exists
}

// ContainsAll checks if all keys are in the cache
func (s *LRUCache[K, V]) ContainsAll(keys ...K) bool {
	for _, key := range keys {
		if !s.Contains(key) {
			return false
		}
	}
	return true
}

// ContainsAny checks if any keys are in the cache
func (s *LRUCache[K, V]) ContainsAny(keys ...K) bool {
	for _, key := range keys {
		if s.Contains(key) {
			return true
		}
	}
	return false
}

// GetValue returns value associated with the key
func (s *LRUCache[K, V]) GetValue(key K) (res V, found bool) {
	listElement, ok := s.hashMap[key]
	if !ok {
		return res, false
	}

	s.linkedList.MoveToFront(listElement)
	lruEntry := listElement.Value.(*LRUEntry[K, V])
	return lruEntry.Value, true
}

// Remove removes a key and its value
func (s *LRUCache[K, V]) Remove(keys ...K) {
	for _, key := range keys {
		s.remove(key)
	}
}

// Clear clears the multiMap
func (s *LRUCache[K, V]) Clear() {
	s.hashMap = make(map[K]*list.Element)
	s.linkedList.Init()
}

// IsEmpty checks if the multiMap is empty
func (s *LRUCache[K, V]) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the multiMap
func (s *LRUCache[K, V]) Size() int {
	return len(s.hashMap)
}

func (s *LRUCache[K, V]) remove(key K) {
	listElement, ok := s.hashMap[key]
	if !ok {
		return
	}

	s.linkedList.Remove(listElement)
	delete(s.hashMap, key)
}
