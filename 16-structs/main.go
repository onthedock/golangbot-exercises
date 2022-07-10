package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	var emp1 Employee
	emp1.firstName = "John"
	emp1.lastName = "Doe"
	fmt.Println("First name", emp1.firstName)
	fmt.Println("Last name", emp1.lastName)
	fmt.Println("Age", emp1.age)
}
