package wordbreakers

import "github.com/shomali11/go-interview/datastructures/sets/hashsets"

const (
	empty = ""
)

// BreakWord breaks into possible words
func BreakWord(input string, values []string) string {
	dictionary := hashsets.NewHashSet(values...)
	return breakWord(input, dictionary, make(map[string]string))
}

func breakWord(input string, dictionary *hashsets.HashSet, cache map[string]string) string {
	if dictionary.Contains(input) {
		return input
	}

	cachedValue, exists := cache[input]
	if exists {
		return cachedValue
	}

	inputRunes := []rune(input)
	for i := 1; i < len(input); i++ {
		prefix := string(inputRunes[0:i])
		if dictionary.Contains(prefix) {
			suffix := string(inputRunes[i:])
			brokenSuffix := breakWord(suffix, dictionary, cache)
			if len(brokenSuffix) > 0 {
				cache[input] = prefix + " " + brokenSuffix
				return prefix + " " + brokenSuffix
			}
		}
	}
	return empty
}

// ExtractWords extracts all possible words
func ExtractWords(input string, values []string) []string {
	dictionary := hashsets.NewHashSet(values...)
	results := extractWords(input, dictionary, make(map[string]string))
	return results.List()
}

func extractWords(input string, dictionary *hashsets.HashSet, cache map[string]string) *hashsets.HashSet {
	results := hashsets.NewHashSet()
	if dictionary.Contains(input) {
		results.Add(input)
	}

	cachedValue, exists := cache[input]
	if exists {
		results.Add(cachedValue)
		return results
	}

	inputRunes := []rune(input)
	for i := 1; i < len(input); i++ {
		prefix := string(inputRunes[0:i])
		if dictionary.Contains(prefix) {
			suffix := string(inputRunes[i:])
			brokenSuffixes := extractWords(suffix, dictionary, cache)
			for _, brokenSuffix := range brokenSuffixes.List() {
				cache[input] = prefix + " " + brokenSuffix
				results.Add(prefix + " " + brokenSuffix)
			}
		}
	}
	return results
}
