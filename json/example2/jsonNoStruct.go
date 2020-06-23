package example2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadJsonWithoutStruct(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	byte, _ := ioutil.ReadAll(file)

	// we know the key will be a string
	// we just don't know the value but we want to be flexible
	var result map[string]interface{}

	// We don't need to assign to variable because it is assigning it to the result address
	json.Unmarshal(byte, &result)

	fmt.Println(result["users"])

}
