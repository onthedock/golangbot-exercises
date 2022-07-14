package main

import "fmt"

type customInt int

func (a customInt) add(b customInt) customInt {
	return a + b
}

func main() {
	num1 := customInt(5)
	num2 := customInt(20)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}
