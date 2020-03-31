package repeatingfractions

import (
	"bytes"
	"fmt"
	"strconv"
)

const (
	resultFormat     = "%d.%s"
	errorResult      = "ERROR"
	leftParenthesis  = "("
	rightParenthesis = ")"
)

// Divide divides two integer values and displays the repeating part of the fraction in parenthesis.
func Divide(a int, b int) string {
	if b == 0 {
		return errorResult
	}

	quotient := a / b
	remainder := a % b
	fractionString := divide(remainder, b)

	// If no fraction result
	if len(fractionString) == 0 {
		return fmt.Sprint(quotient)
	}
	return fmt.Sprintf(resultFormat, quotient, fractionString)
}

// divide returns the fraction with repeating parts between paranthesis
func divide(a int, b int) string {
	remainderIndexMap := make(map[int]int)
	digits := make([]int, 0)

	digit := 0
	remainder := a

	// While we have not seen that digit before
	for !exists(remainderIndexMap, remainder) {
		remainderIndexMap[remainder] = len(digits)

		remainder *= 10
		digit = remainder / b
		remainder = remainder % b
		digits = append(digits, digit)
	}

	repeatingDigitIndex := remainderIndexMap[remainder]
	return concat(digits, repeatingDigitIndex)
}

// concat create the string result
func concat(digits []int, repeatingDigitIndex int) string {
	var buffer bytes.Buffer
	for i := 0; i < repeatingDigitIndex; i++ {
		buffer.WriteString(strconv.Itoa(digits[i]))
	}

	// If the repeating result is a 0, return
	if repeatingDigitIndex == len(digits)-1 && digits[repeatingDigitIndex] == 0 {
		return buffer.String()
	}

	// Collect repeating digits between paranthesis
	buffer.WriteString(leftParenthesis)
	for i := repeatingDigitIndex; i < len(digits); i++ {
		buffer.WriteString(strconv.Itoa(digits[i]))
	}
	buffer.WriteString(rightParenthesis)
	return buffer.String()
}

// exists Determines whether a key exists
func exists(myMap map[int]int, key int) bool {
	_, ok := myMap[key]
	return ok
}
