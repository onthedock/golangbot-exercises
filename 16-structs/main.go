package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	var emp1 Employee // Not initialized -> zero values
	fmt.Println("First name", emp1.firstName)
	fmt.Println("Last name", emp1.lastName)
	fmt.Println("Age", emp1.age)
	emp1.firstName = "Samuel"
	fmt.Println("New name", emp1.firstName)
}
