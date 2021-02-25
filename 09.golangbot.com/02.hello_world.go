package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

// One subtle difference between the `go run` and `go build`/`go install` commands is `go run` rquires the name of the `.go` file as an argument.

// Every go file must start with the `package name` statement.

// Packages are used to provide code compartmentalization and reusability. The main function should always reside in the main package.
