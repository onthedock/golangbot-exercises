package main

import (
	"fmt"
)

func main() {
	switch num := 25; {
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough // When fallthrough is encountered the control moves to the first statement of the next case
	case num > 100: // skipped because of the fallthrough!!!
		fmt.Printf("%d is greater than 100\n", num) // this is printed!!!
	}
}
