package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // do not print 'i' if it is even
		}
		fmt.Printf("%d ", i)
	}
}
