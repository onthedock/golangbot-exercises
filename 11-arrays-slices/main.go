package main

import "fmt"

func main() {
	a := []int{12, 21, 32, 48, 95}
	var b []int = a[1:3] // gets items from the element a[1] to the element befon a[3]=a[3-1]=a[2]
	fmt.Println(b)
}
