package methods

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadString: to read a string from console input

// PrintString: to return the string in upper case.

type Rp struct {
	Result string
}

// ReadString: to read a string from console input
func (r *Rp) ReadString() string {
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please insert a strings: ")
	input.Scan()
	err := input.Err()
	if err != nil {
		fmt.Printf("Error! %v", err.Error())
	}
	result := input.Text()
	return result
}

func (r *Rp) PrintString(input string) string {
	result := strings.ToUpper(input)
	return result
}
