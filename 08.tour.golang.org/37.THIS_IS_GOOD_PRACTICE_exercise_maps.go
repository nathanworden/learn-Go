// Exercise: Maps

// Implemnt `WordCount`. It should return a map of the counts of each "word" in the string `s`. The `wc.Test` function runs a test suite against the provided function and prints success or failure.

// You might find strings.Fields helpful

package main

import (
	"fmt"
	"strings"
	// "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	a := strings.Fields(s)
	for _, v := range a {
		m[v]++
	}
	return m
}

func main() {
	fmt.Println(WordCount("I am learning Go!"))
	fmt.Println(WordCount("The quick brown fox jumped over the lazy dog."))
	fmt.Println(WordCount("I ate a donut. Then I ate another donut."))
	fmt.Println(WordCount("A man a plan a canal panama."))
	// wc.Test(WordCount)
}
