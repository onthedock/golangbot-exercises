package main

import "fmt"

func main() {
	employeeAge := map[string]int{
		"John":    33,
		"Jane":    32,
		"Marc":    45,
		"Anthony": 55,
	}
	for k, v := range employeeAge {
		fmt.Printf("The employee '%s' is '%d' years old\n", k, v)
	}
}
