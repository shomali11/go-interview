package primes

import (
	"math"
)

// IsPrime determines if a number is prime
func IsPrime(number int) bool {
	if number <= 1 {
		return false
	}

	if number%2 == 0 {
		return false
	}

	for i := 3; i <= sqrt(number); i = i + 2 {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func sqrt(number int) int {
	return int(math.Sqrt(float64(number)))
}
