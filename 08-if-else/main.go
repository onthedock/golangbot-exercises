package main

import "fmt"

func main() {
	if num := 10; num%2 == 0 { // 'num' is only available inside the if...else statement
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}
