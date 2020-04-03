package reversestrings

// ReverseString reverse a string
func ReverseString(text string) string {
	runes := []rune(text)
	Reverse(runes, 0, len(runes)-1)
	return string(runes)
}

// Reverse reverse a string
func Reverse(runes []rune, start int, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
}
