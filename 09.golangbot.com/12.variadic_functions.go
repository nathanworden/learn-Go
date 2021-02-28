// Variadic Functions

// Functions in general accept only a fixed number of arguments. A variadic function is a function that accepts a variable number of arguments. If the last parameter of a function definition is prefixed by elipsis..., then the function can accept any number of arguments for that parameter.

// Only the last parameter of a function can be variadic.

// func hello(a int, b ...int) {

// }

// In the above function, the parameter `b` is variadic since it's prefized by ellipsis and it can accept any number of arguments. This function can be called by using the syntax.

// Find whether an integer exists in an input list of integers.

package main

import (
	"fmt"
)

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func main() {
	find(89, 89, 90, 95)          // 0
	find(45, 56, 67, 45, 90, 109) // 2
	find(78, 38, 56, 98)          //
	find(87)                      //
}
