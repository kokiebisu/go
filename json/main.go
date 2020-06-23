package main

import (
	"fmt"

	"github.com/kokiebisu/go/json/example1"
	"github.com/kokiebisu/go/json/example2"
)

// Note - It is typically recommended to try and define the
// structs, if you happen to know the structure of the data coming back.

func main() {
	fmt.Println("example1")
	example1.ReadJson("./data.json")
	example2.ReadJsonWithoutStruct("./data.json")
}
