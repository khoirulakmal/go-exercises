package main

import "fmt"

// With a given integral number n, write a program to generate a map that contains (i, i*i)
// such that is an integral number between 1 and n (both included),
// and then the program should print the map with representation of the value

// Suppose the following input is supplied to the program: 8

// Then, the output should be: map[1:1 2:4 3:9 4:16 5:25 6:36 7:49 8:64]

func main() {
	num := make(map[int]int)
	var a int
	print("Please input a number:  ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		print("Please input a number!\n")
		return
	}
	for i := 1; i <= a; i++ {
		num[i] = i * i
	}
	fmt.Printf("%v\n", num)
}
