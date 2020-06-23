package example2

import (
	"fmt"
	"os"
	"bufio"
)
// If you are only needing single character input then use ReadRune()
func ReadUnicode() {
	reader := bufio.NewReader(os.Stdin)


	for {
		char, _, err := reader.ReadRune()

		if err != nil {
			fmt.Println("something went wrong when initializing rune")
		}
		switch char {
		case 'a':
			fmt.Println("small a pressed")
			break
		case 'A':
			fmt.Println("capital a presse")
			break
		}
	}

}