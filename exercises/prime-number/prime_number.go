package main

import (
	"fmt"
	"math"
)

// using basic approach
// func checkPrime(num int) string {
//	for i := 2; i <= num/2; i++ {
//		if num%i == 0 {
//			return "Not Prime"
//		}
//	}
//
//	return "Prime"
//}

// checkPrime function using sq-root approach to check whether the number is prime or not
func checkPrime(num int) string {
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return "Not Prime"
		}
	}

	return "Prime"
}

func main() {
	var num int

	fmt.Scanf("%d", &num)
	fmt.Println(checkPrime(num))
}
