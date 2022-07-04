package main

import "fmt"

func main() {
	a := [...]float64{67.7, 89.8, 21, 78}
	var sum float64 = 0

	for n, v := range a { // range returns the index and the value of the array
		fmt.Printf("The %dth element of the array is %.2f\n", n, v)
		sum += v
	}
	fmt.Printf("The sum of all elements is %.2f\n", sum)
}
