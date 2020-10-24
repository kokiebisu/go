package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/unmarshal", unmarshal)
	http.ListenAndServe(":8080", nil)
}

type user struct {
	FirstName string
	LastName string
}

var u = user{
	"Kenichi",
	"Okiebisu",
}

var jsonuser = `{"FirstName":"Kenichi","LastName":"Okiebisu"}`
