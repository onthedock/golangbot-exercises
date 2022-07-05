package main

import (
	"fmt"
)

func charsAndBytePosition(s string) {
	for index, rune := range s { // returns the position of the byte where the rune starts along with the rune
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func main() {
	name := "Señor"
	charsAndBytePosition(name)
}

// Output:
// S starts at byte 0 <-- 1 byte
// e starts at byte 1 <-- 1 byte
// ñ starts at byte 2 <-- 2 bytes (2 and 3) are used by 'ñ'
// o starts at byte 4 ...
// r starts at byte 5
