package main

import (
	"fmt"
)

func number() int {
	num := 15 * 5 // num = 75
	return num
}

func main() {
	switch num := number(); { //num is not a constant
	case num < 50: // this evaluates to true
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100: // this results in 'true'
		fmt.Printf("%d is lesser than 100\n", num) // print the message
		fallthrough                                // continue evaluating 'case' statements
	case num < 200: // it is not checked (because of the previous fallthrough)!!!
		fmt.Printf("%d is lesser than 200\n", num) // print the message
	}
}
