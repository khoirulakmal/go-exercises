package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Write a program which accepts a sequence of comma-separated numbers from console
// and generate an slice out of them. Return the slice.

// Suppose the following input is supplied to the program: 34, 67, 55, 33, 12, 98.

// Then, the output should be: [34 67 55 33 12 98]

func main() {
	a := play("34, 67, 55, 33, 12, 98")
	fmt.Printf("%v\n", a)
}

func play(input string) []int {
	var numbers []int
	list := strings.Split(input, ",")
	for _, num := range list {
		a := strings.TrimSpace(num)
		b, err := strconv.Atoi(a)
		if err != nil {
			fmt.Print(err.Error())
		}
		numbers = append(numbers, b)
	}
	return numbers
}
