package main

import "fmt"

// countOccurrence function returns the map containing the count of occurrences of each character in a string
func countOccurrence(word string) map[string]int {
	count := make(map[string]int)

	for _, char := range word {
		count[string(char)]++
	}

	return count
}

func main() {
	var word string

	fmt.Scanf("%s", &word)
	fmt.Println(countOccurrence(word))
}
