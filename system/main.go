package main

import (
	"fmt"
	"runtime"

	"github.com/kokiebisu/go/system/example1"
)

func main() {
	// Checks operating system
	if runtime.GOOS == "windows" {
		fmt.Println("this is run by windows")
	} else {
		example1.Execute()
	}
}
