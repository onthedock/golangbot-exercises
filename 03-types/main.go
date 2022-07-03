package main

import "fmt"

func main() {
	int8Min := -128
	int8Max := 127
	fmt.Println("'int8' values range from", int8Min, "to", int8Max)
	int16Min := -32768
	int16Max := 32767
	fmt.Println("'int16' values range from", int16Min, "to", int16Max)
	int32Min := -2147483648
	int32Max := 2147483647
	fmt.Println("'int32' values range from", int32Min, "to", int32Max)
	int64Min := -9223372036854775808
	int64Max := 9223372036854775807
	fmt.Println("'int64' values range from", int64Min, "to", int64Max)
	fmt.Println("The 'int' type represents 32 or 64 bits integer depending on the platform")
	fmt.Println("In 32 bits platforms, 'int' is 'int32' and in 64 bits platforms, ' int64'")

	fmt.Println("")
	unsignedInt8Min := 0
	unsignedInt8Max := 255
	fmt.Println("'uint8' values range from", unsignedInt8Min, "to", unsignedInt8Max)
	unsignedInt16Min := 0
	unsignedInt16Max := 65535
	fmt.Println("'uint16' values range from", unsignedInt16Min, "to", unsignedInt16Max)
	unsignedInt32Min := 0
	unsignedInt32Max := 4294967295
	fmt.Println("'uint32' values range from", unsignedInt32Min, "to", unsignedInt32Max)
	unsignedInt64Min := 0
	var unsignedInt64Max uint64 = 18446744073709551615 // Inference failed to detect the correct type and throw an error
	fmt.Println("'uint64' values range from", unsignedInt64Min, "to", unsignedInt64Max)

	fmt.Println("")
	floatNumber := 3.14
	fmt.Println("'float32' and 'float64' can hold 32 (64) bit floatin point numbers, like", floatNumber)

	fmt.Println("")
	complexNumber := 1 + 5i
	fmt.Println("'complex64' and 'complex128' hold values with 'float32' (and 'float64') integer and imaginary values, like", complexNumber)

	fmt.Println("")
	fmt.Println("'byte' is an alias for 'uint8'")
	fmt.Println("'rune' is an alias for 'uint32'")

	fmt.Println("")
	fmt.Println("'string' is a collection of bytes")
}
