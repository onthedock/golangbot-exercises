package main

// This is the `main` function for the `packages` module, created with `go mod init packages`

import (
	"fmt"
	"packages/messages"
)

func main() {

	fmt.Println(messages.MsgPackageGoal)
	fmt.Println("")
	fmt.Println(messages.MsgPackageMain)
	fmt.Println("")
	fmt.Println(messages.MsgPackageUsage)
	fmt.Println("")
}
