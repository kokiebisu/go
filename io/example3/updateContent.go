package example3

import (
	"fmt"
	"os"
)

// os.O_WRONLY: write only tells the computer you are only going to writ to the file, not read
// os.O_CREATE tells the computer to create the file if it doesn't exist
// os.O_APPEND: append tells the computer to append to the end of the file instead of overwritting or truncating it

func UpdateFile(body string, path string) {
	f, err := os.OpenFile(path, 0, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(body)
	if err != nil {
		panic(err)
	}

	fmt.Println("successfully updated file content")

}
