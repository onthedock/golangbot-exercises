package main

import "fmt"

func main() {
	employeeAge := map[string]int{
		"John": 33,
		"Jane": 32,
		"Ann":  21,
	}
	fmt.Println(employeeAge) // prints -> map[Ann:21 Jane:32 John:33]
}
