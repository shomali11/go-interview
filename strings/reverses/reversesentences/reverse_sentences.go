package reversesentences

import "github.com/shomali11/go-interview/strings/reverses/reversestrings"

// ReverseSentence reverse a sentence
func ReverseSentence(sentence string) string {
	runes := []rune(sentence)
	size := len(runes)

	// Reverse entire string
	reversestrings.Reverse(runes, 0, size-1)

	start := 0
	end := 0

	for start < size && end < size {
		if runes[start] == ' ' {
			start++
			end++
			continue
		}

		if runes[end] != ' ' {
			end++
			continue
		}

		reversestrings.Reverse(runes, start, end-1)
		start = end
	}

	reversestrings.Reverse(runes, start, end-1)
	return string(runes)
}
