package main

import "fmt"

func main() {
	employeeAge := map[string]int{
		"John": 33,
		"Jane": 32,
	}
	employee := "John"
	age := employeeAge[employee]
	fmt.Printf("%s is %d years old\n", employee, age)
}
