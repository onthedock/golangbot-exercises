package main

import "fmt"

func main() {
	employeeAge := map[string]int{
		"John": 33,
		"Jane": 32,
	}
	employee := "Xavi"
	age, ok := employeeAge[employee] // returns the value and ok=true if the 'key' exists in the map
	if !ok {
		fmt.Printf("The age of employee %s is not defined in the map\n", employee)
		return
	}
	fmt.Printf("%s is %d years old\n", employee, age)
}
