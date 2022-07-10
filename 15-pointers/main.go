package main

import "fmt"

func main() {
	size := new(int)
	fmt.Printf("Size value is %d, type is %T and address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New 'size' value is", *size)
}
