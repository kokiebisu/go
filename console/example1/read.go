package example1

import (
	"bufio"
	"fmt"
	"strings"
	"os"

)
// if you are wanting to read in full new line delimited sentences
func ReadFromConsole() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("------------")

	for {
		fmt.Print("->")
		text, _ := reader.ReadString('\n')
		// text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi\n", text) == 0 {
			fmt.Println("hello, Yourself")
		}
	}
}