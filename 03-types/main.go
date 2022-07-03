package main

import "fmt"

func main() {
	i := 3       // int
	j := 2.3     // float
	sum := i + j // ERROR: mismatched types int and float64
	fmt.Println(sum)
}
