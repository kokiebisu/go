package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type data struct {
	Code int64
	Descrip string
}

func main() {
	rcvd := `[{"Code":200,"Descrip":"StatusOK"},{"Code":301,"Descrip":"StatusMovedPermanently"}]`
	
	var data []data
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln("Not correct")
	}

	for _, v := range data {
		fmt.Println("code: ", v.Code)
		fmt.Println("description: ", v.Descrip)
	}
}