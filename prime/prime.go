package prime

import "math"

// checkPrime function is defined to validate whether a given number is a prime or not
func checkPrime(num int) bool {
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

// findPrime function generates the slice of all the primes till the given number
func findPrime(num int) []int {
	if num <= 0 {
		return nil
	}

	primes := make([]int, 0)

	for i := 2; i <= num; i++ {
		if checkPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes
}
