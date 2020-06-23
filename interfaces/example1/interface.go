package example1

import "fmt"

// By defining a function that takes in an interface{}, 
// we essentially give ourselves the flexibility to pass 
// in anything we want. It’s a Go programmers way of saying, 
// this function takes in something, but I don’t necessarily
// care about its type.

func myFunc(a interface{}) {
	fmt.Println("this function can print any type")
	fmt.Println(a)
}

func Run() {
	var myAge int
	myAge = 15

	myFunc(myAge)
}