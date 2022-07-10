package main

import "fmt"

func changeVal(val *int) {
	*val = 55
}

func main() {
	a := 58
	fmt.Println("value of 'a' before function call", a)
	b := &a
	changeVal(b)
	fmt.Println("Value of 'a' after function call", a)
}
