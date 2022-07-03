package main

import (
	"fmt"
)

func main() {
	for number, index := 10, 1; index <= 10 && number <= 19; index, number = index+1, number+1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", number, index, number*index)
	}
}
