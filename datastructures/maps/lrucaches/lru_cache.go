package lrucaches

import (
	"container/list"
)

// New factory that creates a new LRU Cache
func New(capacity int) *LRUCache {
	multiMap := LRUCache{
		capacity:   capacity,
		linkedList: list.New(),
		hashMap:    make(map[interface{}]*list.Element),
	}
	return &multiMap
}

// LRUEntry holds a key/value pair
type LRUEntry struct {
	Key   interface{}
	Value interface{}
}

// LRUCache a data structure representing a map of keys with lists of values
type LRUCache struct {
	capacity   int
	linkedList *list.List
	hashMap    map[interface{}]*list.Element
}

// Merge merge multiple lru caches
func (s *LRUCache) Merge(caches ...*LRUCache) {
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
func (s *LRUCache) Put(key interface{}, value interface{}) {
	if s.capacity == s.Size() {
		element := s.linkedList.Back()
		if element == nil {
			return
		}
		s.remove(element.Value.(*LRUEntry).Key)
	}

	lruEntry := &LRUEntry{Key: key, Value: value}
	listElement := s.linkedList.PushFront(lruEntry)
	s.hashMap[key] = listElement
}

// GetKeys returns a list of the cache's keys in most recently used order
func (s *LRUCache) GetKeys() []interface{} {
	keys := make([]interface{}, 0, s.Size())
	currentListElement := s.linkedList.Front()
	for currentListElement != nil {
		entry := currentListElement.Value.(*LRUEntry)
		keys = append(keys, entry.Key)
		currentListElement = currentListElement.Next()
	}
	return keys
}

// GetEntries returns a list of the cache's entries in most recently used order
func (s *LRUCache) GetEntries() []*LRUEntry {
	entries := make([]*LRUEntry, 0, s.Size())
	currentListElement := s.linkedList.Front()
	for currentListElement != nil {
		entry := currentListElement.Value.(*LRUEntry)
		entries = append(entries, entry)
		currentListElement = currentListElement.Next()
	}
	return entries
}

// Contains checks if a key is in cache
func (s *LRUCache) Contains(key interface{}) bool {
	_, exists := s.hashMap[key]
	return exists
}

// ContainsAll checks if all keys are in the cache
func (s *LRUCache) ContainsAll(keys ...interface{}) bool {
	for _, key := range keys {
		if !s.Contains(key) {
			return false
		}
	}
	return true
}

// ContainsAny checks if any keys are in the cache
func (s *LRUCache) ContainsAny(keys ...interface{}) bool {
	for _, key := range keys {
		if s.Contains(key) {
			return true
		}
	}
	return false
}

// GetValue returns value associated with the key
func (s *LRUCache) GetValue(key interface{}) (interface{}, bool) {
	listElement, ok := s.hashMap[key]
	if !ok {
		return nil, false
	}

	s.linkedList.MoveToFront(listElement)
	lruEntry := listElement.Value.(*LRUEntry)
	return lruEntry.Value, true
}

// Remove removes a key and its value
func (s *LRUCache) Remove(keys ...interface{}) {
	for _, key := range keys {
		s.remove(key)
	}
}

// Clear clears the multiMap
func (s *LRUCache) Clear() {
	s.hashMap = make(map[interface{}]*list.Element)
	s.linkedList.Init()
}

// IsEmpty checks if the multiMap is empty
func (s *LRUCache) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the multiMap
func (s *LRUCache) Size() int {
	return len(s.hashMap)
}

func (s *LRUCache) remove(key interface{}) {
	listElement, ok := s.hashMap[key]
	if !ok {
		return
	}

	s.linkedList.Remove(listElement)
	delete(s.hashMap, key)
}
