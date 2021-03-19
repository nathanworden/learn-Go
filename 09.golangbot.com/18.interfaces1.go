// Interfaces part 1

// What is an interface?

// In Go, an interface is a set of method signatures. When a type provides definition for all the methods in the interface, it is said to implement the interface. It is much similar to the OOP world. Interface specifies what methods a type should have and the type decides how to implement these methods.

// For example WashingMachine can be an interface with method signatures Cleaning() and Drying(). An type which provides definition for Cleaning() and Drying() methods is said to implement the WashingMachine interface.

// Declaring and implementing an interface.

package main

// interface definition

// type VowelsFinder interface {
// 	FindVowels() []rune
// }

// type MyString string

// // MyString implements VowelsFinder
// func (ms MyString) FindVowels() []rune {
// 	var vowels []rune
// 	for _, rune := range ms {
// 		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
// 			vowels = append(vowels, rune)
// 		}
// 	}
// 	return vowels
// }

// func main() {
// 	name := MyString("Sam Anderson")
// 	var v VowelsFinder
// 	v = name // possible since MyString implements VowelsFinder
// 	fmt.Printf("Vowels are %c", v.FindVowels())
// }

// PRactical use of an interface

// In the example above if we used name.FindVowels() instead of v.Findvowels(), it would have also worked and there would have been no use for the interface.

// Now let's look at a practical use of an interface.

// type SalaryCalculator interface {
// 	CalculateSalary() int
// }

// type Permanent struct {
// 	empId    int
// 	basicpay int
// 	pf       int
// }

// type Contract struct {
// 	empId    int
// 	basicpay int
// }

// type Freelancer struct {
// 	empId       int
// 	ratePerHour int
// 	totalHours  int
// }

// // salary of permanent employee is the sum of basic pay and pf
// func (p Permanent) CalculateSalary() int {
// 	return p.basicpay + p.pf
// }

// // salary of contract employee is the basic pay alone
// func (c Contract) CalculateSalary() int {
// 	return c.basicpay
// }

// // salary of freelancer
// func (f Freelancer) CalculateSalary() int {
// 	return f.ratePerHour * f.totalHours
// }

// /*
// total expense is calculated by iterating through the SalaryCalculator slice and summing the salaries of the individual employees
// */

// func totalExpense(s []SalaryCalculator) {
// 	expense := 0
// 	for _, v := range s {
// 		expense = expense + v.CalculateSalary()
// 	}
// 	fmt.Printf("Total Expense Per Month $%d", expense)
// }

// func main() {
// 	pemp1 := Permanent{
// 		empId:    1,
// 		basicpay: 5000,
// 		pf:       20,
// 	}
// 	pemp2 := Permanent{
// 		empId:    2,
// 		basicpay: 6000,
// 		pf:       30,
// 	}
// 	cemp1 := Contract{
// 		empId:    3,
// 		basicpay: 3000,
// 	}
// 	freelancer1 := Freelancer{
// 		empId:       4,
// 		ratePerHour: 70,
// 		totalHours:  120,
// 	}
// 	freelancer2 := Freelancer{
// 		empId:       5,
// 		ratePerHour: 100,
// 		totalHours:  100,
// 	}
// 	employees := []SalaryCalculator{pemp1, pemp2, cemp1, freelancer1, freelancer2}
// 	totalExpense(employees)

// }

// Interface internal representation

// An interface can be thought of as being represented internally by a typle (type, value). Type is the underlying concrete type of the interface and value holds the value of the concrete type.

// type Worker interface {
// 	Work()
// }

// type Person struct {
// 	name string
// }

// func (p Person) Work() {
// 	fmt.Println(p.name, "is working")
// }

// func describe(w Worker) {
// 	fmt.Printf("Interface type %T value %v\n", w, w)
// }

// func main() {
// 	p := Person{
// 		name: "Naveen",
// 	}
// 	var w Worker = p
// 	describe(w)
// 	w.Work()
// }

// Interface type: Person value: Naveen
// "Naveen is working"

// Empty interface

// An interface that has zero methods is called an empty interface. It is represented as `interface{}`. Since the empty interface has zero methods, all types implement the empty interface.

// func describe(i interface{}) {
// 	fmt.Printf("Type = %T, value = %v\n", i, i)
// }

// func main() {
// 	s := "Hello World"
// 	describe(s)
// 	i := 55
// 	describe(i)
// 	strt := struct {
// 		name string
// 	}{
// 		name: "Naveen R",
// 	}
// 	describe(strt)
// }

// Type = string, value = Hello World
// Type = int, value = 55
// Type = struct, value = {name: "Naveen R"}

// Type assertion

// Type assertion is used to extract the underlying value of the interface.

// i.(T) is the syntax which is used to get the underlying value of interface i whose concrete type is T

import (
	"fmt"
)

func assert(i interface{}) {
	s := i.(int) // get the underlying int value from i
	fmt.Println(s)
}

func main() {
	var s interface{} = 56
	assert(s)
}
