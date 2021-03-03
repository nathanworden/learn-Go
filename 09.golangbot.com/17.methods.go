// Methods

// A method is just a function with a special receiver type between the func keyword and the method name. The receiver can either be a struct type or non-struct type.

//  The syntax of a method declaration is:

// func (t Type) methodName(parameter list) {

// }

// package main

// import (
// 	"fmt"
// )

// type Employee struct {
// 	name     string
// 	salary   int
// 	currency string
// }

// // displaySalary() method has Employee as the receiver

// func (e Employee) displaySalary() {
// 	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
// }

// func main() {
// 	emp1 := Employee{
// 		name:     "Sam Adolf",
// 		salary:   5000,
// 		currency: "$",
// 	}
// 	emp1.displaySalary() // Calling displaySalary() method of Employee type

// Why do we have methods when we can write the same program using functions?

// Go is not a pure object oriented programming language and it does not support classes. Hence methods on types are a way to achieve behavior similar to classes. Methods allow a logical grouping of behavior related to a type similar to classes.

// Methods with the same name can be defined on different types whereas functions with the same names are not allowed.

// Pointer receivers vs value receivers

// The difference between a value and a pointer receiver is, changese made inside a method with a pointer receiver is visible to the caller, whereas this is not the case with a value receiver.

package main

import (
	"fmt"
)

type Employee struct {
	name string
	age  int
}

/*
Method with value receiver
*/

func (e Employee) changeName(newName string) {
	e.name = newName
}

func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

func main() {
	e := Employee{
		name: "Mark Andrew",
		age:  50,
	}

	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	(&e).changeAge(51)
	fmt.Printf("\nEmployee age after change: %d", e.age)
}

// Mark Andrew
// Mark Andrew

// 50
// 51

// When to use a pointer receiver and when to use a value receiver

// Generally, pointer receivers can be used when changes made to the receiver inside the method should be visible to the caller.

// Methods of anonymous struct fields

// Methods belonging to anonymouse fields of a struct can be called as if they belong to the structure where the anonymouse field is defined.

// Value receivers in methods vs Value arguments in functions

// When a function has a value argument, it will accept only a value argument

// When a method has a value receiver, it will accept both pointer and value receivers

// Pointer receivers in methods vs Pointer arguments in functions

// Similar to value arguments, functions with pointer arguments will accept only pointers whereas methods with pointer recivers will accept both pointer and value receivers.

// Methods with non-struct receivers

// So far we have defined methods only on struct types. It is also possible to define methods on non-struct types, but there is a catch. To define a method on a type, the definition of the receiver type and the definition of the method should be present in the same package.
