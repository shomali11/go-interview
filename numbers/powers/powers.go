package powers

// Power calculates the power
func Power(number int, exponent int) float64 {
	if exponent == 0 {
		return 1
	}

	if exponent == 1 {
		return float64(number)
	}

	if exponent < 0 {
		return 1.0 / Power(number, -exponent)
	}

	result := Power(number, exponent/2)
	if exponent%2 == 0 {
		return result * result
	}
	return float64(number) * result * result
}
