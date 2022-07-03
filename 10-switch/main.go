package main

import "fmt"

const msg string = `A switch is a conditional statement that evaluates an expression and compares it
against a list of possible matches and executes the corresponding block of code.
It can be considered as an idiomatic way of replacing complex "if...else" clauses.
`

func main() {
	fmt.Println(msg)

	finger := 30
	fmt.Printf("Finger %d is ... ", finger)

	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("Invalid finger number")
	}
}
