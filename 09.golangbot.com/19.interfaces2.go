// Implementing interfaces using pointer receivers vs value receivers

// It is possible to implement interfaces using pointer receivers. There is a subtlety to be noted when implementing interfaces using pointer receivers.

package main

import "fmt"

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

func (p Person) Describe() { // implemented using value receiver
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() { // implemented using pointer receiver
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func main() {

}
