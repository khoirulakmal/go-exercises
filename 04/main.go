package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Write a program which accepts a sequence of comma-separated numbers from console
// and generate an slice out of them. Return the slice.

// Suppose the following input is supplied to the program: 34, 67, 55, 33, 12, 98.

// Then, the output should be: [34 67 55 33 12 98]

func main() {
	print("Please input a number:  ")
	numbers, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Printf("Please enter correct number!\n %v", err.Error())
		return
	}
	fmt.Print(numbers)
	list := strings.Split(numbers, ",")
	fmt.Printf("%v\n", list)
}
