// Pointers

// What is a pointer?
// A pointer is a variable which stores the memory address of another variable.

// Declaring pointers
// *T is the type of the pointer variable which points to a value of type T.

package main

import (
	"fmt"
)

// func main() {
// 	b := 255
// 	var a *int = &b
// 	fmt.Printf("Type of a is %T\n", a)
// 	fmt.Println("address of b is", a)
// }

// The & operator is used to get the address of a variable. In line no. 9 of the above program we are assigning the address of `b` to `a` whose type is `*int`. Now a is said to point to b. When we print the value in `a`, the address of `b` will be printed.

// Zero value of a pointer
// The zero value of a pointer is nil

// func main() {
// 	a := 25
// 	var b *int
// 	if b == nil {
// 		fmt.Println("b is", b)
// 		b = &a
// 		fmt.Println("b after initialization is", b)
// 	}
// }

// Creating pointers using new function

// The `new` function takes a type as argument and returns a pointer to a newly allocated zero value of the type passed as argument.

// func main() {
// 	size := new(int)
// 	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
// 	*size = 85
// 	fmt.Println("New size value is", *size)
// }

// Dereferencing a pointer
// Dereferencing a pointer means accessing the value of the variable which the pointer points to. `*a` is the syntax to deference a.

// func main() {
// 	b := 255
// 	a := &b
// 	fmt.Println("address of b is", a)
// 	fmt.Println("value of b is", *a)
// }

// Passing pointer to a function

// func change(val *int) {
// 	*val = 55
// }

// func main() {
// 	a := 58
// 	fmt.Println("value of a before function call is", a)
// 	b := &a
// 	change(b)
// 	fmt.Println("value of a after function call is", a)
// }

// Returning a pointer from a function

// func hello() *int {
// 	i := 5
// 	return &i
// }

// func main() {
// 	d := hello()
// 	fmt.Println("Value of d is", *d)
// }

// Do not pass a pointer to an array as an argument to a fucntion. Use a slice instead.

// Lets assume that we want to make some modifications to an array inside the functino and the changes made to that array insdie the function should be visible to the caller. One way of doing this is to pass a pointer to an array as an argument to the function.

// func modify(arr *[3]int) {
// 	(*arr)[0] = 90
// }

// func main() {
// 	a := [3]int{89, 90, 91}
// 	modify(&a)
// 	fmt.Println(a)
// }

func modify(sls []int) {
	sls[0] = 90
}

func main() {
	a := [3]int{89, 90, 91}
	modify(a[:])
	fmt.Println(a)
}

// Go does not supprot pointer arithemtic.
