package example1

import (
	"fmt"
	"os/exec"
)

func Execute() {

	// Checks the ls command
	output, err := exec.Command("ls").Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output[:]))

	// Checks the pwd command
	output, err = exec.Command("pwd").Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
