package main

import "fmt"

func functionName(spanishGreeting, name string) string {
	// function body
	return spanishGreeting + " " + name
}

func main() {
	greeting := functionName("Hola", "Xavi")
	fmt.Println(greeting)
}
