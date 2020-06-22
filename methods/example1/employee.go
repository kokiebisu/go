package example1

import "fmt"

type Employee struct {
	Name string
}

//One thing you should note when working with methods is that, 
// like functions, they create copies of the arguments passed into it.
// To avoid this, we can use pointer receivers when defining our methods.

// Updates the name of the employee
func (e *Employee) UpdateName(newName string) {
	// Without the pointer, it only makes a copy
	e.Name = newName
}

// Prints the name
func (e *Employee) PrintName() {
	fmt.Println(e.Name)
}
