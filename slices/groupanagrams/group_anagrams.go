package groupanagrams

import (
	"github.com/shomali11/go-interview/datastructures/maps/hashmultimaps"
	"github.com/shomali11/go-interview/strings/sorts"
)

// GroupAnagrams groups anagrams
func GroupAnagrams(list []string) *hashmultimaps.HashMultiMap[string, string] {
	multimap := hashmultimaps.New[string, string]()
	for _, value := range list {
		sortedValue := sorts.Sort(value)
		multimap.Put(sortedValue, value)
	}
	return multimap
}
