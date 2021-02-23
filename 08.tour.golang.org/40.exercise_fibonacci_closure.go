// Exercise: Fibonacci closure

// Let's have some fun with functions.

//  Implement a `fibonacci` function that returns a fucntion (a clousre) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5...)

// Solution: https://gist.github.com/weppos/7843653

package main

import "fmt"

// fibonacci is a function that returns a function that returns an int

func fibonacci() func() int {
	f2, f1 := 0, 1
	return func() int {
		f := f2
		f2, f1 = f1, f+f1

		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i += 1 {
		fmt.Println(f())
	}
}
