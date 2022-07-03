package main

import "fmt"

func functionName(parameter1 string, parameter2 string) string {
	// function body
	return parameter1 + " " + parameter2
}

func main() {
	greeting := functionName("Hola", "Xavi")
	fmt.Println(greeting)
}
