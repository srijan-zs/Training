package main

import (
	"fmt"
	"unicode"
)

func swap(slice []rune, i, j int) []rune {
	if slice[i] <= unicode.MaxASCII {
		if slice[j] <= unicode.MaxASCII {
			slice[i], slice[j] = slice[j], slice[i]
		} else {
			j--
			swap(slice, i, j)
		}
	} else {
		i++
		swap(slice, i, j)
	}

	return slice
}

func reverse(str string) string {
	slice := []rune(str)

	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		swap(slice, i, j)
	}
	fmt.Println(slice)

	return string(slice)
}

func main() {
	var str string

	fmt.Scanf("%s", &str)
	fmt.Println([]rune(str))
	fmt.Println(reverse(str))
}
