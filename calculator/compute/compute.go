package main

import (
	"fmt"

	"github.com/srijan-zs/calculator/calc"
)

func main() {
	fmt.Println(calc.Compute(23, 45, calc.Add))
	fmt.Println(calc.Compute(34, 12, calc.Sub))
	fmt.Println(calc.Compute(13, 7, calc.Multiply))
	fmt.Println(calc.Compute(45, 9, calc.Divide))

	// user specific needs
	var x, y float64

	var op string

	fmt.Print("Give 2 operands and an operator: ")
	fmt.Scanf("%f %f %s", &x, &y, &op)

	switch op {
	case "+":
		fmt.Println(calc.Compute(x, y, calc.Add))
	case "-":
		fmt.Println(calc.Compute(x, y, calc.Sub))
	case "*":
		fmt.Println(calc.Compute(x, y, calc.Multiply))
	case "/":
		fmt.Println(calc.Compute(x, y, calc.Divide))
	default:
		fmt.Println("Invalid parameters given!!!")
	}
}
