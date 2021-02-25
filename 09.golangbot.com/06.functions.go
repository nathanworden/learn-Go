// function syntax is:

// func functionname(parametername type) returntype {
// 		function body
// }

// It is possible to return multiple values from a function.

package main

import "fmt"

func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	otherStuff := 89.9
	area += otherStuff
	return
}

func main() {
	fmt.Println(rectProps(2, 5))
}
