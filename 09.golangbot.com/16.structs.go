// Structs

// What is a struct?

// A struct is a user-defined type that represents a collection of fields. It can be used in places where it makes sense to group the data into a single unit rather than having each of them as separate values.

// Declaring a struct

// type Emplyee stuct {
// 	firstName string
// 	lastName string
// 	age int
// }

// The above `Employee` struct is called a named struct because it creates anew data type named `Employee` using which `Employee` structs can be created.

// Creating named structs

// Let's declare a named struct Employee using the follwing simple program

package main

import "fmt"

// type Employee struct {
// 	firstName string
// 	lastName  string
// 	age       int
// 	salary    int
// }

// func main() {
// 	// creating struct specifying field names
// 	emp1 := Employee{
// 		firstName: "Sam",
// 		age:       25,
// 		salary:    500,
// 		lastName:  "Anderson",
// 	}

// 	// creating struct without specifying field names
// 	emp2 := Employee{"Thomas", "Paul", 29, 800}

// 	fmt.Println("Employee 1", emp1)
// 	fmt.Println("Employee 2", emp2)
// }

//Creating anonymous structs

// It is possible to declare structs without creating a new data type. These types of structs are called anonymous structs.

func main() {
	emp3 := struct {
		firstName string
		lastName  string
		age       int
		salary    int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee 3", emp3)
}
