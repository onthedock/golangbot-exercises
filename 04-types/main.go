package main

import "fmt"

func main() {
	i := 3                // int
	j := 2.3              // float
	sum := float64(i) + j // Explicit type conversion
	fmt.Println(sum)      // prints 5.3
}
