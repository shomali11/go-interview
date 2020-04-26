package balancedparantheses

import (
	"github.com/shomali11/go-interview/datastructures/stacks/slicestacks"
)

const (
	openParanthesis1 = '['
	openParanthesis2 = '{'
	openParanthesis3 = '('

	closeParanthesis1 = ']'
	closeParanthesis2 = '}'
	closeParanthesis3 = ')'
)

// IsBalancedParantheses determines if the parantheses are balanced
func IsBalancedParantheses(expression string) bool {
	stack := slicestacks.New()

	for _, char := range []rune(expression) {
		if isOpenParantheses(char) {
			stack.Push(char)
			continue
		}

		if !isCloseParantheses(char) {
			continue
		}

		value, err := stack.Pop()
		if err != nil {
			return false
		}

		if !isMatchingParanthesis(value.(rune), char) {
			return false
		}
	}
	return stack.IsEmpty()
}

func isOpenParantheses(character rune) bool {
	return character == openParanthesis1 || character == openParanthesis2 || character == openParanthesis3
}

func isCloseParantheses(character rune) bool {
	return character == closeParanthesis1 || character == closeParanthesis2 || character == closeParanthesis3
}

func isMatchingParanthesis(open rune, close rune) bool {
	switch open {
	case openParanthesis1:
		return close == closeParanthesis1
	case openParanthesis2:
		return close == closeParanthesis2
	case openParanthesis3:
		return close == closeParanthesis3
	}
	return false
}
