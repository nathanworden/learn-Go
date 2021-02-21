// Defer

// A defer statement defers the execution of a function until the surrounding function returns.

// The deferred call's argument are evaluated immediatley, but the function call is not executed until the surrounding fucntion returns.

package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
