// an interface is a set of method signatures. When a type provides definition for all the methods in the interface, it is said to implement the interface.
// WashingMachine can be an interface with method signatures Cleaning() and Drying(). Any type which provides definition for Cleaning() and Drying() methods is said to implement the WashingMachine interface

package main

import (
	"fmt"
	"unicode"
)

type VowelsFinder interface {
	FindVowels() []rune
}
type myString string

// MyString implements VowelsFinder
func (ms myString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		rune = unicode.ToLower(rune)
		switch rune {
		case 'a', 'e', 'i', 'o', 'u':
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	var name myString = myString("Antonio Machin")
	fmt.Printf("Vowels are %c\n", name.FindVowels())
}
