package main

import (
	"fmt"
	"unicode"
)

// countNonASCII function returns the map of strings to int which consists of count of occurrences of each non-ASCII character present in the string
func countNonASCII(str string) map[string]int {
	count := make(map[string]int)

	for _, c := range str {
		if c > unicode.MaxASCII {
			c := string(c)
			_, ok := count[c]

			if ok {
				count[c]++
			} else {
				count[c] = 1
			}
		}
	}

	return count
}

func main() {
	var str string

	fmt.Scanf("%s", &str)
	fmt.Println(countNonASCII(str))
}
