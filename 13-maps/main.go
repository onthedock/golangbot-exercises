package main

import "fmt"

func main() {
	employeeAge := make(map[string]int)
	employeeAge["John"] = 33
	employeeAge["Jane"] = 32
	employeeAge["Ann"] = 21
	fmt.Println(employeeAge) // prints -> map[Ann:21 Jane:32 John:33]
}
