package main

import (
	"fmt"

	"github.com/kokiebisu/go/functions/example1"
	"github.com/kokiebisu/go/functions/example2"
)

func main() {
	fmt.Println("testing 1")
	firstname := "ken"
	lastname := "okiebisu"
	fullname, err := example1.RandomFunction(firstname, lastname)
	if err != nil {
		fmt.Println("there was an error")
	}
	fmt.Println(fullname)

	fmt.Println("testing 2")
	address := example2.AddOne()
	fmt.Println(address)
}
