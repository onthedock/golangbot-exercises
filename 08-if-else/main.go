package main

import "fmt"

func main() {
	number := 10
	if number%2 == 0 {
		fmt.Println("Number", number, "is even")
		return
	}
	fmt.Println("Number", number, "is odd")
}
