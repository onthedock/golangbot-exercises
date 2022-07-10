package main

import "fmt"

type Address struct {
	city  string
	state string
}

type Person struct {
	name    string
	age     int
	address Address
}

func main() {
	p := Person{
		name: "Frank",
		age:  67,
		address: Address{
			city:  "Chicago",
			state: "Ilinois",
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)
}
