package postfixes

import (
	"errors"
	"strconv"
	"strings"

	"github.com/shomali11/go-interview/datastructures/stacks/slicestacks"
)

var (
	errERR = errors.New("#ERR")
)

// Evaluate evaluates postfix expressions
func Evaluate(expression string) (float64, error) {
	stack := slicestacks.New()
	tokens := strings.Fields(expression)

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		number, ok := isNumber(token)
		if ok {
			stack.Push(number)
			continue
		}

		num2, err := pop(stack)
		if err != nil {
			return 0, err
		}

		num1, err := pop(stack)
		if err != nil {
			return 0, err
		}

		switch token {
		case "+":
			stack.Push(num1 + num2)
		case "-":
			stack.Push(num1 - num2)
		case "*":
			stack.Push(num1 * num2)
		case "/":
			if num2 == 0 {
				return 0, errERR
			}
			stack.Push(num1 / num2)
		default:
			return 0, errERR
		}
	}

	if stack.Size() != 1 {
		return 0, errERR
	}
	return pop(stack)
}

func isNumber(text string) (float64, bool) {
	number, err := strconv.ParseFloat(text, 64)
	return number, err == nil
}

func pop(stack *slicestacks.Stack) (float64, error) {
	numberInterface, err := stack.Pop()
	if err != nil {
		return 0, err
	}

	number, ok := numberInterface.(float64)
	if !ok {
		return 0, errERR
	}
	return number, nil
}
