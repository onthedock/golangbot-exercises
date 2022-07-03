package main

import (
	"fmt"
)

func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width            // declaration in the function signature
	perimeter = (length + width) * 2 // declaration in the function signature
	return                           // no need to specify the name of the returned variables
}

func main() {
	area, _ := rectProps(10.8, 5.6) // The returned value for `perimeter` is assigned to the blank identifier
	fmt.Printf("Area %f\n", area)   // this allows to "ignore" returned values
}
