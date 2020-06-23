package main

import (
	"github.com/kokiebisu/go/pointers/example1"
)

// In the above code, myTestFunc takes in an integer
// variable and makes a copy of it for use within
// the context of the function body. Any changes
// we make to a within myTestFunc will only persist
// inside the body of the myTestFunc function.
// func myTestFunc(a int) {
// 	a += 3
// 	fmt.Println(a)
// }

// func myTestFunc1(a *int) {
// 	*a += 3
// 	fmt.Println(*a)
// }

func main() {
	a := 2
	// myTestFunc(a)
	// fmt.Println(a) // the value will stay 2
	// myTestFunc1(&a)
	// fmt.Println(a)

	example1.YearsUntilRetirement(&a)
}
