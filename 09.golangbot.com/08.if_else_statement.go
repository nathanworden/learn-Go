// If with assignment

// There is a variant of the if statement which includes an optional shorthand assignment statement that is executed before the condition is evaluated. Its syntx is

// if assignment-statement; condition {

// }

package main

import (
	"fmt"
)

func main() {
	if num := 10; num%2 == 0 { //checks if number is even
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}
