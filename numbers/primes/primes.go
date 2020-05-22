package primes

// IsPrime determines if a number is prime
func IsPrime(number int) bool {
	if number <= 1 {
		return false
	}

	if number%2 == 0 {
		return false
	}

	for i := 3; i*i <= number; i = i + 2 {
		if number%i == 0 {
			return false
		}
	}
	return true
}
