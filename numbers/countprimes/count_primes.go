package countprimes

// CountPrimes returns the number of prime numbers up to given number
func CountPrimes(number int) int {
	if number <= 1 {
		return 0
	}

	isPrimeNumbers := make([]bool, number+1)
	for i := range isPrimeNumbers {
		isPrimeNumbers[i] = true
	}

	for i := 2; i*i <= number; i++ {
		if isPrimeNumbers[i] {
			for j := i * i; j <= number; j += i {
				isPrimeNumbers[j] = false
			}
		}
	}

	count := 0
	for i := 2; i <= number; i++ {
		if isPrimeNumbers[i] {
			count++
		}
	}
	return count
}
