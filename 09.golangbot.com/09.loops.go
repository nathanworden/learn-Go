// Loops

// The `for` is the only loop available in Go. Go doesn't have while or do while loops which are present in other languages.

// for loop syntx
// for initialization; condition; post {

// }

package main

import (
	"fmt"
)

// func main() {
// 	n := 5
// 	for i := 0; i < n; i++ {
// 		for j := 0; j <= i; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

// func repeat(times int, str string) string {
// 	n := times
// 	var output string = ""
// 	for i := 0; i < n; i++ {
// 		output = output + str
// 	}
// 	return output
// }

// func diamond(num int) {
// 	fmt.Println(repeat(4, " "), repeat(1, "*"), repeat(4, " "))
// 	fmt.Println(repeat(3, " "), repeat(3, "*"), repeat(3, " "))
// 	fmt.Println(repeat(2, " "), repeat(5, "*"), repeat(2, " "))
// 	fmt.Println(repeat(1, " "), repeat(7, "*"), repeat(1, " "))
// 	fmt.Println(repeat(2, " "), repeat(5, "*"), repeat(2, " "))
// 	fmt.Println(repeat(3, " "), repeat(3, "*"), repeat(3, " "))
// 	fmt.Println(repeat(4, " "), repeat(1, "*"), repeat(4, " "))
// }

// func main() {

// 	diamond(9)
// }

// func main() {
// 	i := 0
// 	for ; i <= 10; i += 2 { // initialization and post are omitted
// 		fmt.Printf("%d ", i)
// 	}
// }

func main() {
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
}
