package excels

import (
	"bytes"
	"strings"

	"github.com/shomali11/go-interview/strings/reverses/reversestrings"
)

const (
	characterA     = 'A'
	characterZ     = 'Z'
	alphabetsCount = 26
)

// Encode converts a number to an Excel Column Name
func Encode(number int) string {
	var buffer bytes.Buffer
	remainder := 0
	for number > 0 {
		remainder = (number - 1) % alphabetsCount
		buffer.WriteString(string(characterA + remainder))
		number = (number - remainder) / alphabetsCount
	}
	return reversestrings.ReverseString(buffer.String())
}

// Decode converts an Excel Column Name to a number
func Decode(value string) int {
	value = strings.ToUpper(value)
	number := 0
	pow := 1
	for i := len(value) - 1; i >= 0; i-- {
		if value[i] < characterA || value[i] > characterZ {
			continue
		}

		number += (int(value[i]-characterA) + 1) * pow
		pow *= alphabetsCount
	}
	return number

}
