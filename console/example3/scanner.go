package example3

import (
	"bufio"
	"fmt"
	"os"
)

func Scan() {
	reader := bufio.NewScanner(os.Stdin)

	for reader.Scan() {
		fmt.Println(reader.Text())
	}

}
