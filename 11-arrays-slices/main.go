package main

func main() {
	a := [3]int{12, 78, 50}
	var b [5]int
	b = a // Error: The size of the array is a part of the type; therefore, they are different types and 'a' cannot be assigned to 'b'
}
