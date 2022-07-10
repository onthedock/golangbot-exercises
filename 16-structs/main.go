package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	// creates a struct specifying field names
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		lastName:  "Anderson",
	}
	// creates a struct without specifying field names
	emp2 := Employee{"Federico", "Mercurio", 33}
	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)
}
