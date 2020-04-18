package palindromestrings

// IsPalindrome determines if the input is a palindrome
func IsPalindrome(number int) bool {
	copy := number
	reverse := 0
	for number != 0 {
		reverse *= 10
		reverse += number % 10
		number /= 10
	}
	return copy == reverse
}
