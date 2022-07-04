package main

import (
	"fmt"
)

func main() {
	darr := [...]int{57, 89, 78, 67, 69, 59} // declares an array
	dslice := darr[2:5]                      // creates a slice (from the previous array)
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++ // modifies the values in the slice. As the slice is just a reference, the underlying array its updated.
	}
	fmt.Println("array after", darr)
}
