package main

import (
	"fmt"
	"log"
	"net/http"
)

func greetings(rw http.ResponseWriter, w *http.Request) {
	fmt.Fprint(rw, "Hello World")
}

func runServer() {
	http.HandleFunc("/", greetings)
	log.Fatal(http.ListenAndServeTLS(":8000", "server.crt", "server.key", nil))
}

func main() {
	runServer()
}
