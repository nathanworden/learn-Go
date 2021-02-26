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