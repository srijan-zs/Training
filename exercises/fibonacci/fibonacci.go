package main

import "fmt"

// fibonacci function returns the fibonacci series till the given rows as a slice of int
func fibonacci(till int) []int {
	series := make([]int, 0)
	first, second := 0, 1
	series = append(series, first, second)

	for i := 2; i < till; i++ {
		first, second = second, first+second
		series = append(series, second)
	}

	return series
}

func main() {
	var till int

	fmt.Scanf("%d", &till)
	fmt.Println(fibonacci(till))
}
