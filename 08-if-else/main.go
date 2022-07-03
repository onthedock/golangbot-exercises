package main

import "fmt"

func main() {
	num := 10
	// Idiomatic Go avoids innecessary branching and indentation of code.
	// Also, return as early as possible.
	if num%2 == 0 {
		fmt.Println(num, "is even")
		return
	}
	fmt.Println(num, "is odd")
}
