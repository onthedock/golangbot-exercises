package main

import (
	"fmt"
)

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func main() {
	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
	fmt.Printf("\n\n")
	name = "Señor"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
	// Output
	// String: Hello World
	// Characters: H e l l o   W o r l d
	// Bytes: 48 65 6c 6c 6f 20 57 6f 72 6c 64

	// String: Señor
	// Characters: S e Ã ± o r <-- FAIL!!!
	// // Bytes: 53 65 c3 b1 6f 72
}
