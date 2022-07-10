package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	emp1 := &Employee{
		firstName: "Samuel",
		lastName:  " Beckett",
		age:       25,
	}
	fmt.Println("First name:", (*emp1).firstName)
	fmt.Println("Age:", (*emp1).age)
}
