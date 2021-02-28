// Arrays and Slices

// An array is a collection of elements that belong to the same type. Mixing values of different types, for example an array that contains both strings and integers is not allowed in Go.

// Declaration
// An array belongs to type [n]T. n denotes the number of elements in an array and T represents the type of each element. The number of elements n is also part of the type.

// There are different ways to declare arrays. Lets look at them one by one.

package main

import (
	"fmt"
)

func main() {
	var a [3]int // int array with length 3
	fmt.Println(a)
}

// var a [3]int declares an integer array of length 3. All elements in an array are automatically assigned the zero value of the array type. In this case `a` is an integer array and hence all elements of `a` are assigned to `0`, the zero value of int. Running the above program will output [0 0 0]

// The index of an array starts from `0` and ends at length - 1.

// The size of the array is part of the type. Hence `[5]int` and `[25]int` are distinct types. Because of this, arrays cannot be resized. Don't worry about this restriction since `slices` exist to overcome this.

// Arrays are value types
// Arrays in Go are value types and not reference types. This means that when they are assigned to a new variable, a copy of the original array is assigned to the new variabe. If changes are made to the new variable, it will not be reflected in the original array.

package main

import "fmt"

func main() {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
}


// Iterating arrays using range
// The for loop can be used to iterate over elements of an arary

func main() {
	a := [...]float64{67.7, 89.8, 21, 78}
	for i := 0l i < len(a); i++ { //looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	}
}

// Go provides a better and concise way to iterate over an array by using the range form of the for loop. range returns both the index and the value at that index. Let's rewrite the code using range. We will also find the sum of all the elements of the array.

func main() {
	a := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range a { // range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)
}

// Multidimensional arrays

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

// Slices
// A slice is a convenient, flexible and powerful wrapper on top of an array. Slices do not own any data on their own. They are just references to existing arrays.

// creating a slice
// A slice with elements of type T is represented by []T

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] // creates a slice from a[1] to a[3]
	fmt.Println(b)
}

// Lets look at another way to create a slice

func main() {
	c := []int{6, 7, 8} // creates an array and return a slice reference
	fmt.Println(c)
}

// Modifying a slice
// A slice does not own any data of its own. It is just a preresentation of the undrelying array. Any modifications done to the slice will be reflected in the underlying array.

func main() {
	darr := [...]int{57, 89, 90, 82, 100 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)
}


57, 89, 90, 82, 100, 78, 67, 69, 59
57, 89, 91, 83, 101, 78, 67, 69, 59

// Length and capacity of a slice

// The length of the slice is the number of elements in the slice. THe capaacity of the slice is the number of elements in the underlying array starting from the index from which the slice is created.



// creating a slice using make

// func make([]T, len, cap) []T can be used to create a slice by passing the type, length and capacity. The capacity parameter is optional and defaults to the length. The make function creates an array and returns a slice reference to it.

func main() {
	i := make([]int, 5, 5)
	fmt.Println(i)
}

// Apending to a slice

// As we already know, arrays are restricted to fixed length and their length cannot be increased. Slices are dynamic and new elements can be appended to the slice using the `append` function. The definition of the `append` function is func append( s []T, x ... T) []T

// x ...T in the function definition means that the function accepts variable number of arguments for the parameter x. These type of functions are called variadic functions.

// Passing a slice to a function
// slice can be though of as being represented internally by a structure type. This is ow it looks:

type slice struct {
	Length int
	Capacity int
	ZerothElement *byte
}


func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func main() {
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)
	fumt.Println("slice after function call", nos)
}




// When arrays are passed to functions as parameters, they are passed by value and the original array is unchanged.

package main

import "fmt"

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)

}

func main() {
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) // num is passed by value
	fmt.Println("after passing to function ", num)
}

