// The init and post statements are optional in a for loop

package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
	}

	fmt.Println(sum)
}
