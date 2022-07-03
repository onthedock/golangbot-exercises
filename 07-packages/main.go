package main

import "fmt"

func main() {
	msgPackageGoal := `Packages are used to organize Go source code for better reusability and readability.
Packages are a collection of Go sources files that reside in the same directory.
Packages provide code compartmentalization and hence it becomes easy to maintain Go projects.`
	msgPackageMain := `Every executable Go application must contain the main function.
This function is the entry point for execution.
The main function should reside in the main package.`
	msgPackageUsage := `'package packagename' specifies that a particular source file belongs to package 'packagename'.
This should be the first line of every go source file.`
	fmt.Println(msgPackageGoal)
	fmt.Println("")
	fmt.Println(msgPackageMain)
	fmt.Println("")
	fmt.Println(msgPackageUsage)
	fmt.Println("")
}
