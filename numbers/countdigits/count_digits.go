package countdigits

// CountDigits counts the number of digits
func CountDigits(number int) int {
	if number == 0 {
		return 1
	}

	count := 0
	for number != 0 {
		number /= 10
		count++
	}
	return count
}
