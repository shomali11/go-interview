package addbinaries

import (
	"bytes"
	"errors"
	"strconv"

	"github.com/shomali11/go-interview/strings/reverses/reversestrings"
)

const (
	empty = ""
)

var (
	errRuneNotInt = errors.New("digit is not an integer")
)

// Add adds two binary string numbers
func Add(number1 string, number2 string) (string, error) {
	var result bytes.Buffer

	number1Runes := []rune(number1)
	number2Runes := []rune(number2)

	number1Index := len(number1) - 1
	number2Index := len(number2) - 1

	carry := 0

	for number1Index >= 0 && number2Index >= 0 {
		value1, err := getNumber(number1Runes, number1Index)
		if err != nil {
			return empty, err
		}

		value2, err := getNumber(number2Runes, number2Index)
		if err != nil {
			return empty, err
		}

		result.WriteString(strconv.Itoa((value1 + value2 + carry) % 2))
		carry = (value1 + value2 + carry) / 2

		number1Index--
		number2Index--
	}

	for ; number1Index >= 0; number1Index-- {
		value, err := getNumber(number1Runes, number1Index)
		if err != nil {
			return empty, err
		}

		result.WriteString(strconv.Itoa((value + carry) % 2))
		carry = (value + carry) / 2
	}

	for ; number2Index >= 0; number2Index-- {
		value, err := getNumber(number2Runes, number2Index)
		if err != nil {
			return empty, err
		}

		result.WriteString(strconv.Itoa((value + carry) % 2))
		carry = (value + carry) / 2
	}

	if carry > 0 {
		result.WriteString(strconv.Itoa(carry))
	}
	return reversestrings.ReverseString(result.String()), nil
}

func getNumber(numberRunes []rune, index int) (int, error) {
	number, err := characterToNumber(numberRunes[index])
	if err != nil {
		return 0, err
	}
	return number, nil
}

func characterToNumber(r rune) (int, error) {
	if '0' <= r && r <= '9' {
		return int(r) - '0', nil
	}
	return 0, errRuneNotInt
}
