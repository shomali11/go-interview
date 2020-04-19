package palindromestrings

import "github.com/shomali11/go-interview/numbers/reverses"

// IsPalindrome determines if the input is a palindrome
func IsPalindrome(number int) bool {
	return number == reverses.Reverse(number)
}
