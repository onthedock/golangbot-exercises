package main

import (
	"fmt"
)

type Employee struct {
	name string
	age  int
}

/*
Method with value receiver
*/
func (e Employee) changeName(newName string) {
	e.name = newName
}

/*
Method with pointer receiver
*/
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

func main() {
	e := Employee{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Println("Value receiver:")
	fmt.Printf("Employee name before change: %s\n", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("Employee name after change: %s\n", e.name)

	fmt.Println("\nPointer Receivers:")
	fmt.Printf("Employee age before change: %d\n", e.age)
	(&e).changeAge(51)
	fmt.Printf("Employee age after change: %d\n", e.age)
}
