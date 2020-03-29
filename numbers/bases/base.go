package bases

import (
	"bytes"
	"math"
	"strings"
)

const (
	codes = "0123456789ABCDEF"
)

// Encode converts a number to a given base
func Encode(number int, base int) string {
	var buffer bytes.Buffer
	remainder := 0
	for number > 0 {
		remainder = number % base
		buffer.WriteString(string(codes[remainder]))
		number /= base
	}
	return reverse(buffer.String())
}

// Decode converts a number of a specific base
func Decode(number string, base int) int {
	sum := 0
	exponent := len(number) - 1
	digits := []rune(number)
	for i := 0; i < len(number); i++ {
		index := float64(strings.Index(codes, string(digits[i])))
		power := math.Pow(float64(base), float64(exponent))
		sum += int(index * power)
		exponent--
	}
	return int(sum)
}

func reverse(text string) string {
	runes := []rune(text)
	length := len(runes)
	for i, j := 0, length-1; i < length/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
