package main

import "fmt"

func main() {
	number := 11
	if number%2 == 0 {
		fmt.Println("Number", number, "is even")
	} else {
		fmt.Println("Number", number, "is odd")
	}
}
