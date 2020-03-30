package perfects

// IsPerfect determines if a number is perfect
func IsPerfect(number int) bool {
	if number <= 0 {
		return false
	}

	sum := 0
	for i := 1; i <= number/2; i++ {
		if number%i == 0 {
			sum += i
		}
	}
	return sum == number
}
