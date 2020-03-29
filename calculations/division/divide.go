package division

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

	value := a / b
	remainder := a % b
	return fmt.Sprintf(resultFormat, value, divide(remainder, b))
}

func divide(a int, b int) string {
	remainderIndexMap := make(map[int]int)
	values := make([]int, 0)

	value := 0
	remainder := a
	for !exists(remainderIndexMap, remainder) {
		remainderIndexMap[remainder] = len(values)

		remainder *= 10
		value = remainder / b
		remainder = remainder % b
		values = append(values, value)
	}

	index, _ := remainderIndexMap[remainder]

	var buffer bytes.Buffer
	for i := 0; i < index; i++ {
		buffer.WriteString(strconv.Itoa(values[i]))
	}

	buffer.WriteString(leftParenthesis)

	for i := index; i < len(values); i++ {
		buffer.WriteString(strconv.Itoa(values[i]))
	}

	buffer.WriteString(rightParenthesis)
	return buffer.String()
}

func exists(myMap map[int]int, key int) bool {
	_, ok := myMap[key]
	return ok
}
