package reverse

// Reverse reverse a string
func Reverse(text string) string {
	runes := []rune(text)
	length := len(runes)
	for i, j := 0, length-1; i < length/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
