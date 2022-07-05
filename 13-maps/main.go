package main

import "fmt"

type employee struct {
	country string
	age     int
}

func main() {
	e1 := employee{"USA", 21}
	e2 := employee{"Maxico", 19}
	e3 := employee{"France", 22}
	employeeInfo := map[string]employee{
		"Mick":     e1,
		"Miguel":   e2,
		"Delphine": e3,
	}
	fmt.Println(len(employeeInfo))
}
