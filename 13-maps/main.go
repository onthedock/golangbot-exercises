package main

import "fmt"

func main() {
	employeeAge := map[string]int{
		"John":    33,
		"Jane":    32,
		"Marc":    45,
		"Anthony": 55,
	}
	fmt.Println("Initial map:", employeeAge)
	delete(employeeAge, "Marc")
	fmt.Println("Final map:", employeeAge)
}
