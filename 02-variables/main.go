package main

import "fmt"

func main() {
	name, age := "Xavi", 123
	fmt.Println("My name is", name, "and I am", age, "years old")
	name, age := "Lisa", 321
	fmt.Println("My name is", name, "and I am", age, "years old")
}
