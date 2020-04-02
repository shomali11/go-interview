package wordbreakers

import (
	"fmt"

	"github.com/shomali11/go-interview/datastructures/sets/hashsets"
)

const (
	empty                   = ""
	stringFormat            = "%s"
	stringSpaceStringFormat = "%s %s"
)

// BreakWord breaks into possible words
func BreakWord(input string, values []string) string {
	dictionary := hashsets.New()
	for _, value := range values {
		dictionary.Add(value)
	}
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
				result := fmt.Sprintf(stringSpaceStringFormat, prefix, brokenSuffix)
				cache[input] = result
				return result
			}
		}
	}
	return empty
}

// ExtractWords extracts all possible words
func ExtractWords(input string, values []string) []string {
	dictionary := hashsets.New()
	for _, value := range values {
		dictionary.Add(value)
	}

	resultSet := extractWords(input, dictionary, make(map[string]string))

	results := make([]string, 0)
	for _, item := range resultSet.GetValues() {
		results = append(results, fmt.Sprintf(stringFormat, item))
	}
	return results
}

func extractWords(input string, dictionary *hashsets.HashSet, cache map[string]string) *hashsets.HashSet {
	resultSet := hashsets.New()
	if dictionary.Contains(input) {
		resultSet.Add(input)
	}

	cachedValue, exists := cache[input]
	if exists {
		resultSet.Add(cachedValue)
		return resultSet
	}

	inputRunes := []rune(input)
	for i := 1; i < len(input); i++ {
		prefix := string(inputRunes[0:i])
		if dictionary.Contains(prefix) {
			suffix := string(inputRunes[i:])
			brokenSuffixes := extractWords(suffix, dictionary, cache)
			for _, brokenSuffix := range brokenSuffixes.GetValues() {
				result := fmt.Sprintf(stringSpaceStringFormat, prefix, brokenSuffix)
				cache[input] = result
				resultSet.Add(result)
			}
		}
	}
	return resultSet
}
