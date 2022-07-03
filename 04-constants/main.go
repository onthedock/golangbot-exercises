package main

import "fmt"

func main() {
	const (
		name  = "Xavi"
		age   = 123
		color = "Green"
	)
	fmt.Println("My name is", name, "and I am", age, "years old. I like the color", color)
	color = "Red" // Error: the value of a constant cannot be changed: it's constant ;)
}
