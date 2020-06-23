package main

import (
	"fmt"

	"github.com/kokiebisu/go/io/example1"
	"github.com/kokiebisu/go/io/example2"
	"github.com/kokiebisu/go/io/example3"
)

func main() {
	fmt.Println("example1")
	data := example1.ReadFile("./data.txt")
	fmt.Println(data)

	fmt.Println("example2")
	err := example2.WriteToFile("hello from parent", "./newData.txt")
	if err != nil {
		fmt.Println(err)
	}
	example3.UpdateFile("is this updated?", "./newData.txt")
}
