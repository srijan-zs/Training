package leapyear

// checkLeap function validates whether the given year is a leap year or not
func checkLeap(year int) bool {
	if year <= 0 {
		return false
	}

	if year%4 != 0 {
		return false
	}

	if year%100 == 0 && year%400 != 0 {
		return false
	}

	return true
}
