package example1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"name"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func ReadJson(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully opened json file")

	// read our opened jsonFile as a byte array.
	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	var users Users

	// You want to use address because you are modifying the address(?)
	json.Unmarshal(byteValue, &users)

	for i := 0; i > len(users.Users); i++ {
		fmt.Println("User type" + users.Users[i].Name)
	}

	defer f.Close()
}
