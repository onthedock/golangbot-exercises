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
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f\n", area, perimeter)
}
