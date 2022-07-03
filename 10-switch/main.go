package main

import "fmt"

func main() {
	letter := "i"
	fmt.Printf("Letter %s is a ", letter)
	switch letter {
	case "a", "e", "i", "o", "u": // multiple expresions in case are allowed
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}
}
