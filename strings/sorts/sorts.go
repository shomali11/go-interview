package sorts

import "sort"

// Sort sort a string
func Sort(text string) string {
	runes := []rune(text)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
