package main

import (
	"fmt"

	overflow "github.com/kokiebisu/go/types/example1"
	addition "github.com/kokiebisu/go/types/example2"
)

func main() {
	// test 1
	var test1 overflow.Overflow
	fmt.Println("test1 runs")
	test1.Run()

	// test 2
	var test2 addition.Addition
	fmt.Println("test2 runs")
	test2.Run()
}
