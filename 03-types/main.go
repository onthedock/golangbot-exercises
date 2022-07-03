package main

import "fmt"

func main() {
	i := 3                // int
	j := 2.3              // float
	sum := float32(i) + j // ERROR: infered type for j is 'float64'!!!
	fmt.Println(sum)
}
