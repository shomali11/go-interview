package duplicates

import "github.com/shomali11/go-interview/datastructures/sets/hashsets"

// ContainsDuplicates checks if the list contains duplicates
func ContainsDuplicates[T comparable](values ...T) bool {
	set := hashsets.New[T]()
	for _, value := range values {
		if set.Contains(value) {
			return true
		}
		set.Add(value)
	}
	return false
}
