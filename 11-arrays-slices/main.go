package main

import "fmt"

func main() {
	// Arrays in Go are value types and not reference types.
	// This means that when they are assigned to a new variable,
	// a copy of the original array is assigned to the new variable.
	// If changes are made to the new variable, it will not be
	// reflected in the original array.
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
}
