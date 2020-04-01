package duplicates

import "github.com/shomali11/go-interview/datastructures/sets/hashsets"

// ContainsDuplicates checks if the list contains duplicates
func ContainsDuplicates(values ...interface{}) bool {
	set := hashsets.New()
	for _, value := range values {
		if set.Contains(value) {
			return true
		}
		set.Add(value)
	}
	return false
}
