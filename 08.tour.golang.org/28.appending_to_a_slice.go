package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	//append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// len=0 cap=0 []
// len=1 cap=1 [0]
// len=2 cap=2 [0, 1]
// len=5 cap=5 [0, 1, 2, 3, 4]

// It is common to append new elements to a slice, and so Go provides a built-in append funciton. The documentation of the built-in package describes append.

//  https://golang.org/pkg/builtin/#append

// func append(s []T, vs ...T) []T

// The first parameter 's' of append is a slice of type 'T', and the rest are 'T' values to append to the slice.

// The resulting value of 'append' is a slice containing all the elements of the original slice plus the provided values.

//  If the backing array of 's' is too small to fit all the given values a bigger array will be allocated. The returned slice will poitn to the newly allocated array.
