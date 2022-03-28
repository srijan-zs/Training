package main

import "fmt"

// checkPerfectNumber function validates whether the given number is a perfect number of not
func checkPerfectNumber(num int) string {
	divisors := make([]int, 5)

	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
		}
	}

	sum := 0

	for _, val := range divisors {
		sum += val
	}

	if sum == num {
		return "Is a perfect number"
	}

	return "Not a perfect number"
}

func main() {
	var num int

	fmt.Scanf("%d", &num)
	fmt.Println(checkPerfectNumber(num))
}
