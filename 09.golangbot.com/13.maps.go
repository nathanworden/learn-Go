// Maps

// A map is a builtin type in Go that is used to stroe key-value pairs. Let's take the example of a startup with a few employees. For simplicity, let's assume that the first name of all these employeees is unique. We are looking for a data structure to store teh salary of each employee. A map will be a perfect fit for this use case. The name of the emplyee can be the key and the salary can be the value. Maps are similar to dictionaries in other languages such as Python.

// A map can be created by passing the type of key and value to the make function. The following is the syntax to create a new map.

// make(map[tyoe of key]typeof value)

// emplyeeSalary := make(map[string]int)

package main

import (
	"fmt"
)

// func main() {
// 	employeeSalary := make(map[string]int)
// 	employeeSalary["steve"] = 12000
// 	employeeSalary["jamie"] = 15000
// 	employeeSalary["mike"] = 9000
// 	fmt.Println("emplyeeSalary map contents:", employeeSalary)
// }

// func main() {
// 	employeeSalary := map[string]int{
// 		"steve": 12000,
// 		"jamie": 15000,
// 		"mike":  9000,
// 	}
// 	employee := "jamie"
// 	salary := employeeSalary[employee]
// 	fmt.Println("Salary of", employee, "is", salary)
// }

// Here is the syntax to find out whether a particular key is present in a map:

// value, ok := map[key]

// If `ok` is true, then the key is present and its value is present in the variable `value`, else the key is absent

// Iterate over all elements in a map

// The `range` form of the `for` loop is used to iterate over all elements of a map

func main() {
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike":  9000,
	}
	fmt.Println("Contents of the map")
	for key, value := range employeeSalary {
		fmt.Printf("employeeSalary[%s] = %d\n", key, value)
	}
}
