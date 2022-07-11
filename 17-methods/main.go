package main

import "fmt"

// A method is just a function with a special receiver type between the `func` keyworkd and the method name. The receiver can be a struct type or a non-struct type

// func (t Type) methodName(parameter list) {...}

type Employee struct {
	name     string
	salary   int
	currency string
}

func displaySalary(e Employee) {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func main() {
	emp1 := Employee{
		name:     "Sara Connor",
		salary:   5000,
		currency: "$",
	}
	displaySalary(emp1) // Call displayMethod()
}
