// The empty interface

// The interface type that specifies zero methods is known as the empty interface

// interface{}

// An empty interface may hold values of any type. (Every type implements at least zero methods)

//  Empty interfaces are used by code that handles value of unkown type. For example, fmt.Print takes any number of arguments of type `interface{}`

package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// <nil> <nil>
// 42  int
// (hello, string)
