package main

import (
	"fmt"
)

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) // %x is the format specifier for hexadecimal
	}
}

func main() {
	name := "Hello World"
	fmt.Printf("String: %s\n", name) // %s is the format specifier to print a string
	printBytes(name)                 // Prints --> Bytes: 48 65 6c 6c 6f 20 57 6f 72 6c 64
}
