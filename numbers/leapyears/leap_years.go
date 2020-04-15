package leapyears

// IsLeapYear determines if a year is a leap year
func IsLeapYear(number int) bool {
	if number%4 != 0 {
		return false
	}

	if number%100 != 0 {
		return true
	}

	return number%400 == 0
}
