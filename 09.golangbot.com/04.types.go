// Types

// The following are the basic types available in go

// bool
// Numeric Types:
// int8, int16, int32, int64, int
// uint8, uint16, uint64, uint
// float32, float64
// complex64, complex128
// byte
// rune
// string

// The following is a simple program to illustrate integer and floating point types

package main

import (
	"fmt"
)

func main() {
	a, b := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", a, b)
	sum := a + b
	diff := a - b
	fmt.Println("sum", sum, "diff", diff)
}


// Complex types

// The builtin function complex is used to construct a complex number with real and imaginary parts. The complex function has the following definition:

// func complex(r, i FloatType) ComplexType

// Type Conversion

// Go is very strict about explicit typing. There is no automatic type promotion or conversion. Let's look at what this means with an example.

package main

import (
	"fmt"
)

func main() {
	i := 55 		//int
	j := 67.8 	// float64
	sum := i + j // int + float64 not allowed
	fmt.Println(sum)
}

func main() {
	i := 55 			//int
	j := 67.8			//flat64
	sum := i + int(j) // j is converted to int
	fmt.Println(sum)
}