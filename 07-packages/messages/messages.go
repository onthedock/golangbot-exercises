package messages

// Constants names need to be capitalized to be "exported"
// i.e. accessible from other packages (like the "parent" package, "packages")

const (
	MsgPackageGoal string = `Packages are used to organize Go source code for better reusability and readability.
Packages are a collection of Go sources files that reside in the same directory.
Packages provide code compartmentalization and hence it becomes easy to maintain Go projects.`
	MsgPackageMain string = `Every executable Go application must contain the main function.
This function is the entry point for execution.
The main function should reside in the main package.`
	MsgPackageUsage string = `'package packagename' specifies that a particular source file belongs to package 'packagename'.
This should be the first line of every go source file.`
)
