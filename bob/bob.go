package main

// importing required packages
import (
	"fmt"
	"strings"
)

// myFunction for bob problem using if-else nested condition

// func myFunction(input string) interface{} {
//	if len(strings.TrimSpace(input)) == 0 {
//		return "Fine. Be that way!"
//	} else if strings.ToUpper(input) == input {
//		if input[len(input)-1] == '?' {
//			return "Calm down, I know what I'm doing!"
//		}
//		return "Whoa, chill out!"
//	} else if input[len(input)-1] == '?' {
//		return "Sure"
//	} else {
//		return "Whatever"
//	}
//}

// myFunction checks for the given conditions as per the question
func myFunction(input string) interface{} {
	switch {
	// check for empty strings
	case strings.TrimSpace(input) == "":
		return "Fine. Be that way!"
	// check for yelling questions
	case strings.ToUpper(input) == input && input[len(input)-1] == '?':
		return "Calm down, I know what I'm doing!"
	// check for yelling statements
	case strings.ToUpper(input) == input:
		return "Whoa, chill out!"
	// check for questions
	case input[len(input)-1] == '?':
		return "Sure"
	default:
		// default response
		return "Whatever"
	}
}

// main provides the test cases and calls the myFunction for the respective outcome
func main() {
	fmt.Println(myFunction("You there?"))       // question case
	fmt.Println(myFunction("HEYYYYYY"))         // yelling case
	fmt.Println(myFunction("WHATS YOUR NAME?")) // yelling question
	fmt.Println(myFunction("               "))  // empty string case
	fmt.Println(myFunction("What's happening")) // whatever case
}
