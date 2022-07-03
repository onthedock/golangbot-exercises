package main

import "fmt"

func main() {
	i := 3            // int
	j := 2.3          // float
	sum := i + int(j) // Explicit conversion to type 'int'
	fmt.Println(sum)  // prints 5
}
