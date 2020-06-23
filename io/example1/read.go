package example1

import (
	"fmt"
	"io/ioutil"
)

func ReadFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("there was an error")
		return "something wrong"
	}
	return string(data)
}
