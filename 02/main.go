package main

import (
	"fmt"
)

// Write a program which can compute the factorial of a given numbers.
// The results should be printed in a comma-separated sequence on a single line.

// Suppose the following input is supplied to the program: 8

// Then, the output should be: 40320

func main() {
	var input int
	print("Please input a number:  ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		print("Please input a number!\n")
		return
	}
	factorial(&input)
}

func factorial(num *int) {
	if *num < 0 {
		print("Input number larger than 0!\n")
		return
	}
	for i := *num - 1; i > 0; i-- {
		*num = *num * i
	}
	fmt.Printf("Factorial Output is: %v\n", *num)
}
