package main

import "fmt"

func main() {
	number := 75
	switch { // the switch is considered 'true'
	case number >= 0 && number < 50: // each condition in the case statement is evaluated
		fmt.Printf("%d is greater or equal than 0, but less than 50\n", number)
	case number >= 50 && number < 100:
		fmt.Printf("%d is greater or equal than 50, but less than 100\n", number)
	default:
		fmt.Printf("%d is greater than 100\n", number)
	}
}
