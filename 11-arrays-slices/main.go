package main

import "fmt"

func main() {
	a := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(a); i++ {
		fmt.Printf("The %dth element of the array is %.2f\n", i, a[i])
	}

}
