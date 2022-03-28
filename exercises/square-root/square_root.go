package main

import (
	"fmt"
)

// squareRoot function returns the square-root of the given value (float64)
func squareRoot(num float64) float64 {
	var tmp1, tmp2 float64

	tmp1 = num / 2

	for {
		tmp2 = tmp1
		tmp1 = (tmp2 + (num / tmp2)) / 2

		if tmp1 == tmp2 {
			break
		}
	}

	return tmp1
}

func main() {
	var num float64

	fmt.Scanf("%f", &num)
	fmt.Println(squareRoot(num))
}
