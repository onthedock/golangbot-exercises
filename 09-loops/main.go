package main

import "fmt"

func main() {
	const msg string = `A loop statement is used to execute a block of code repeatedly.
The "for" loop structure is
	for <initialisation>; <condition>; <post> {
		...
	}
<initialisation>, <condition> and <post> are all optional.

<initialisation> is executed only once.
<condition> is evaluated after initialisation. If the result is "true"
After each successful execution of the "for" block, the <post> statement is executed.
Then, the <condition> is checked again.`
	fmt.Println(msg)
}
