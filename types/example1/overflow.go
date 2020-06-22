package example1

import "fmt"

// Runs the run command from the Overflow struct
type Overflow struct{}

// This is the function that displays an overflow
func (o Overflow) Run() {
	fmt.Println("This will show the example of an overflow")

	var myInt int8

	for i := 0; i < 130; i++ {
		myInt++
	}
	fmt.Println(myInt)
}
