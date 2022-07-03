package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break // the loop is terminated
			// execution continues after the 'for' loop
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n(break ends the loop)\nafter the loop")
}
