package main

import (
	"exercise/05/methods"
	"fmt"
)

// Create a seperate file (module) which has at least two methods:

//     ReadString: to read a string from console input

//     PrintString: to return the string in upper case.

// Also create a main.go file that acts as calling class.

func main() {
	call := methods.Rp{
		Result: "",
	}
	text := call.ReadString()
	call.Result = call.PrintString(text)
	fmt.Print(call)
}
