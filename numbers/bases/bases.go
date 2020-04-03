package bases

import (
	"bytes"
	"math"
	"strings"

	"github.com/shomali11/go-interview/strings/reverses/reversestrings"
)

const (
	codes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
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
	return reversestrings.ReverseString(buffer.String())
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
